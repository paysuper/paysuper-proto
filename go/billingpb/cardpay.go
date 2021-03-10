package billingpb

var (
	cardPayPaymentCallbackAllowedStatuses = map[string]bool{
		CardPayPaymentResponseStatusCompleted:  true,
		CardPayPaymentResponseStatusDeclined:   true,
		CardPayPaymentResponseStatusCancelled:  true,
		CardPayPaymentResponseStatusAuthorized: true,
	}
)

func (m *CardPayPaymentCallback) IsPaymentAllowedStatus() bool {
	v, ok := cardPayPaymentCallbackAllowedStatuses[m.GetStatus()]

	return ok && v == true
}

func (m *CardPayRefundCallback) IsRefundAllowedStatus() bool {
	v, ok := cardPayPaymentCallbackAllowedStatuses[m.RefundData.Status]

	return ok && v == true
}

func (m *CardPayPaymentCallback) GetBankCardTxnParams() map[string]string {
	params := make(map[string]string)

	params[PaymentCreateFieldPan] = m.CardAccount.MaskedPan
	params[PaymentCreateFieldHolder] = m.CardAccount.Holder
	params[TxnParamsFieldBankCardEmissionCountry] = m.CardAccount.IssuingCountryCode
	params[TxnParamsFieldBankCardToken] = m.CardAccount.Token
	params[TxnParamsFieldBankCardRrn] = m.GetRrn()
	params[TxnParamsFieldBankCardIs3DS] = m.GetIs3DS()

	if m.GetStatus() == CardPayPaymentResponseStatusDeclined {
		params[TxnParamsFieldDeclineCode] = m.GetDeclineCode()
		params[TxnParamsFieldDeclineReason] = m.GetDeclineReason()
	}

	return params
}

func (m *CardPayPaymentCallback) GetEWalletTxnParams() map[string]string {
	params := make(map[string]string)

	params[PaymentCreateFieldEWallet] = m.EwalletAccount.Id

	if m.GetStatus() == CardPayPaymentResponseStatusDeclined {
		params[TxnParamsFieldDeclineCode] = m.GetDeclineCode()
		params[TxnParamsFieldDeclineReason] = m.GetDeclineReason()
	}

	return params
}

func (m *CardPayPaymentCallback) GetCryptoCurrencyTxnParams() map[string]string {
	params := make(map[string]string)

	params[PaymentCreateFieldCrypto] = m.CryptocurrencyAccount.CryptoAddress
	params[TxnParamsFieldCryptoTransactionId] = m.CryptocurrencyAccount.CryptoTransactionId
	params[TxnParamsFieldCryptoAmount] = m.CryptocurrencyAccount.PrcAmount
	params[TxnParamsFieldCryptoCurrency] = m.CryptocurrencyAccount.PrcCurrency

	if m.GetStatus() == CardPayPaymentResponseStatusDeclined {
		params[TxnParamsFieldDeclineCode] = m.GetDeclineCode()
		params[TxnParamsFieldDeclineReason] = m.GetDeclineReason()
	}

	return params
}

func (m *CardPayPaymentCallback) IsRecurring() bool {
	return m.RecurringData != nil
}

func (m *CardPayPaymentCallback) IsRecurringWithFiling() bool {
	return m.RecurringData != nil && m.RecurringData.Filing != nil
}

func (m *CardPayPaymentCallback) GetId() string {
	if m.PaymentData != nil {
		return m.PaymentData.Id
	}

	return m.RecurringData.Id
}

func (m *CardPayPaymentCallback) GetAmount() float64 {
	if m.PaymentData != nil {
		return m.PaymentData.Amount
	}

	return m.RecurringData.Amount
}

func (m *CardPayPaymentCallback) GetCurrency() string {
	if m.PaymentData != nil {
		return m.PaymentData.Currency
	}

	return m.RecurringData.Currency
}

func (m *CardPayPaymentCallback) GetStatus() string {
	if m.PaymentData != nil {
		return m.PaymentData.Status
	}

	return m.RecurringData.Status
}

func (m *CardPayPaymentCallback) GetDeclineCode() string {
	if m.PaymentData != nil {
		return m.PaymentData.DeclineCode
	}

	return m.RecurringData.DeclineCode
}

func (m *CardPayPaymentCallback) GetDeclineReason() string {
	if m.PaymentData != nil {
		return m.PaymentData.DeclineReason
	}

	return m.RecurringData.DeclineReason
}

func (m *CardPayPaymentCallback) GetRrn() string {
	if m.PaymentData != nil {
		return m.PaymentData.Rrn
	}

	return m.RecurringData.Rrn
}

func (m *CardPayPaymentCallback) GetIs3DS() string {
	is3DS := false
	result := "0"

	if m.PaymentData != nil {
		is3DS = m.PaymentData.Is_3D
	} else {
		is3DS = m.RecurringData.Is_3D
	}

	if is3DS == true {
		result = "1"
	}

	return result
}

func (m *CardPayPaymentCallback) IsSuccess() bool {
	return m.RecurringData != nil &&
		(m.RecurringData.Status == CardPayPaymentResponseStatusCompleted ||
			m.RecurringData.Status == CardPayPaymentResponseStatusAuthorized)
}
