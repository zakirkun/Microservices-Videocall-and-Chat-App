// Code generated by mockery v2.14.1. DO NOT EDIT.

package mock

import (
	events "github.com/Slimo300/MicroservicesChatApp/backend/lib/msgqueue/events"
	mock "github.com/stretchr/testify/mock"

	models "github.com/Slimo300/MicroservicesChatApp/backend/message-service/models"

	uuid "github.com/google/uuid"
)

// MockMessageDB is an autogenerated mock type for the DBLayer type
type MockMessageDB struct {
	mock.Mock
}

// AddMessage provides a mock function with given fields: event
func (_m *MockMessageDB) AddMessage(event events.MessageSentEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.MessageSentEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroupMembers provides a mock function with given fields: event
func (_m *MockMessageDB) DeleteGroupMembers(event events.GroupDeletedEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.GroupDeletedEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMember provides a mock function with given fields: event
func (_m *MockMessageDB) DeleteMember(event events.MemberDeletedEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.MemberDeletedEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessageForEveryone provides a mock function with given fields: userID, messageID
func (_m *MockMessageDB) DeleteMessageForEveryone(userID uuid.UUID, messageID uuid.UUID) (models.Message, error) {
	ret := _m.Called(userID, messageID)

	var r0 models.Message
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) models.Message); ok {
		r0 = rf(userID, messageID)
	} else {
		r0 = ret.Get(0).(models.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessageForYourself provides a mock function with given fields: userID, messageID
func (_m *MockMessageDB) DeleteMessageForYourself(userID uuid.UUID, messageID uuid.UUID) (models.Message, error) {
	ret := _m.Called(userID, messageID)

	var r0 models.Message
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) models.Message); ok {
		r0 = rf(userID, messageID)
	} else {
		r0 = ret.Get(0).(models.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupMessages provides a mock function with given fields: userID, groupID, offset, num
func (_m *MockMessageDB) GetGroupMessages(userID uuid.UUID, groupID uuid.UUID, offset int, num int) ([]models.Message, error) {
	ret := _m.Called(userID, groupID, offset, num)

	var r0 []models.Message
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID, int, int) []models.Message); ok {
		r0 = rf(userID, groupID, offset, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID, int, int) error); ok {
		r1 = rf(userID, groupID, offset, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyMember provides a mock function with given fields: event
func (_m *MockMessageDB) ModifyMember(event events.MemberUpdatedEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.MemberUpdatedEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMember provides a mock function with given fields: event
func (_m *MockMessageDB) NewMember(event events.MemberCreatedEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.MemberCreatedEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockMessageDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMessageDB creates a new instance of MockMessageDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMessageDB(t mockConstructorTestingTNewMockMessageDB) *MockMessageDB {
	mock := &MockMessageDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
