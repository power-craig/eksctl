// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/eksctl/pkg/drain"
	"github.com/weaveworks/eksctl/pkg/drain/evictor"
	v1 "k8s.io/api/core/v1"
)

type FakeDrainer struct {
	CanUseEvictionsStub        func() error
	canUseEvictionsMutex       sync.RWMutex
	canUseEvictionsArgsForCall []struct {
	}
	canUseEvictionsReturns struct {
		result1 error
	}
	canUseEvictionsReturnsOnCall map[int]struct {
		result1 error
	}
	EvictOrDeletePodStub        func(v1.Pod) error
	evictOrDeletePodMutex       sync.RWMutex
	evictOrDeletePodArgsForCall []struct {
		arg1 v1.Pod
	}
	evictOrDeletePodReturns struct {
		result1 error
	}
	evictOrDeletePodReturnsOnCall map[int]struct {
		result1 error
	}
	GetPodsForDeletionStub        func(string) (*evictor.PodDeleteList, []error)
	getPodsForDeletionMutex       sync.RWMutex
	getPodsForDeletionArgsForCall []struct {
		arg1 string
	}
	getPodsForDeletionReturns struct {
		result1 *evictor.PodDeleteList
		result2 []error
	}
	getPodsForDeletionReturnsOnCall map[int]struct {
		result1 *evictor.PodDeleteList
		result2 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDrainer) CanUseEvictions() error {
	fake.canUseEvictionsMutex.Lock()
	ret, specificReturn := fake.canUseEvictionsReturnsOnCall[len(fake.canUseEvictionsArgsForCall)]
	fake.canUseEvictionsArgsForCall = append(fake.canUseEvictionsArgsForCall, struct {
	}{})
	fake.recordInvocation("CanUseEvictions", []interface{}{})
	fake.canUseEvictionsMutex.Unlock()
	if fake.CanUseEvictionsStub != nil {
		return fake.CanUseEvictionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.canUseEvictionsReturns
	return fakeReturns.result1
}

func (fake *FakeDrainer) CanUseEvictionsCallCount() int {
	fake.canUseEvictionsMutex.RLock()
	defer fake.canUseEvictionsMutex.RUnlock()
	return len(fake.canUseEvictionsArgsForCall)
}

func (fake *FakeDrainer) CanUseEvictionsCalls(stub func() error) {
	fake.canUseEvictionsMutex.Lock()
	defer fake.canUseEvictionsMutex.Unlock()
	fake.CanUseEvictionsStub = stub
}

func (fake *FakeDrainer) CanUseEvictionsReturns(result1 error) {
	fake.canUseEvictionsMutex.Lock()
	defer fake.canUseEvictionsMutex.Unlock()
	fake.CanUseEvictionsStub = nil
	fake.canUseEvictionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDrainer) CanUseEvictionsReturnsOnCall(i int, result1 error) {
	fake.canUseEvictionsMutex.Lock()
	defer fake.canUseEvictionsMutex.Unlock()
	fake.CanUseEvictionsStub = nil
	if fake.canUseEvictionsReturnsOnCall == nil {
		fake.canUseEvictionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.canUseEvictionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDrainer) EvictOrDeletePod(arg1 v1.Pod) error {
	fake.evictOrDeletePodMutex.Lock()
	ret, specificReturn := fake.evictOrDeletePodReturnsOnCall[len(fake.evictOrDeletePodArgsForCall)]
	fake.evictOrDeletePodArgsForCall = append(fake.evictOrDeletePodArgsForCall, struct {
		arg1 v1.Pod
	}{arg1})
	fake.recordInvocation("EvictOrDeletePod", []interface{}{arg1})
	fake.evictOrDeletePodMutex.Unlock()
	if fake.EvictOrDeletePodStub != nil {
		return fake.EvictOrDeletePodStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.evictOrDeletePodReturns
	return fakeReturns.result1
}

func (fake *FakeDrainer) EvictOrDeletePodCallCount() int {
	fake.evictOrDeletePodMutex.RLock()
	defer fake.evictOrDeletePodMutex.RUnlock()
	return len(fake.evictOrDeletePodArgsForCall)
}

func (fake *FakeDrainer) EvictOrDeletePodCalls(stub func(v1.Pod) error) {
	fake.evictOrDeletePodMutex.Lock()
	defer fake.evictOrDeletePodMutex.Unlock()
	fake.EvictOrDeletePodStub = stub
}

func (fake *FakeDrainer) EvictOrDeletePodArgsForCall(i int) v1.Pod {
	fake.evictOrDeletePodMutex.RLock()
	defer fake.evictOrDeletePodMutex.RUnlock()
	argsForCall := fake.evictOrDeletePodArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDrainer) EvictOrDeletePodReturns(result1 error) {
	fake.evictOrDeletePodMutex.Lock()
	defer fake.evictOrDeletePodMutex.Unlock()
	fake.EvictOrDeletePodStub = nil
	fake.evictOrDeletePodReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDrainer) EvictOrDeletePodReturnsOnCall(i int, result1 error) {
	fake.evictOrDeletePodMutex.Lock()
	defer fake.evictOrDeletePodMutex.Unlock()
	fake.EvictOrDeletePodStub = nil
	if fake.evictOrDeletePodReturnsOnCall == nil {
		fake.evictOrDeletePodReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evictOrDeletePodReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDrainer) GetPodsForDeletion(arg1 string) (*evictor.PodDeleteList, []error) {
	fake.getPodsForDeletionMutex.Lock()
	ret, specificReturn := fake.getPodsForDeletionReturnsOnCall[len(fake.getPodsForDeletionArgsForCall)]
	fake.getPodsForDeletionArgsForCall = append(fake.getPodsForDeletionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetPodsForDeletion", []interface{}{arg1})
	fake.getPodsForDeletionMutex.Unlock()
	if fake.GetPodsForDeletionStub != nil {
		return fake.GetPodsForDeletionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPodsForDeletionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDrainer) GetPodsForDeletionCallCount() int {
	fake.getPodsForDeletionMutex.RLock()
	defer fake.getPodsForDeletionMutex.RUnlock()
	return len(fake.getPodsForDeletionArgsForCall)
}

func (fake *FakeDrainer) GetPodsForDeletionCalls(stub func(string) (*evictor.PodDeleteList, []error)) {
	fake.getPodsForDeletionMutex.Lock()
	defer fake.getPodsForDeletionMutex.Unlock()
	fake.GetPodsForDeletionStub = stub
}

func (fake *FakeDrainer) GetPodsForDeletionArgsForCall(i int) string {
	fake.getPodsForDeletionMutex.RLock()
	defer fake.getPodsForDeletionMutex.RUnlock()
	argsForCall := fake.getPodsForDeletionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDrainer) GetPodsForDeletionReturns(result1 *evictor.PodDeleteList, result2 []error) {
	fake.getPodsForDeletionMutex.Lock()
	defer fake.getPodsForDeletionMutex.Unlock()
	fake.GetPodsForDeletionStub = nil
	fake.getPodsForDeletionReturns = struct {
		result1 *evictor.PodDeleteList
		result2 []error
	}{result1, result2}
}

func (fake *FakeDrainer) GetPodsForDeletionReturnsOnCall(i int, result1 *evictor.PodDeleteList, result2 []error) {
	fake.getPodsForDeletionMutex.Lock()
	defer fake.getPodsForDeletionMutex.Unlock()
	fake.GetPodsForDeletionStub = nil
	if fake.getPodsForDeletionReturnsOnCall == nil {
		fake.getPodsForDeletionReturnsOnCall = make(map[int]struct {
			result1 *evictor.PodDeleteList
			result2 []error
		})
	}
	fake.getPodsForDeletionReturnsOnCall[i] = struct {
		result1 *evictor.PodDeleteList
		result2 []error
	}{result1, result2}
}

func (fake *FakeDrainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.canUseEvictionsMutex.RLock()
	defer fake.canUseEvictionsMutex.RUnlock()
	fake.evictOrDeletePodMutex.RLock()
	defer fake.evictOrDeletePodMutex.RUnlock()
	fake.getPodsForDeletionMutex.RLock()
	defer fake.getPodsForDeletionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDrainer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ drain.Drainer = new(FakeDrainer)
