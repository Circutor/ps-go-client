// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	modelss "github.com/circutor/ps-go-client/pkg/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PowerStudioAPIMock is an autogenerated mock type for the PowerStudioAPI type
type PowerStudioAPIMock struct {
	mock.Mock
}

// PsAllDevices provides a mock function with given fields:
func (_m *PowerStudioAPIMock) PsAllDevices() (*modelss.Devices, error) {
	ret := _m.Called()

	var r0 *modelss.Devices
	if rf, ok := ret.Get(0).(func() *modelss.Devices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelss.Devices)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PsDeviceInfo provides a mock function with given fields: ids
func (_m *PowerStudioAPIMock) PsDeviceInfo(ids []string) (*modelss.DevicesInfo, error) {
	ret := _m.Called(ids)

	var r0 *modelss.DevicesInfo
	if rf, ok := ret.Get(0).(func([]string) *modelss.DevicesInfo); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelss.DevicesInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PsDevicesSelectionInfo provides a mock function with given fields:
func (_m *PowerStudioAPIMock) PsDevicesSelectionInfo() (*modelss.DevicesSelectionInfo, error) {
	ret := _m.Called()

	var r0 *modelss.DevicesSelectionInfo
	if rf, ok := ret.Get(0).(func() *modelss.DevicesSelectionInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelss.DevicesSelectionInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PsRecords provides a mock function with given fields: begin, end, period, vars
func (_m *PowerStudioAPIMock) PsRecords(begin time.Time, end time.Time, period int, vars []string) (*modelss.RecordGroup, error) {
	ret := _m.Called(begin, end, period, vars)

	var r0 *modelss.RecordGroup
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, []string) *modelss.RecordGroup); ok {
		r0 = rf(begin, end, period, vars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelss.RecordGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int, []string) error); ok {
		r1 = rf(begin, end, period, vars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PsVarInfo provides a mock function with given fields: ids, vars
func (_m *PowerStudioAPIMock) PsVarInfo(ids []string, vars []string) (*modelss.VarInfo, error) {
	ret := _m.Called(ids, vars)

	var r0 *modelss.VarInfo
	if rf, ok := ret.Get(0).(func([]string, []string) *modelss.VarInfo); ok {
		r0 = rf(ids, vars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelss.VarInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, []string) error); ok {
		r1 = rf(ids, vars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PsVarValue provides a mock function with given fields: ids, vars
func (_m *PowerStudioAPIMock) PsVarValue(ids []string, vars []string) (*modelss.Values, error) {
	ret := _m.Called(ids, vars)

	var r0 *modelss.Values
	if rf, ok := ret.Get(0).(func([]string, []string) *modelss.Values); ok {
		r0 = rf(ids, vars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelss.Values)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, []string) error); ok {
		r1 = rf(ids, vars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPowerStudioAPIMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewPowerStudioAPIMock creates a new instance of PowerStudioAPIMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPowerStudioAPIMock(t mockConstructorTestingTNewPowerStudioAPIMock) *PowerStudioAPIMock {
	mock := &PowerStudioAPIMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}