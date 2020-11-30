// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/unistack-org/micro/v3/client"

	mock "github.com/stretchr/testify/mock"

	recurringpb "github.com/paysuper/paysuper-proto/go/recurringpb"
)

// RepositoryService is an autogenerated mock type for the RepositoryService type
type RepositoryService struct {
	mock.Mock
}

// AddSubscription provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) AddSubscription(ctx context.Context, req *recurringpb.Subscription, opts ...client.CallOption) (*recurringpb.AddSubscriptionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.AddSubscriptionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.Subscription, ...client.CallOption) *recurringpb.AddSubscriptionResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.AddSubscriptionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.Subscription, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSavedCard provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) DeleteSavedCard(ctx context.Context, req *recurringpb.DeleteSavedCardRequest, opts ...client.CallOption) (*recurringpb.DeleteSavedCardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.DeleteSavedCardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.DeleteSavedCardRequest, ...client.CallOption) *recurringpb.DeleteSavedCardResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.DeleteSavedCardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.DeleteSavedCardRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSubscription provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) DeleteSubscription(ctx context.Context, req *recurringpb.Subscription, opts ...client.CallOption) (*recurringpb.DeleteSubscriptionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.DeleteSubscriptionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.Subscription, ...client.CallOption) *recurringpb.DeleteSubscriptionResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.DeleteSubscriptionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.Subscription, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSavedCardById provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) FindSavedCardById(ctx context.Context, req *recurringpb.FindByStringValue, opts ...client.CallOption) (*recurringpb.SavedCard, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.SavedCard
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.FindByStringValue, ...client.CallOption) *recurringpb.SavedCard); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.SavedCard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.FindByStringValue, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSavedCards provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) FindSavedCards(ctx context.Context, req *recurringpb.SavedCardRequest, opts ...client.CallOption) (*recurringpb.SavedCardList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.SavedCardList
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.SavedCardRequest, ...client.CallOption) *recurringpb.SavedCardList); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.SavedCardList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.SavedCardRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSubscriptions provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) FindSubscriptions(ctx context.Context, req *recurringpb.FindSubscriptionsRequest, opts ...client.CallOption) (*recurringpb.FindSubscriptionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.FindSubscriptionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.FindSubscriptionsRequest, ...client.CallOption) *recurringpb.FindSubscriptionsResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.FindSubscriptionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.FindSubscriptionsRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscription provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) GetSubscription(ctx context.Context, req *recurringpb.GetSubscriptionRequest, opts ...client.CallOption) (*recurringpb.GetSubscriptionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.GetSubscriptionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.GetSubscriptionRequest, ...client.CallOption) *recurringpb.GetSubscriptionResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.GetSubscriptionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.GetSubscriptionRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertSavedCard provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) InsertSavedCard(ctx context.Context, req *recurringpb.SavedCardRequest, opts ...client.CallOption) (*recurringpb.Result, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.Result
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.SavedCardRequest, ...client.CallOption) *recurringpb.Result); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.SavedCardRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSubscription provides a mock function with given fields: ctx, req, opts
func (_m *RepositoryService) UpdateSubscription(ctx context.Context, req *recurringpb.Subscription, opts ...client.CallOption) (*recurringpb.UpdateSubscriptionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *recurringpb.UpdateSubscriptionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.Subscription, ...client.CallOption) *recurringpb.UpdateSubscriptionResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recurringpb.UpdateSubscriptionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recurringpb.Subscription, ...client.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
