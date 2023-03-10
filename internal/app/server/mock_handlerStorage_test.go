// Code generated by mockery v2.16.0. DO NOT EDIT.

package server

import mock "github.com/stretchr/testify/mock"

// mockHandlerStorage is an autogenerated mock type for the handlerStorage type
type mockHandlerStorage struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *mockHandlerStorage) Get(key string) (string, bool) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, value
func (_m *mockHandlerStorage) Set(key string, value string) {
	_m.Called(key, value)
}

type mockConstructorTestingTnewMockHandlerStorage interface {
	mock.TestingT
	Cleanup(func())
}

// newMockHandlerStorage creates a new instance of mockHandlerStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockHandlerStorage(t mockConstructorTestingTnewMockHandlerStorage) *mockHandlerStorage {
	mock := &mockHandlerStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
