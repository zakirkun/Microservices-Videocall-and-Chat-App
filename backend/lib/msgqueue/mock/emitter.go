// Code generated by mockery v2.14.1. DO NOT EDIT.

package mock

import (
	msgqueue "github.com/Slimo300/MicroservicesChatApp/backend/lib/msgqueue"
	mock "github.com/stretchr/testify/mock"
)

// MockEmitter is an autogenerated mock type for the EventEmitter type
type MockEmitter struct {
	mock.Mock
}

// Emit provides a mock function with given fields: event
func (_m *MockEmitter) Emit(event msgqueue.Event) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(msgqueue.Event) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockEmitter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEmitter creates a new instance of MockEmitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEmitter(t mockConstructorTestingTNewMockEmitter) *MockEmitter {
	mock := &MockEmitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
