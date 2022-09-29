// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TCPDelivery is an autogenerated mock type for the TCPDelivery type
type TCPDelivery struct {
	mock.Mock
}

// ListenAndServeWitSignal provides a mock function with given fields:
func (_m *TCPDelivery) ListenAndServeWitSignal() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTCPDelivery interface {
	mock.TestingT
	Cleanup(func())
}

// NewTCPDelivery creates a new instance of TCPDelivery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTCPDelivery(t mockConstructorTestingTNewTCPDelivery) *TCPDelivery {
	mock := &TCPDelivery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}