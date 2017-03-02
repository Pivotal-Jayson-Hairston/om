// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ProductUploader struct {
	UploadStub        func(api.UploadProductInput) (api.UploadProductOutput, error)
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		arg1 api.UploadProductInput
	}
	uploadReturns struct {
		result1 api.UploadProductOutput
		result2 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 api.UploadProductOutput
		result2 error
	}
	CheckProductAvailabilityStub        func(string, string) (bool, error)
	checkProductAvailabilityMutex       sync.RWMutex
	checkProductAvailabilityArgsForCall []struct {
		arg1 string
		arg2 string
	}
	checkProductAvailabilityReturns struct {
		result1 bool
		result2 error
	}
	checkProductAvailabilityReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	TrashStub        func() error
	trashMutex       sync.RWMutex
	trashArgsForCall []struct{}
	trashReturns     struct {
		result1 error
	}
	trashReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProductUploader) Upload(arg1 api.UploadProductInput) (api.UploadProductOutput, error) {
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		arg1 api.UploadProductInput
	}{arg1})
	fake.recordInvocation("Upload", []interface{}{arg1})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uploadReturns.result1, fake.uploadReturns.result2
}

func (fake *ProductUploader) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *ProductUploader) UploadArgsForCall(i int) api.UploadProductInput {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].arg1
}

func (fake *ProductUploader) UploadReturns(result1 api.UploadProductOutput, result2 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 api.UploadProductOutput
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) UploadReturnsOnCall(i int, result1 api.UploadProductOutput, result2 error) {
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 api.UploadProductOutput
			result2 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 api.UploadProductOutput
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) CheckProductAvailability(arg1 string, arg2 string) (bool, error) {
	fake.checkProductAvailabilityMutex.Lock()
	ret, specificReturn := fake.checkProductAvailabilityReturnsOnCall[len(fake.checkProductAvailabilityArgsForCall)]
	fake.checkProductAvailabilityArgsForCall = append(fake.checkProductAvailabilityArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CheckProductAvailability", []interface{}{arg1, arg2})
	fake.checkProductAvailabilityMutex.Unlock()
	if fake.CheckProductAvailabilityStub != nil {
		return fake.CheckProductAvailabilityStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkProductAvailabilityReturns.result1, fake.checkProductAvailabilityReturns.result2
}

func (fake *ProductUploader) CheckProductAvailabilityCallCount() int {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return len(fake.checkProductAvailabilityArgsForCall)
}

func (fake *ProductUploader) CheckProductAvailabilityArgsForCall(i int) (string, string) {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return fake.checkProductAvailabilityArgsForCall[i].arg1, fake.checkProductAvailabilityArgsForCall[i].arg2
}

func (fake *ProductUploader) CheckProductAvailabilityReturns(result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	fake.checkProductAvailabilityReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) CheckProductAvailabilityReturnsOnCall(i int, result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	if fake.checkProductAvailabilityReturnsOnCall == nil {
		fake.checkProductAvailabilityReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkProductAvailabilityReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) Trash() error {
	fake.trashMutex.Lock()
	ret, specificReturn := fake.trashReturnsOnCall[len(fake.trashArgsForCall)]
	fake.trashArgsForCall = append(fake.trashArgsForCall, struct{}{})
	fake.recordInvocation("Trash", []interface{}{})
	fake.trashMutex.Unlock()
	if fake.TrashStub != nil {
		return fake.TrashStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.trashReturns.result1
}

func (fake *ProductUploader) TrashCallCount() int {
	fake.trashMutex.RLock()
	defer fake.trashMutex.RUnlock()
	return len(fake.trashArgsForCall)
}

func (fake *ProductUploader) TrashReturns(result1 error) {
	fake.TrashStub = nil
	fake.trashReturns = struct {
		result1 error
	}{result1}
}

func (fake *ProductUploader) TrashReturnsOnCall(i int, result1 error) {
	fake.TrashStub = nil
	if fake.trashReturnsOnCall == nil {
		fake.trashReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.trashReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ProductUploader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	fake.trashMutex.RLock()
	defer fake.trashMutex.RUnlock()
	return fake.invocations
}

func (fake *ProductUploader) recordInvocation(key string, args []interface{}) {
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
