// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/micro/go-micro/client"

	currenciespb "github.com/paysuper/paysuper-proto/go/currenciespb"

	mock "github.com/stretchr/testify/mock"
)

// CurrencyRatesService is an autogenerated mock type for the CurrencyRatesService type
type CurrencyRatesService struct {
	mock.Mock
}

// AddCommonRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) AddCommonRateCorrectionRule(ctx context.Context, in *currenciespb.CommonCorrectionRule, opts ...client.CallOption) (*currenciespb.EmptyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.EmptyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.CommonCorrectionRule, ...client.CallOption) *currenciespb.EmptyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.EmptyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.CommonCorrectionRule, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddMerchantRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) AddMerchantRateCorrectionRule(ctx context.Context, in *currenciespb.CorrectionRule, opts ...client.CallOption) (*currenciespb.EmptyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.EmptyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.CorrectionRule, ...client.CallOption) *currenciespb.EmptyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.EmptyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.CorrectionRule, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyByDateCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) ExchangeCurrencyByDateCommon(ctx context.Context, in *currenciespb.ExchangeCurrencyByDateCommonRequest, opts ...client.CallOption) (*currenciespb.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyByDateCommonRequest, ...client.CallOption) *currenciespb.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.ExchangeCurrencyByDateCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyByDateForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) ExchangeCurrencyByDateForMerchant(ctx context.Context, in *currenciespb.ExchangeCurrencyByDateForMerchantRequest, opts ...client.CallOption) (*currenciespb.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyByDateForMerchantRequest, ...client.CallOption) *currenciespb.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.ExchangeCurrencyByDateForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyCurrentCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) ExchangeCurrencyCurrentCommon(ctx context.Context, in *currenciespb.ExchangeCurrencyCurrentCommonRequest, opts ...client.CallOption) (*currenciespb.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyCurrentCommonRequest, ...client.CallOption) *currenciespb.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.ExchangeCurrencyCurrentCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyCurrentForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) ExchangeCurrencyCurrentForMerchant(ctx context.Context, in *currenciespb.ExchangeCurrencyCurrentForMerchantRequest, opts ...client.CallOption) (*currenciespb.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyCurrentForMerchantRequest, ...client.CallOption) *currenciespb.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.ExchangeCurrencyCurrentForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountingCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetAccountingCurrencies(ctx context.Context, in *currenciespb.EmptyRequest, opts ...client.CallOption) (*currenciespb.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) *currenciespb.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetCommonRateCorrectionRule(ctx context.Context, in *currenciespb.CommonCorrectionRuleRequest, opts ...client.CallOption) (*currenciespb.CorrectionRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CorrectionRule
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.CommonCorrectionRuleRequest, ...client.CallOption) *currenciespb.CorrectionRule); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CorrectionRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.CommonCorrectionRuleRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrenciesPrecision provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetCurrenciesPrecision(ctx context.Context, in *currenciespb.EmptyRequest, opts ...client.CallOption) (*currenciespb.CurrenciesPrecisionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CurrenciesPrecisionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) *currenciespb.CurrenciesPrecisionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CurrenciesPrecisionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMerchantRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetMerchantRateCorrectionRule(ctx context.Context, in *currenciespb.MerchantCorrectionRuleRequest, opts ...client.CallOption) (*currenciespb.CorrectionRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CorrectionRule
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.MerchantCorrectionRuleRequest, ...client.CallOption) *currenciespb.CorrectionRule); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CorrectionRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.MerchantCorrectionRuleRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPriceCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetPriceCurrencies(ctx context.Context, in *currenciespb.EmptyRequest, opts ...client.CallOption) (*currenciespb.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) *currenciespb.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateByDateCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetRateByDateCommon(ctx context.Context, in *currenciespb.GetRateByDateCommonRequest, opts ...client.CallOption) (*currenciespb.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateByDateCommonRequest, ...client.CallOption) *currenciespb.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.GetRateByDateCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateByDateForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetRateByDateForMerchant(ctx context.Context, in *currenciespb.GetRateByDateForMerchantRequest, opts ...client.CallOption) (*currenciespb.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateByDateForMerchantRequest, ...client.CallOption) *currenciespb.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.GetRateByDateForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateCurrentCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetRateCurrentCommon(ctx context.Context, in *currenciespb.GetRateCurrentCommonRequest, opts ...client.CallOption) (*currenciespb.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateCurrentCommonRequest, ...client.CallOption) *currenciespb.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.GetRateCurrentCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateCurrentForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetRateCurrentForMerchant(ctx context.Context, in *currenciespb.GetRateCurrentForMerchantRequest, opts ...client.CallOption) (*currenciespb.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateCurrentForMerchantRequest, ...client.CallOption) *currenciespb.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.GetRateCurrentForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSettlementCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetSettlementCurrencies(ctx context.Context, in *currenciespb.EmptyRequest, opts ...client.CallOption) (*currenciespb.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) *currenciespb.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSupportedCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetSupportedCurrencies(ctx context.Context, in *currenciespb.EmptyRequest, opts ...client.CallOption) (*currenciespb.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) *currenciespb.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVatCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyRatesService) GetVatCurrencies(ctx context.Context, in *currenciespb.EmptyRequest, opts ...client.CallOption) (*currenciespb.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currenciespb.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) *currenciespb.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currenciespb.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currenciespb.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
