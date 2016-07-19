// This file was generated by counterfeiter
package spacesfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/securitygroups/spaces"
)

type FakeSecurityGroupSpaceBinder struct {
	BindSpaceStub        func(securityGroupGUID string, spaceGUID string) error
	bindSpaceMutex       sync.RWMutex
	bindSpaceArgsForCall []struct {
		securityGroupGUID string
		spaceGUID         string
	}
	bindSpaceReturns struct {
		result1 error
	}
	UnbindSpaceStub        func(securityGroupGUID string, spaceGUID string) error
	unbindSpaceMutex       sync.RWMutex
	unbindSpaceArgsForCall []struct {
		securityGroupGUID string
		spaceGUID         string
	}
	unbindSpaceReturns struct {
		result1 error
	}
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpace(securityGroupGUID string, spaceGUID string) error {
	fake.bindSpaceMutex.Lock()
	fake.bindSpaceArgsForCall = append(fake.bindSpaceArgsForCall, struct {
		securityGroupGUID string
		spaceGUID         string
	}{securityGroupGUID, spaceGUID})
	fake.bindSpaceMutex.Unlock()
	if fake.BindSpaceStub != nil {
		return fake.BindSpaceStub(securityGroupGUID, spaceGUID)
	} else {
		return fake.bindSpaceReturns.result1
	}
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceCallCount() int {
	fake.bindSpaceMutex.RLock()
	defer fake.bindSpaceMutex.RUnlock()
	return len(fake.bindSpaceArgsForCall)
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceArgsForCall(i int) (string, string) {
	fake.bindSpaceMutex.RLock()
	defer fake.bindSpaceMutex.RUnlock()
	return fake.bindSpaceArgsForCall[i].securityGroupGUID, fake.bindSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceReturns(result1 error) {
	fake.BindSpaceStub = nil
	fake.bindSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpace(securityGroupGUID string, spaceGUID string) error {
	fake.unbindSpaceMutex.Lock()
	fake.unbindSpaceArgsForCall = append(fake.unbindSpaceArgsForCall, struct {
		securityGroupGUID string
		spaceGUID         string
	}{securityGroupGUID, spaceGUID})
	fake.unbindSpaceMutex.Unlock()
	if fake.UnbindSpaceStub != nil {
		return fake.UnbindSpaceStub(securityGroupGUID, spaceGUID)
	} else {
		return fake.unbindSpaceReturns.result1
	}
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceCallCount() int {
	fake.unbindSpaceMutex.RLock()
	defer fake.unbindSpaceMutex.RUnlock()
	return len(fake.unbindSpaceArgsForCall)
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceArgsForCall(i int) (string, string) {
	fake.unbindSpaceMutex.RLock()
	defer fake.unbindSpaceMutex.RUnlock()
	return fake.unbindSpaceArgsForCall[i].securityGroupGUID, fake.unbindSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceReturns(result1 error) {
	fake.UnbindSpaceStub = nil
	fake.unbindSpaceReturns = struct {
		result1 error
	}{result1}
}

var _ spaces.SecurityGroupSpaceBinder = new(FakeSecurityGroupSpaceBinder)
