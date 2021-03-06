// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/micro/go-micro/client"

	mock "github.com/stretchr/testify/mock"

	notifierpb "github.com/paysuper/paysuper-proto/go/notifierpb"
)

// NotifierService is an autogenerated mock type for the NotifierService type
type NotifierService struct {
	mock.Mock
}

// CheckUser provides a mock function with given fields: ctx, in, opts
func (_m *NotifierService) CheckUser(ctx context.Context, in *notifierpb.CheckUserRequest, opts ...client.CallOption) (*notifierpb.CheckUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *notifierpb.CheckUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *notifierpb.CheckUserRequest, ...client.CallOption) *notifierpb.CheckUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notifierpb.CheckUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *notifierpb.CheckUserRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
