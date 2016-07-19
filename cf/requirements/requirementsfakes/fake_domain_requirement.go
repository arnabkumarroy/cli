// This file was generated by counterfeiter
package requirementsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeDomainRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 error
	}
	GetDomainStub        func() models.DomainFields
	getDomainMutex       sync.RWMutex
	getDomainArgsForCall []struct{}
	getDomainReturns     struct {
		result1 models.DomainFields
	}
}

func (fake *FakeDomainRequirement) Execute() error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeDomainRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeDomainRequirement) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDomainRequirement) GetDomain() models.DomainFields {
	fake.getDomainMutex.Lock()
	fake.getDomainArgsForCall = append(fake.getDomainArgsForCall, struct{}{})
	fake.getDomainMutex.Unlock()
	if fake.GetDomainStub != nil {
		return fake.GetDomainStub()
	} else {
		return fake.getDomainReturns.result1
	}
}

func (fake *FakeDomainRequirement) GetDomainCallCount() int {
	fake.getDomainMutex.RLock()
	defer fake.getDomainMutex.RUnlock()
	return len(fake.getDomainArgsForCall)
}

func (fake *FakeDomainRequirement) GetDomainReturns(result1 models.DomainFields) {
	fake.GetDomainStub = nil
	fake.getDomainReturns = struct {
		result1 models.DomainFields
	}{result1}
}

var _ requirements.DomainRequirement = new(FakeDomainRequirement)
