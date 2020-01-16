// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/micro/go-micro/client"

	mock "github.com/stretchr/testify/mock"

	reporterpb "github.com/paysuper/paysuper-proto/go/reporterpb"
)

// ReporterService is an autogenerated mock type for the ReporterService type
type ReporterService struct {
	mock.Mock
}

// CreateFile provides a mock function with given fields: ctx, in, opts
func (_m *ReporterService) CreateFile(ctx context.Context, in *reporterpb.ReportFile, opts ...client.CallOption) (*reporterpb.CreateFileResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *reporterpb.CreateFileResponse
	if rf, ok := ret.Get(0).(func(context.Context, *reporterpb.ReportFile, ...client.CallOption) *reporterpb.CreateFileResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*reporterpb.CreateFileResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *reporterpb.ReportFile, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}