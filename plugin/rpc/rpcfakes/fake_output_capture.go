// This file was generated by counterfeiter
package rpcfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/cli/plugin/rpc"
)

type FakeOutputCapture struct {
	SetOutputBucketStub        func(io.Writer)
	setOutputBucketMutex       sync.RWMutex
	setOutputBucketArgsForCall []struct {
		arg1 io.Writer
	}
}

func (fake *FakeOutputCapture) SetOutputBucket(arg1 io.Writer) {
	fake.setOutputBucketMutex.Lock()
	fake.setOutputBucketArgsForCall = append(fake.setOutputBucketArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.setOutputBucketMutex.Unlock()
	if fake.SetOutputBucketStub != nil {
		fake.SetOutputBucketStub(arg1)
	}
}

func (fake *FakeOutputCapture) SetOutputBucketCallCount() int {
	fake.setOutputBucketMutex.RLock()
	defer fake.setOutputBucketMutex.RUnlock()
	return len(fake.setOutputBucketArgsForCall)
}

func (fake *FakeOutputCapture) SetOutputBucketArgsForCall(i int) io.Writer {
	fake.setOutputBucketMutex.RLock()
	defer fake.setOutputBucketMutex.RUnlock()
	return fake.setOutputBucketArgsForCall[i].arg1
}

var _ rpc.OutputCapture = new(FakeOutputCapture)
