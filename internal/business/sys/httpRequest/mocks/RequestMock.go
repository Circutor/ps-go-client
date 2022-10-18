// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// RequestMock is an autogenerated mock type for the Request type
type RequestMock struct {
	mock.Mock
}

// NewRequest provides a mock function with given fields: method, url, body, query
func (_m *RequestMock) NewRequest(method string, url string, body io.Reader, query []map[string]interface{}) ([]byte, int, error) {
	ret := _m.Called(method, url, body, query)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string, io.Reader, []map[string]interface{}) []byte); ok {
		r0 = rf(method, url, body, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, string, io.Reader, []map[string]interface{}) int); ok {
		r1 = rf(method, url, body, query)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, io.Reader, []map[string]interface{}) error); ok {
		r2 = rf(method, url, body, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewRequestMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequestMock creates a new instance of RequestMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequestMock(t mockConstructorTestingTNewRequestMock) *RequestMock {
	mock := &RequestMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}