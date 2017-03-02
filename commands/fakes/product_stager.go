// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ProductStager struct {
	StageStub        func(api.StageProductInput) error
	stageMutex       sync.RWMutex
	stageArgsForCall []struct {
		arg1 api.StageProductInput
	}
	stageReturns struct {
		result1 error
	}
	stageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProductStager) Stage(arg1 api.StageProductInput) error {
	fake.stageMutex.Lock()
	ret, specificReturn := fake.stageReturnsOnCall[len(fake.stageArgsForCall)]
	fake.stageArgsForCall = append(fake.stageArgsForCall, struct {
		arg1 api.StageProductInput
	}{arg1})
	fake.recordInvocation("Stage", []interface{}{arg1})
	fake.stageMutex.Unlock()
	if fake.StageStub != nil {
		return fake.StageStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stageReturns.result1
}

func (fake *ProductStager) StageCallCount() int {
	fake.stageMutex.RLock()
	defer fake.stageMutex.RUnlock()
	return len(fake.stageArgsForCall)
}

func (fake *ProductStager) StageArgsForCall(i int) api.StageProductInput {
	fake.stageMutex.RLock()
	defer fake.stageMutex.RUnlock()
	return fake.stageArgsForCall[i].arg1
}

func (fake *ProductStager) StageReturns(result1 error) {
	fake.StageStub = nil
	fake.stageReturns = struct {
		result1 error
	}{result1}
}

func (fake *ProductStager) StageReturnsOnCall(i int, result1 error) {
	fake.StageStub = nil
	if fake.stageReturnsOnCall == nil {
		fake.stageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ProductStager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stageMutex.RLock()
	defer fake.stageMutex.RUnlock()
	return fake.invocations
}

func (fake *ProductStager) recordInvocation(key string, args []interface{}) {
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
