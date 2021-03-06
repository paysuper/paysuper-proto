// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	taxpb "github.com/paysuper/paysuper-proto/go/taxpb"
	mock "github.com/stretchr/testify/mock"
)

// TaxServiceHandler is an autogenerated mock type for the TaxServiceHandler type
type TaxServiceHandler struct {
	mock.Mock
}

// CreateOrUpdate provides a mock function with given fields: _a0, _a1, _a2
func (_m *TaxServiceHandler) CreateOrUpdate(_a0 context.Context, _a1 *taxpb.TaxRate, _a2 *taxpb.TaxRate) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.TaxRate, *taxpb.TaxRate) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRateById provides a mock function with given fields: _a0, _a1, _a2
func (_m *TaxServiceHandler) DeleteRateById(_a0 context.Context, _a1 *taxpb.DeleteRateRequest, _a2 *taxpb.DeleteRateResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.DeleteRateRequest, *taxpb.DeleteRateResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRate provides a mock function with given fields: _a0, _a1, _a2
func (_m *TaxServiceHandler) GetRate(_a0 context.Context, _a1 *taxpb.GeoIdentity, _a2 *taxpb.TaxRate) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.GeoIdentity, *taxpb.TaxRate) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRates provides a mock function with given fields: _a0, _a1, _a2
func (_m *TaxServiceHandler) GetRates(_a0 context.Context, _a1 *taxpb.GetRatesRequest, _a2 *taxpb.GetRatesResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.GetRatesRequest, *taxpb.GetRatesResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
