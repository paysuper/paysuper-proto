package reporterpb

const (
	ServiceName    = "p1payreporter"
	ServiceVersion = "latest"

	ReportTypeTransactions                  = "transactions"
	ReportTypeVat                           = "vat"
	ReportTypeVatTransactions               = "vat_transactions"
	ReportTypeRoyalty                       = "royalty"
	ReportTypeRoyaltyTransactions           = "royalty_transactions"
	ReportTypeRoyaltyAccountant             = "royalty_accountant"
	ReportTypeRoyaltyTransactionsAccountant = "royalty_transactions_accountant"
	ReportTypePayout                        = "payout"
	ReportTypePayoutFinance                 = "payout_finance"
	ReportTypeAgreement                     = "agreement"

	OutputExtensionXlsx = "xlsx"
	OutputExtensionCsv  = "csv"
	OutputExtensionPdf  = "pdf"

	ParamsFieldId            = "id"
	ParamsFieldCountry       = "country"
	ParamsFieldStatus        = "status"
	ParamsFieldPaymentMethod = "payment_method"
	ParamsFieldDateFrom      = "date_from"
	ParamsFieldDateTo        = "date_to"
	ParamsFieldMerchantId    = "merchant_id"

	RequestParameterAgreementNumber                             = "number"
	RequestParameterAgreementLegalName                          = "legal_name"
	RequestParameterAgreementAddress                            = "address"
	RequestParameterAgreementRegistrationNumber                 = "registration_number"
	RequestParameterAgreementPayoutCost                         = "payout_cost"
	RequestParameterAgreementMinimalPayoutLimit                 = "minimal_payout_limit"
	RequestParameterAgreementPayoutCurrency                     = "payout_currency"
	RequestParameterAgreementPSRate                             = "ps_rate"
	RequestParameterAgreementHomeRegion                         = "home_region"
	RequestParameterAgreementMerchantAuthorizedName             = "merchant_authorized_name"
	RequestParameterAgreementMerchantAuthorizedPosition         = "merchant_authorized_position"
	RequestParameterAgreementOperatingCompanyLegalName          = "oc_name"
	RequestParameterAgreementOperatingCompanyAddress            = "oc_address"
	RequestParameterAgreementOperatingCompanyRegistrationNumber = "oc_registration_number"
	RequestParameterAgreementOperatingCompanyAuthorizedName     = "oc_authorized_name"
	RequestParameterAgreementOperatingCompanyAuthorizedPosition = "oc_authorized_position"

	FileMask          = "report_%s_%s.%s"
	FileMaskAgreement = "License Agreement_%s_#%s.%s"
)
