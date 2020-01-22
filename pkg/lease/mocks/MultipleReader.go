// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import lease "github.com/Optum/dce/pkg/lease"
import mock "github.com/stretchr/testify/mock"

// MultipleReader is an autogenerated mock type for the MultipleReader type
type MultipleReader struct {
	mock.Mock
}

// List provides a mock function with given fields: _a0
func (_m *MultipleReader) List(_a0 *lease.Lease) (*lease.Leases, error) {
	ret := _m.Called(_a0)

	var r0 *lease.Leases
	if rf, ok := ret.Get(0).(func(*lease.Lease) *lease.Leases); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*lease.Lease) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
