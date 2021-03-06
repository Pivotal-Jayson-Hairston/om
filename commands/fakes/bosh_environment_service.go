// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type BoshEnvironmentService struct {
	GetBoshEnvironmentStub        func() (api.GetBoshEnvironmentOutput, error)
	getBoshEnvironmentMutex       sync.RWMutex
	getBoshEnvironmentArgsForCall []struct {
	}
	getBoshEnvironmentReturns struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}
	getBoshEnvironmentReturnsOnCall map[int]struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}
	ListCertificateAuthoritiesStub        func() (api.CertificateAuthoritiesOutput, error)
	listCertificateAuthoritiesMutex       sync.RWMutex
	listCertificateAuthoritiesArgsForCall []struct {
	}
	listCertificateAuthoritiesReturns struct {
		result1 api.CertificateAuthoritiesOutput
		result2 error
	}
	listCertificateAuthoritiesReturnsOnCall map[int]struct {
		result1 api.CertificateAuthoritiesOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BoshEnvironmentService) GetBoshEnvironment() (api.GetBoshEnvironmentOutput, error) {
	fake.getBoshEnvironmentMutex.Lock()
	ret, specificReturn := fake.getBoshEnvironmentReturnsOnCall[len(fake.getBoshEnvironmentArgsForCall)]
	fake.getBoshEnvironmentArgsForCall = append(fake.getBoshEnvironmentArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBoshEnvironment", []interface{}{})
	fake.getBoshEnvironmentMutex.Unlock()
	if fake.GetBoshEnvironmentStub != nil {
		return fake.GetBoshEnvironmentStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBoshEnvironmentReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BoshEnvironmentService) GetBoshEnvironmentCallCount() int {
	fake.getBoshEnvironmentMutex.RLock()
	defer fake.getBoshEnvironmentMutex.RUnlock()
	return len(fake.getBoshEnvironmentArgsForCall)
}

func (fake *BoshEnvironmentService) GetBoshEnvironmentCalls(stub func() (api.GetBoshEnvironmentOutput, error)) {
	fake.getBoshEnvironmentMutex.Lock()
	defer fake.getBoshEnvironmentMutex.Unlock()
	fake.GetBoshEnvironmentStub = stub
}

func (fake *BoshEnvironmentService) GetBoshEnvironmentReturns(result1 api.GetBoshEnvironmentOutput, result2 error) {
	fake.getBoshEnvironmentMutex.Lock()
	defer fake.getBoshEnvironmentMutex.Unlock()
	fake.GetBoshEnvironmentStub = nil
	fake.getBoshEnvironmentReturns = struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}{result1, result2}
}

func (fake *BoshEnvironmentService) GetBoshEnvironmentReturnsOnCall(i int, result1 api.GetBoshEnvironmentOutput, result2 error) {
	fake.getBoshEnvironmentMutex.Lock()
	defer fake.getBoshEnvironmentMutex.Unlock()
	fake.GetBoshEnvironmentStub = nil
	if fake.getBoshEnvironmentReturnsOnCall == nil {
		fake.getBoshEnvironmentReturnsOnCall = make(map[int]struct {
			result1 api.GetBoshEnvironmentOutput
			result2 error
		})
	}
	fake.getBoshEnvironmentReturnsOnCall[i] = struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}{result1, result2}
}

func (fake *BoshEnvironmentService) ListCertificateAuthorities() (api.CertificateAuthoritiesOutput, error) {
	fake.listCertificateAuthoritiesMutex.Lock()
	ret, specificReturn := fake.listCertificateAuthoritiesReturnsOnCall[len(fake.listCertificateAuthoritiesArgsForCall)]
	fake.listCertificateAuthoritiesArgsForCall = append(fake.listCertificateAuthoritiesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListCertificateAuthorities", []interface{}{})
	fake.listCertificateAuthoritiesMutex.Unlock()
	if fake.ListCertificateAuthoritiesStub != nil {
		return fake.ListCertificateAuthoritiesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listCertificateAuthoritiesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BoshEnvironmentService) ListCertificateAuthoritiesCallCount() int {
	fake.listCertificateAuthoritiesMutex.RLock()
	defer fake.listCertificateAuthoritiesMutex.RUnlock()
	return len(fake.listCertificateAuthoritiesArgsForCall)
}

func (fake *BoshEnvironmentService) ListCertificateAuthoritiesCalls(stub func() (api.CertificateAuthoritiesOutput, error)) {
	fake.listCertificateAuthoritiesMutex.Lock()
	defer fake.listCertificateAuthoritiesMutex.Unlock()
	fake.ListCertificateAuthoritiesStub = stub
}

func (fake *BoshEnvironmentService) ListCertificateAuthoritiesReturns(result1 api.CertificateAuthoritiesOutput, result2 error) {
	fake.listCertificateAuthoritiesMutex.Lock()
	defer fake.listCertificateAuthoritiesMutex.Unlock()
	fake.ListCertificateAuthoritiesStub = nil
	fake.listCertificateAuthoritiesReturns = struct {
		result1 api.CertificateAuthoritiesOutput
		result2 error
	}{result1, result2}
}

func (fake *BoshEnvironmentService) ListCertificateAuthoritiesReturnsOnCall(i int, result1 api.CertificateAuthoritiesOutput, result2 error) {
	fake.listCertificateAuthoritiesMutex.Lock()
	defer fake.listCertificateAuthoritiesMutex.Unlock()
	fake.ListCertificateAuthoritiesStub = nil
	if fake.listCertificateAuthoritiesReturnsOnCall == nil {
		fake.listCertificateAuthoritiesReturnsOnCall = make(map[int]struct {
			result1 api.CertificateAuthoritiesOutput
			result2 error
		})
	}
	fake.listCertificateAuthoritiesReturnsOnCall[i] = struct {
		result1 api.CertificateAuthoritiesOutput
		result2 error
	}{result1, result2}
}

func (fake *BoshEnvironmentService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBoshEnvironmentMutex.RLock()
	defer fake.getBoshEnvironmentMutex.RUnlock()
	fake.listCertificateAuthoritiesMutex.RLock()
	defer fake.listCertificateAuthoritiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BoshEnvironmentService) recordInvocation(key string, args []interface{}) {
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
