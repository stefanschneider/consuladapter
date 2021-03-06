package consuladapter

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
)

type LostLockError string

func (e LostLockError) Error() string {
	return fmt.Sprintf("Lost lock '%s'", string(e))
}

var ErrInvalidSession = errors.New("invalid session")
var ErrDestroyed = errors.New("already destroyed")
var ErrCancelled = errors.New("cancelled")

type Session struct {
	kv *api.KV

	name       string
	sessionMgr SessionManager
	ttl        time.Duration
	noChecks   bool

	errCh chan error

	lock      sync.Mutex
	id        string
	destroyed bool
	doneCh    chan struct{}
	lostLock  string
}

func NewSession(sessionName string, ttl time.Duration, client *api.Client, sessionMgr SessionManager) (*Session, error) {
	return newSession(sessionName, ttl, false, client.KV(), sessionMgr)
}

func NewSessionNoChecks(sessionName string, ttl time.Duration, client *api.Client, sessionMgr SessionManager) (*Session, error) {
	return newSession(sessionName, ttl, true, client.KV(), sessionMgr)
}

func newSession(sessionName string, ttl time.Duration, noChecks bool, kv *api.KV, sessionMgr SessionManager) (*Session, error) {
	doneCh := make(chan struct{}, 1)
	errCh := make(chan error, 1)

	s := &Session{
		kv:         kv,
		name:       sessionName,
		sessionMgr: sessionMgr,
		ttl:        ttl,
		noChecks:   noChecks,
		doneCh:     doneCh,
		errCh:      errCh,
	}

	return s, nil
}

func (s *Session) ID() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.id
}

func (s *Session) Err() chan error {
	return s.errCh
}

func (s *Session) Destroy() {
	s.lock.Lock()
	s.destroy()
	s.lock.Unlock()
}

// Lock must be held
func (s *Session) destroy() {
	if s.destroyed == false {
		close(s.doneCh)

		if s.id != "" {
			s.sessionMgr.Destroy(s.id, nil)
		}

		s.destroyed = true
	}
}

// Lock must be held
func (s *Session) createSession() error {
	if s.destroyed {
		return ErrDestroyed
	}

	if s.id != "" {
		return nil
	}

	se := &api.SessionEntry{
		Name:      s.name,
		Behavior:  api.SessionBehaviorDelete,
		TTL:       s.ttl.String(),
		LockDelay: 1 * time.Nanosecond,
	}

	id, renewTTL, err := create(se, s.noChecks, s.sessionMgr)
	if err != nil {
		return err
	}

	s.id = id

	go func() {
		err := s.sessionMgr.RenewPeriodic(renewTTL, id, nil, s.doneCh)
		s.lock.Lock()
		lostLock := s.lostLock
		s.destroy()
		s.lock.Unlock()

		if lostLock != "" {
			err = LostLockError(lostLock)
		} else {
			err = convertError(err)
		}
		s.errCh <- err
	}()

	return err
}

func (s *Session) Recreate() (*Session, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	session, err := newSession(s.name, s.ttl, s.noChecks, s.kv, s.sessionMgr)
	if err != nil {
		return nil, err
	}

	err = session.createSession()
	if err != nil {
		return nil, err
	}

	return session, err
}

func (s *Session) AcquireLock(key string, value []byte) error {
	s.lock.Lock()
	err := s.createSession()
	s.lock.Unlock()
	if err != nil {
		return err
	}

	lock, err := s.sessionMgr.NewLock(s.id, key, value)
	if err != nil {
		return convertError(err)
	}

	lostCh, err := lock.Lock(s.doneCh)
	if err != nil {
		return convertError(err)
	}
	if lostCh == nil {
		return ErrCancelled
	}

	go func() {
		select {
		case <-lostCh:
			s.lock.Lock()
			s.lostLock = key
			s.destroy()
			s.lock.Unlock()
		case <-s.doneCh:
		}
	}()

	return nil
}

func (s *Session) SetPresence(key string, value []byte) (<-chan string, error) {
	s.lock.Lock()
	err := s.createSession()
	s.lock.Unlock()
	if err != nil {
		return nil, err
	}

	lock, err := s.sessionMgr.NewLock(s.id, key, value)
	if err != nil {
		return nil, convertError(err)
	}

	lostCh, err := lock.Lock(s.doneCh)
	if err != nil {
		return nil, convertError(err)
	}
	if lostCh == nil {
		return nil, ErrCancelled
	}

	presenceLost := make(chan string, 1)
	go func() {
		select {
		case <-lostCh:
			presenceLost <- key
		case <-s.doneCh:
		}
	}()

	return presenceLost, nil
}

func create(se *api.SessionEntry, noChecks bool, sessionMgr SessionManager) (string, string, error) {
	nodeName, err := sessionMgr.NodeName()
	if err != nil {
		return "", "", err
	}

	nodeSessions, _, err := sessionMgr.Node(nodeName, nil)
	if err != nil {
		return "", "", err
	}

	sessions := findSessions(se.Name, nodeSessions)
	if sessions != nil {
		for _, s := range sessions {
			_, err = sessionMgr.Destroy(s.ID, nil)
			if err != nil {
				return "", "", err
			}
		}
	}

	var f func(*api.SessionEntry, *api.WriteOptions) (string, *api.WriteMeta, error)
	if noChecks {
		f = sessionMgr.CreateNoChecks
	} else {
		f = sessionMgr.Create
	}

	id, _, err := f(se, nil)
	if err != nil {
		return "", "", err
	}

	return id, se.TTL, nil
}

func findSessions(name string, sessions []*api.SessionEntry) []*api.SessionEntry {
	var matches []*api.SessionEntry
	for _, session := range sessions {
		if session.Name == name {
			matches = append(matches, session)
		}
	}

	return matches
}

func convertError(err error) error {
	if err == nil {
		return err
	}

	if strings.Contains(err.Error(), "500 (Invalid session)") {
		return ErrInvalidSession
	}

	return err
}
