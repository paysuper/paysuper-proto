// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	currenciespb "github.com/paysuper/paysuper-proto/go/currenciespb"
	mock "github.com/stretchr/testify/mock"
)

// CurrencyRatesServiceHandler is an autogenerated mock type for the CurrencyRatesServiceHandler type
type CurrencyRatesServiceHandler struct {
	mock.Mock
}

// AddCommonRateCorrectionRule provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) AddCommonRateCorrectionRule(_a0 context.Context, _a1 *currenciespb.CommonCorrectionRule, _a2 *currenciespb.EmptyResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.CommonCorrectionRule, *currenciespb.EmptyResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddMerchantRateCorrectionRule provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) AddMerchantRateCorrectionRule(_a0 context.Context, _a1 *currenciespb.CorrectionRule, _a2 *currenciespb.EmptyResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.CorrectionRule, *currenciespb.EmptyResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExchangeCurrencyByDateCommon provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) ExchangeCurrencyByDateCommon(_a0 context.Context, _a1 *currenciespb.ExchangeCurrencyByDateCommonRequest, _a2 *currenciespb.ExchangeCurrencyResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyByDateCommonRequest, *currenciespb.ExchangeCurrencyResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExchangeCurrencyByDateForMerchant provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) ExchangeCurrencyByDateForMerchant(_a0 context.Context, _a1 *currenciespb.ExchangeCurrencyByDateForMerchantRequest, _a2 *currenciespb.ExchangeCurrencyResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyByDateForMerchantRequest, *currenciespb.ExchangeCurrencyResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExchangeCurrencyCurrentCommon provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) ExchangeCurrencyCurrentCommon(_a0 context.Context, _a1 *currenciespb.ExchangeCurrencyCurrentCommonRequest, _a2 *currenciespb.ExchangeCurrencyResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyCurrentCommonRequest, *currenciespb.ExchangeCurrencyResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExchangeCurrencyCurrentForMerchant provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) ExchangeCurrencyCurrentForMerchant(_a0 context.Context, _a1 *currenciespb.ExchangeCurrencyCurrentForMerchantRequest, _a2 *currenciespb.ExchangeCurrencyResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.ExchangeCurrencyCurrentForMerchantRequest, *currenciespb.ExchangeCurrencyResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccountingCurrencies provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetAccountingCurrencies(_a0 context.Context, _a1 *currenciespb.EmptyRequest, _a2 *currenciespb.CurrenciesList) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, *currenciespb.CurrenciesList) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommonRateCorrectionRule provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetCommonRateCorrectionRule(_a0 context.Context, _a1 *currenciespb.CommonCorrectionRuleRequest, _a2 *currenciespb.CorrectionRule) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.CommonCorrectionRuleRequest, *currenciespb.CorrectionRule) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCurrenciesPrecision provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetCurrenciesPrecision(_a0 context.Context, _a1 *currenciespb.EmptyRequest, _a2 *currenciespb.CurrenciesPrecisionResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, *currenciespb.CurrenciesPrecisionResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMerchantRateCorrectionRule provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetMerchantRateCorrectionRule(_a0 context.Context, _a1 *currenciespb.MerchantCorrectionRuleRequest, _a2 *currenciespb.CorrectionRule) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.MerchantCorrectionRuleRequest, *currenciespb.CorrectionRule) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPriceCurrencies provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetPriceCurrencies(_a0 context.Context, _a1 *currenciespb.EmptyRequest, _a2 *currenciespb.CurrenciesList) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, *currenciespb.CurrenciesList) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRateByDateCommon provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetRateByDateCommon(_a0 context.Context, _a1 *currenciespb.GetRateByDateCommonRequest, _a2 *currenciespb.RateData) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateByDateCommonRequest, *currenciespb.RateData) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRateByDateForMerchant provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetRateByDateForMerchant(_a0 context.Context, _a1 *currenciespb.GetRateByDateForMerchantRequest, _a2 *currenciespb.RateData) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateByDateForMerchantRequest, *currenciespb.RateData) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRateCurrentCommon provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetRateCurrentCommon(_a0 context.Context, _a1 *currenciespb.GetRateCurrentCommonRequest, _a2 *currenciespb.RateData) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateCurrentCommonRequest, *currenciespb.RateData) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRateCurrentForMerchant provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetRateCurrentForMerchant(_a0 context.Context, _a1 *currenciespb.GetRateCurrentForMerchantRequest, _a2 *currenciespb.RateData) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.GetRateCurrentForMerchantRequest, *currenciespb.RateData) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSettlementCurrencies provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetSettlementCurrencies(_a0 context.Context, _a1 *currenciespb.EmptyRequest, _a2 *currenciespb.CurrenciesList) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, *currenciespb.CurrenciesList) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSupportedCurrencies provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetSupportedCurrencies(_a0 context.Context, _a1 *currenciespb.EmptyRequest, _a2 *currenciespb.CurrenciesList) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, *currenciespb.CurrenciesList) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetVatCurrencies provides a mock function with given fields: _a0, _a1, _a2
func (_m *CurrencyRatesServiceHandler) GetVatCurrencies(_a0 context.Context, _a1 *currenciespb.EmptyRequest, _a2 *currenciespb.CurrenciesList) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currenciespb.EmptyRequest, *currenciespb.CurrenciesList) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
