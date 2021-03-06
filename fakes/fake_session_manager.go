// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/consuladapter"
	"github.com/hashicorp/consul/api"
)

type FakeSessionManager struct {
	NodeNameStub        func() (string, error)
	nodeNameMutex       sync.RWMutex
	nodeNameArgsForCall []struct{}
	nodeNameReturns     struct {
		result1 string
		result2 error
	}
	NodeStub        func(node string, q *api.QueryOptions) ([]*api.SessionEntry, *api.QueryMeta, error)
	nodeMutex       sync.RWMutex
	nodeArgsForCall []struct {
		node string
		q    *api.QueryOptions
	}
	nodeReturns struct {
		result1 []*api.SessionEntry
		result2 *api.QueryMeta
		result3 error
	}
	CreateStub        func(se *api.SessionEntry, q *api.WriteOptions) (string, *api.WriteMeta, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		se *api.SessionEntry
		q  *api.WriteOptions
	}
	createReturns struct {
		result1 string
		result2 *api.WriteMeta
		result3 error
	}
	CreateNoChecksStub        func(se *api.SessionEntry, q *api.WriteOptions) (string, *api.WriteMeta, error)
	createNoChecksMutex       sync.RWMutex
	createNoChecksArgsForCall []struct {
		se *api.SessionEntry
		q  *api.WriteOptions
	}
	createNoChecksReturns struct {
		result1 string
		result2 *api.WriteMeta
		result3 error
	}
	DestroyStub        func(id string, q *api.WriteOptions) (*api.WriteMeta, error)
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		id string
		q  *api.WriteOptions
	}
	destroyReturns struct {
		result1 *api.WriteMeta
		result2 error
	}
	RenewStub        func(id string, q *api.WriteOptions) (*api.SessionEntry, *api.WriteMeta, error)
	renewMutex       sync.RWMutex
	renewArgsForCall []struct {
		id string
		q  *api.WriteOptions
	}
	renewReturns struct {
		result1 *api.SessionEntry
		result2 *api.WriteMeta
		result3 error
	}
	RenewPeriodicStub        func(initialTTL string, id string, q *api.WriteOptions, doneCh chan struct{}) error
	renewPeriodicMutex       sync.RWMutex
	renewPeriodicArgsForCall []struct {
		initialTTL string
		id         string
		q          *api.WriteOptions
		doneCh     chan struct{}
	}
	renewPeriodicReturns struct {
		result1 error
	}
	NewLockStub        func(sessionID, key string, value []byte) (consuladapter.Lock, error)
	newLockMutex       sync.RWMutex
	newLockArgsForCall []struct {
		sessionID string
		key       string
		value     []byte
	}
	newLockReturns struct {
		result1 consuladapter.Lock
		result2 error
	}
}

func (fake *FakeSessionManager) NodeName() (string, error) {
	fake.nodeNameMutex.Lock()
	fake.nodeNameArgsForCall = append(fake.nodeNameArgsForCall, struct{}{})
	fake.nodeNameMutex.Unlock()
	if fake.NodeNameStub != nil {
		return fake.NodeNameStub()
	} else {
		return fake.nodeNameReturns.result1, fake.nodeNameReturns.result2
	}
}

func (fake *FakeSessionManager) NodeNameCallCount() int {
	fake.nodeNameMutex.RLock()
	defer fake.nodeNameMutex.RUnlock()
	return len(fake.nodeNameArgsForCall)
}

func (fake *FakeSessionManager) NodeNameReturns(result1 string, result2 error) {
	fake.NodeNameStub = nil
	fake.nodeNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionManager) Node(node string, q *api.QueryOptions) ([]*api.SessionEntry, *api.QueryMeta, error) {
	fake.nodeMutex.Lock()
	fake.nodeArgsForCall = append(fake.nodeArgsForCall, struct {
		node string
		q    *api.QueryOptions
	}{node, q})
	fake.nodeMutex.Unlock()
	if fake.NodeStub != nil {
		return fake.NodeStub(node, q)
	} else {
		return fake.nodeReturns.result1, fake.nodeReturns.result2, fake.nodeReturns.result3
	}
}

func (fake *FakeSessionManager) NodeCallCount() int {
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	return len(fake.nodeArgsForCall)
}

func (fake *FakeSessionManager) NodeArgsForCall(i int) (string, *api.QueryOptions) {
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	return fake.nodeArgsForCall[i].node, fake.nodeArgsForCall[i].q
}

func (fake *FakeSessionManager) NodeReturns(result1 []*api.SessionEntry, result2 *api.QueryMeta, result3 error) {
	fake.NodeStub = nil
	fake.nodeReturns = struct {
		result1 []*api.SessionEntry
		result2 *api.QueryMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSessionManager) Create(se *api.SessionEntry, q *api.WriteOptions) (string, *api.WriteMeta, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		se *api.SessionEntry
		q  *api.WriteOptions
	}{se, q})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(se, q)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2, fake.createReturns.result3
	}
}

func (fake *FakeSessionManager) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSessionManager) CreateArgsForCall(i int) (*api.SessionEntry, *api.WriteOptions) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].se, fake.createArgsForCall[i].q
}

func (fake *FakeSessionManager) CreateReturns(result1 string, result2 *api.WriteMeta, result3 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 *api.WriteMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSessionManager) CreateNoChecks(se *api.SessionEntry, q *api.WriteOptions) (string, *api.WriteMeta, error) {
	fake.createNoChecksMutex.Lock()
	fake.createNoChecksArgsForCall = append(fake.createNoChecksArgsForCall, struct {
		se *api.SessionEntry
		q  *api.WriteOptions
	}{se, q})
	fake.createNoChecksMutex.Unlock()
	if fake.CreateNoChecksStub != nil {
		return fake.CreateNoChecksStub(se, q)
	} else {
		return fake.createNoChecksReturns.result1, fake.createNoChecksReturns.result2, fake.createNoChecksReturns.result3
	}
}

func (fake *FakeSessionManager) CreateNoChecksCallCount() int {
	fake.createNoChecksMutex.RLock()
	defer fake.createNoChecksMutex.RUnlock()
	return len(fake.createNoChecksArgsForCall)
}

func (fake *FakeSessionManager) CreateNoChecksArgsForCall(i int) (*api.SessionEntry, *api.WriteOptions) {
	fake.createNoChecksMutex.RLock()
	defer fake.createNoChecksMutex.RUnlock()
	return fake.createNoChecksArgsForCall[i].se, fake.createNoChecksArgsForCall[i].q
}

func (fake *FakeSessionManager) CreateNoChecksReturns(result1 string, result2 *api.WriteMeta, result3 error) {
	fake.CreateNoChecksStub = nil
	fake.createNoChecksReturns = struct {
		result1 string
		result2 *api.WriteMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSessionManager) Destroy(id string, q *api.WriteOptions) (*api.WriteMeta, error) {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		id string
		q  *api.WriteOptions
	}{id, q})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(id, q)
	} else {
		return fake.destroyReturns.result1, fake.destroyReturns.result2
	}
}

func (fake *FakeSessionManager) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeSessionManager) DestroyArgsForCall(i int) (string, *api.WriteOptions) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].id, fake.destroyArgsForCall[i].q
}

func (fake *FakeSessionManager) DestroyReturns(result1 *api.WriteMeta, result2 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 *api.WriteMeta
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionManager) Renew(id string, q *api.WriteOptions) (*api.SessionEntry, *api.WriteMeta, error) {
	fake.renewMutex.Lock()
	fake.renewArgsForCall = append(fake.renewArgsForCall, struct {
		id string
		q  *api.WriteOptions
	}{id, q})
	fake.renewMutex.Unlock()
	if fake.RenewStub != nil {
		return fake.RenewStub(id, q)
	} else {
		return fake.renewReturns.result1, fake.renewReturns.result2, fake.renewReturns.result3
	}
}

func (fake *FakeSessionManager) RenewCallCount() int {
	fake.renewMutex.RLock()
	defer fake.renewMutex.RUnlock()
	return len(fake.renewArgsForCall)
}

func (fake *FakeSessionManager) RenewArgsForCall(i int) (string, *api.WriteOptions) {
	fake.renewMutex.RLock()
	defer fake.renewMutex.RUnlock()
	return fake.renewArgsForCall[i].id, fake.renewArgsForCall[i].q
}

func (fake *FakeSessionManager) RenewReturns(result1 *api.SessionEntry, result2 *api.WriteMeta, result3 error) {
	fake.RenewStub = nil
	fake.renewReturns = struct {
		result1 *api.SessionEntry
		result2 *api.WriteMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSessionManager) RenewPeriodic(initialTTL string, id string, q *api.WriteOptions, doneCh chan struct{}) error {
	fake.renewPeriodicMutex.Lock()
	fake.renewPeriodicArgsForCall = append(fake.renewPeriodicArgsForCall, struct {
		initialTTL string
		id         string
		q          *api.WriteOptions
		doneCh     chan struct{}
	}{initialTTL, id, q, doneCh})
	fake.renewPeriodicMutex.Unlock()
	if fake.RenewPeriodicStub != nil {
		return fake.RenewPeriodicStub(initialTTL, id, q, doneCh)
	} else {
		return fake.renewPeriodicReturns.result1
	}
}

func (fake *FakeSessionManager) RenewPeriodicCallCount() int {
	fake.renewPeriodicMutex.RLock()
	defer fake.renewPeriodicMutex.RUnlock()
	return len(fake.renewPeriodicArgsForCall)
}

func (fake *FakeSessionManager) RenewPeriodicArgsForCall(i int) (string, string, *api.WriteOptions, chan struct{}) {
	fake.renewPeriodicMutex.RLock()
	defer fake.renewPeriodicMutex.RUnlock()
	return fake.renewPeriodicArgsForCall[i].initialTTL, fake.renewPeriodicArgsForCall[i].id, fake.renewPeriodicArgsForCall[i].q, fake.renewPeriodicArgsForCall[i].doneCh
}

func (fake *FakeSessionManager) RenewPeriodicReturns(result1 error) {
	fake.RenewPeriodicStub = nil
	fake.renewPeriodicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionManager) NewLock(sessionID string, key string, value []byte) (consuladapter.Lock, error) {
	fake.newLockMutex.Lock()
	fake.newLockArgsForCall = append(fake.newLockArgsForCall, struct {
		sessionID string
		key       string
		value     []byte
	}{sessionID, key, value})
	fake.newLockMutex.Unlock()
	if fake.NewLockStub != nil {
		return fake.NewLockStub(sessionID, key, value)
	} else {
		return fake.newLockReturns.result1, fake.newLockReturns.result2
	}
}

func (fake *FakeSessionManager) NewLockCallCount() int {
	fake.newLockMutex.RLock()
	defer fake.newLockMutex.RUnlock()
	return len(fake.newLockArgsForCall)
}

func (fake *FakeSessionManager) NewLockArgsForCall(i int) (string, string, []byte) {
	fake.newLockMutex.RLock()
	defer fake.newLockMutex.RUnlock()
	return fake.newLockArgsForCall[i].sessionID, fake.newLockArgsForCall[i].key, fake.newLockArgsForCall[i].value
}

func (fake *FakeSessionManager) NewLockReturns(result1 consuladapter.Lock, result2 error) {
	fake.NewLockStub = nil
	fake.newLockReturns = struct {
		result1 consuladapter.Lock
		result2 error
	}{result1, result2}
}

var _ consuladapter.SessionManager = new(FakeSessionManager)
