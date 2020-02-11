package billingpb

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	tools "github.com/paysuper/paysuper-tools/number"
	"time"
)

const (
	orderBankCardBrandNotFound = "brand for bank card not found"
	orderPaymentMethodNotSet   = "payment method not set"
)

func (m *Order) CanBeRecreated() bool {
	if m.PrivateStatus != recurringpb.OrderStatusPaymentSystemRejectOnCreate &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemReject &&
		m.PrivateStatus != recurringpb.OrderStatusProjectReject &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemDeclined &&
		m.PrivateStatus != recurringpb.OrderStatusNew &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemCreate &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemCanceled {

		return false
	}

	return true
}

func (m *Order) GetMerchantId() string {
	return m.Project.MerchantId
}

func (m *Order) GetPaymentMethodId() string {
	return m.PaymentMethod.Id
}

func (m *Order) IsDeclined() bool {
	return m.PrivateStatus == recurringpb.OrderStatusPaymentSystemDeclined || m.PrivateStatus == recurringpb.OrderStatusPaymentSystemCanceled
}

func (m *Order) GetDeclineReason() string {
	reason, ok := m.PaymentMethodTxnParams[TxnParamsFieldDeclineReason]

	if !ok {
		return ""
	}

	return reason
}

func (m *Order) GetPrivateDeclineCode() string {
	code, ok := m.PaymentMethodTxnParams[TxnParamsFieldDeclineCode]

	if !ok {
		return ""
	}

	return code
}

func (m *Order) GetPublicDeclineCode() string {
	code := m.GetPrivateDeclineCode()

	if code == "" {
		return ""
	}

	code, ok := DeclineCodeMap[code]

	if !ok {
		return ""
	}

	return code
}

func (m *Order) GetMerchantRoyaltyCurrency() string {
	if m.Project == nil {
		return ""
	}
	return m.Project.MerchantRoyaltyCurrency
}

func (m *Order) GetPaymentMethodName() string {
	return m.PaymentMethod.Name
}

func (m *Order) GetBankCardBrand() (string, error) {
	val, ok := m.PaymentRequisites[PaymentCreateBankCardFieldBrand]

	if !ok {
		return "", errors.New(orderBankCardBrandNotFound)
	}

	return val, nil
}

func (m *Order) GetCountry() string {
	if m.BillingAddress != nil && m.BillingAddress.Country != "" {
		return m.BillingAddress.Country
	}
	if m.User != nil && m.User.Address != nil && m.User.Address.Country != "" {
		return m.User.Address.Country
	}
	return ""
}

func (m *Order) GetPostalCode() string {
	if m.BillingAddress != nil && m.BillingAddress.PostalCode != "" {
		return m.BillingAddress.PostalCode
	}
	if m.User != nil && m.User.Address != nil && m.User.Address.PostalCode != "" {
		return m.User.Address.PostalCode
	}
	return ""
}

func (m *Order) HasEndedStatus() bool {
	return m.PrivateStatus == recurringpb.OrderStatusPaymentSystemReject || m.PrivateStatus == recurringpb.OrderStatusProjectComplete ||
		m.PrivateStatus == recurringpb.OrderStatusProjectReject || m.PrivateStatus == recurringpb.OrderStatusRefund ||
		m.PrivateStatus == recurringpb.OrderStatusChargeback
}

func (m *Order) RefundAllowed() bool {
	v, ok := orderRefundAllowedStatuses[m.PrivateStatus]

	return ok && v == true
}

func (m *Order) FormInputTimeIsEnded() bool {
	t, err := ptypes.Timestamp(m.ExpireDateToFormInput)

	return err != nil || t.Before(time.Now())
}

func (m *Order) GetProjectId() string {
	return m.Project.Id
}

func (m *Order) GetPublicStatus() string {
	st, ok := orderStatusPublicMapping[m.PrivateStatus]
	if !ok {
		return recurringpb.OrderPublicStatusPending
	}
	return st
}

func (m *Order) GetReceiptUserEmail() string {
	if m.User != nil {
		return m.User.Email
	}
	return ""
}

func (m *Order) GetReceiptUserPhone() string {
	if m.User != nil {
		return m.User.Phone
	}
	return ""
}

func (m *Order) GetState() string {
	if m.BillingAddress != nil && m.BillingAddress.State != "" {
		return m.BillingAddress.State
	}
	if m.User != nil && m.User.Address != nil && m.User.Address.State != "" {
		return m.User.Address.State
	}
	return ""
}

func (m *Order) SetNotificationStatus(key string, val bool) {
	if m.IsNotificationsSent == nil {
		m.IsNotificationsSent = make(map[string]bool)
	}
	m.IsNotificationsSent[key] = val
}

func (m *Order) GetNotificationStatus(key string) bool {
	if m.IsNotificationsSent == nil {
		return false
	}
	val, ok := m.IsNotificationsSent[key]
	if !ok {
		return false
	}
	return val
}

func (m *Order) GetCostPaymentMethodName() (string, error) {
	if m.PaymentMethod == nil {
		return "", errors.New(orderPaymentMethodNotSet)
	}
	if m.PaymentMethod.IsBankCard() {
		return m.GetBankCardBrand()
	}
	return m.GetPaymentMethodName(), nil
}

func (m *Order) GetPaymentSystemApiUrl() string {
	return m.PaymentMethod.Params.ApiUrl
}

func (m *Order) GetPaymentFormDataChangeResult() *PaymentFormDataChangeResponseItem {
	item := &PaymentFormDataChangeResponseItem{
		UserAddressDataRequired: m.UserAddressDataRequired,
		UserIpData: &UserIpData{
			Country: m.User.Address.Country,
			City:    m.User.Address.City,
			Zip:     m.User.Address.PostalCode,
		},
		Brand:                  "",
		HasVat:                 m.Tax.Rate > 0,
		VatRate:                tools.ToPrecise(m.Tax.Rate),
		Vat:                    m.Tax.Amount,
		VatInChargeCurrency:    tools.FormatAmount(m.GetTaxAmountInChargeCurrency()),
		ChargeAmount:           m.ChargeAmount,
		ChargeCurrency:         m.ChargeCurrency,
		Currency:               m.Currency,
		Amount:                 m.OrderAmount,
		TotalAmount:            m.TotalPaymentAmount,
		Items:                  m.Items,
		CountryPaymentsAllowed: true,
		CountryChangeAllowed:   m.CountryChangeAllowed(),
	}

	brand, err := m.GetBankCardBrand()
	if err == nil {
		item.Brand = brand
	}

	if m.CountryRestriction != nil {
		item.CountryPaymentsAllowed = m.CountryRestriction.PaymentsAllowed
	}

	return item
}

func (m *Order) GetTaxAmountInChargeCurrency() float64 {
	if m.Tax.Amount == 0 {
		return 0
	}
	if m.Currency == m.ChargeCurrency {
		return m.Tax.Amount
	}
	return tools.GetPercentPartFromAmount(m.ChargeAmount, m.Tax.Rate)
}

func (m *Order) IsDeclinedByCountry() bool {
	return m.PrivateStatus == recurringpb.OrderStatusPaymentSystemDeclined &&
		m.CountryRestriction != nil &&
		m.CountryRestriction.PaymentsAllowed == false &&
		m.CountryRestriction.ChangeAllowed == false
}

func (m *Order) CountryChangeAllowed() bool {
	return m.CountryRestriction == nil || m.CountryRestriction.ChangeAllowed == true
}

func (m *Order) NeedCallbackNotification() bool {
	status := m.GetPublicStatus()
	return status == recurringpb.OrderPublicStatusRefunded || status == recurringpb.OrderPublicStatusProcessed ||
		status == recurringpb.OrderPublicStatusChargeback || m.IsDeclined()
}
