package billingpb

import (
	"errors"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
)

const (
	ServiceName    = "p1paybilling"
	ServiceVersion = "latest"

	CardPayPaymentResponseStatusInProgress = "IN_PROGRESS"
	CardPayPaymentResponseStatusPending    = "PENDING"
	CardPayPaymentResponseStatusRefunded   = "REFUNDED"
	CardPayPaymentResponseStatusDeclined   = "DECLINED"
	CardPayPaymentResponseStatusAuthorized = "AUTHORIZED"
	CardPayPaymentResponseStatusCompleted  = "COMPLETED"
	CardPayPaymentResponseStatusCancelled  = "CANCELLED"

	StatusOK              = int32(0)
	StatusErrorValidation = int32(1)
	StatusErrorSystem     = int32(2)
	StatusTemporary       = int32(4)

	ResponseStatusOk          = int32(200)
	ResponseStatusNotModified = int32(304)
	ResponseStatusBadData     = int32(400)
	ResponseStatusNotFound    = int32(404)
	ResponseStatusForbidden   = int32(403)
	ResponseStatusGone        = int32(410)
	ResponseStatusSystemError = int32(500)
	ResponseStatusTemporary   = int32(410)

	OrderTypeOrder   = "order"
	OrderTypeRefund  = "refund"
	OrderTypePayment = "payment"

	VatReportStatusThreshold = "threshold"
	VatReportStatusExpired   = "expired"
	VatReportStatusPending   = "pending"
	VatReportStatusNeedToPay = "need_to_pay"
	VatReportStatusPaid      = "paid"
	VatReportStatusOverdue   = "overdue"
	VatReportStatusCanceled  = "canceled"

	MerchantStatusDraft             = int32(0)
	MerchantStatusAgreementSigning  = int32(3)
	MerchantStatusAgreementSigned   = int32(4)
	MerchantStatusDeleted           = int32(5)
	MerchantStatusRejected          = int32(6)
	MerchantStatusPending           = int32(7)
	MerchantStatusAccepted          = int32(8)
	MerchantStatusOpCompanySelected = int32(9)
	MerchantStatusKycStarted        = int32(10)
	MerchantStatusSuspend           = int32(11)

	ProjectStatusDraft         = int32(0)
	ProjectStatusTestCompleted = int32(1)
	ProjectStatusTestFailed    = int32(2)
	ProjectStatusInProduction  = int32(3)
	ProjectStatusDeleted       = int32(4)

	ProjectCallbackProtocolEmpty   = "empty"
	ProjectCallbackProtocolDefault = "default"

	RoyaltyReportStatusPending        = "pending"
	RoyaltyReportStatusAccepted       = "accepted"
	RoyaltyReportStatusCanceled       = "canceled"
	RoyaltyReportStatusDispute        = "dispute"
	RoyaltyReportStatusWaitForPayment = "waiting_payment"
	RoyaltyReportStatusPaid           = "paid"

	SignerTypeMerchant = int32(0)
	SignerTypePs       = int32(1)

	MccCodeLowRisk  = "5816"
	MccCodeHighRisk = "5967"

	RoleMerchantOwner      = "merchant_owner"
	RoleMerchantDeveloper  = "merchant_developer"
	RoleMerchantAccounting = "merchant_accounting"
	RoleMerchantSupport    = "merchant_support"
	RoleMerchantViewOnly   = "merchant_view_only"
	RoleSystemAdmin        = "system_admin"
	RoleSystemRiskManager  = "system_risk_manager"
	RoleSystemFinancial    = "system_financial"
	RoleSystemSupport      = "system_support"
	RoleSystemViewOnly     = "system_view_only"

	PaymentCreateFieldOrderId         = "order_id"
	PaymentCreateFieldPaymentMethodId = "payment_method_id"
	PaymentCreateFieldEmail           = "email"
	PaymentCreateFieldPan             = "pan"
	PaymentCreateFieldCvv             = "cvv"
	PaymentCreateFieldMonth           = "month"
	PaymentCreateFieldYear            = "year"
	PaymentCreateFieldHolder          = "card_holder"
	PaymentCreateFieldEWallet         = "ewallet"
	PaymentCreateFieldCrypto          = "address"
	PaymentCreateFieldStoreData       = "store_data"
	PaymentCreateFieldRecurringId     = "recurring_id"
	PaymentCreateFieldStoredCardId    = "stored_card_id"
	PaymentCreateFieldUserCountry     = "country"
	PaymentCreateFieldUserCity        = "city"
	PaymentCreateFieldUserState       = "state"
	PaymentCreateFieldUserZip         = "zip"

	TxnParamsFieldBankCardEmissionCountry = "emission_country"
	TxnParamsFieldBankCardToken           = "token"
	TxnParamsFieldBankCardIs3DS           = "is_3ds"
	TxnParamsFieldBankCardRrn             = "rrn"
	TxnParamsFieldDeclineCode             = "decline_code"
	TxnParamsFieldDeclineReason           = "decline_reason"
	TxnParamsFieldCryptoTransactionId     = "transaction_id"
	TxnParamsFieldCryptoAmount            = "amount_crypto"
	TxnParamsFieldCryptoCurrency          = "currency_crypto"

	PaymentCreateBankCardFieldBrand                = "card_brand"
	PaymentCreateBankCardFieldType                 = "card_type"
	PaymentCreateBankCardFieldCategory             = "card_category"
	PaymentCreateBankCardFieldIssuerName           = "bank_issuer_name"
	PaymentCreateBankCardFieldIssuerCountry        = "bank_issuer_country"
	PaymentCreateBankCardFieldIssuerCountryIsoCode = "bank_issuer_country_iso_code"

	CardPayDeclineCodeSystemMalfunction                = "01"
	CardPayDeclineCodeCancelledByCustomer              = "02"
	CardPayDeclineCodeDeclinedByAntiFraud              = "03"
	CardPayDeclineCodeDeclinedBy3DSecure               = "04"
	CardPayDeclineCodeOnly3DSecureTransactionsAllowed  = "05"
	CardPayDeclineCode3DSecureAvailabilityIsUnknown    = "06"
	CardPayDeclineCodeLimitReached                     = "07"
	CardPayDeclineCodeRequestedOperationIsNotSupported = "08"
	CardPayDeclineCodeDeclinedByBankWithoutReason      = "10"
	CardPayDeclineCodeCommonDeclineByBank              = "11"
	CardPayDeclineCodeInsufficientFunds                = "13"
	CardPayDeclineCodeCardLimitReached                 = "14"
	CardPayDeclineCodeIncorrectCardData                = "15"
	CardPayDeclineCodeDeclinedByBankAntiFraud          = "16"
	CardPayDeclineCodeBanksMalfunction                 = "17"
	CardPayDeclineCodeConnectionProblem                = "18"
	CardPayDeclineCodeNoPaymentWasReceived             = "21"
	CardPayDeclineCodeWrongPaymentWasReceived          = "22"
	CardPayDeclineCodeConfirmationsPaymentTimeout      = "23"

	PaySuperDeclineCodeSystemMalfunction                = "ps000001"
	PaySuperDeclineCodeCancelledByCustomer              = "ps000002"
	PaySuperDeclineCodeDeclinedByAntiFraud              = "ps000003"
	PaySuperDeclineCodeDeclinedBy3DSecure               = "ps000004"
	PaySuperDeclineCodeOnly3DSecureTransactionsAllowed  = "ps000005"
	PaySuperDeclineCode3DSecureAvailabilityIsUnknown    = "ps000006"
	PaySuperDeclineCodeLimitReached                     = "ps000007"
	PaySuperDeclineCodeRequestedOperationIsNotSupported = "ps000008"
	PaySuperDeclineCodeDeclinedByBankWithoutReason      = "ps000009"
	PaySuperDeclineCodeCommonDeclineByBank              = "ps000010"
	PaySuperDeclineCodeInsufficientFunds                = "ps000011"
	PaySuperDeclineCodeCardLimitReached                 = "ps000012"
	PaySuperDeclineCodeIncorrectCardData                = "ps000013"
	PaySuperDeclineCodeDeclinedByBankAntiFraud          = "ps000014"
	PaySuperDeclineCodeBanksMalfunction                 = "ps000015"
	PaySuperDeclineCodeConnectionProblem                = "ps000016"
	PaySuperDeclineCodeNoPaymentWasReceived             = "ps000017"
	PaySuperDeclineCodeWrongPaymentWasReceived          = "ps000018"
	PaySuperDeclineCodeConfirmationsPaymentTimeout      = "ps000019"

	ErrorGrpcServiceCallFailed = "gRPC call failed"
	ErrorFieldService          = "service"
	ErrorFieldMethod           = "method"
	ErrorFieldRequest          = "request"
	ErrorFieldStatus           = "status"
	ErrorFieldResponse         = "response"

	ErrorInvalidObjectId = "invalid bson object id"
	ErrorUnableRound     = "Unable to round amount to selected precision"
	ErrorFieldKey        = "key"
	ErrorFieldValue      = "value"

	TariffRegionRussiaAndCis = "russia_and_cis"
	TariffRegionEurope       = "europe"
	TariffRegionAsia         = "asia"
	TariffRegionLatAm        = "latin_america"
	TariffRegionWorldwide    = "worldwide"

	VatPayerBuyer  = "buyer"
	VatPayerSeller = "seller"
	VatPayerNobody = "nobody"

	// MerchantId_UserId
	CasbinMerchantUserMask = "%s_%s"

	RoleTypeMerchant = "merchant"
	RoleTypeSystem   = "system"

	PaymentSystemHandlerCardPay = "cardpay"

	PaylinkUrlDefaultMask       = "/paylink/%s"
	DatabaseRequestDefaultLimit = int64(100)

	RoyaltyReportChangeSourceAuto     = "auto"
	RoyaltyReportChangeSourceMerchant = "merchant"
	RoyaltyReportChangeSourceAdmin    = "admin"

	ProjectRedirectModeAny        = "any"
	ProjectRedirectModeDisable    = "disable"
	ProjectRedirectModeSuccessful = "successful"
	ProjectRedirectModeFail       = "fail"
	ProjectRedirectUsageAny       = "any"

	TestCaseCorrectPayment   = "correct_payment"
	TestCaseNonExistingUser  = "non_existing_user"
	TestCaseExistingUser     = "existing_user"
	TestCaseIncorrectPayment = "invalid_signature"

	PaymentSystemGroupAliasBankCard = "BANKCARD"
	PaymentSystemGroupAliasWebMoney = "WEBMONEY"
	PaymentSystemGroupAliasQiwi     = "QIWI"
	PaymentSystemGroupAliasNeteller = "NETELLER"
	PaymentSystemGroupAliasAlipay   = "ALIPAY"
	PaymentSystemGroupAliasBitcoin  = "BITCOIN"

	VirtualCurrencyPriceGroup = "virtual"

	orderBankCardBrandNotFound = "brand for bank card not found"
	orderPaymentMethodNotSet   = "payment method not set"

	FilterDateFormat     = "2006-01-02"
	FilterDatetimeFormat = "2006-01-02T15:04:05"

	AmountPrecisionDefault = 2

	UserLocaleEn = "en"
	UserLocaleRu = "ru"

	RecurringPeriodMinute = "minute"
	RecurringPeriodDay    = "day"
	RecurringPeriodWeek   = "week"
	RecurringPeriodMonth  = "month"
	RecurringPeriodYear   = "year"

	RecurringSubscriptionStatusActive      = "active"
	RecurringSubscriptionStatusCanceled    = "canceled"
	RecurringSubscriptionStatusNonRenewing = "no_renewing"
)

var (
	DeclineCodeMap = map[string]string{
		CardPayDeclineCodeSystemMalfunction:                PaySuperDeclineCodeSystemMalfunction,
		CardPayDeclineCodeCancelledByCustomer:              PaySuperDeclineCodeCancelledByCustomer,
		CardPayDeclineCodeDeclinedByAntiFraud:              PaySuperDeclineCodeDeclinedByAntiFraud,
		CardPayDeclineCodeDeclinedBy3DSecure:               PaySuperDeclineCodeDeclinedBy3DSecure,
		CardPayDeclineCodeOnly3DSecureTransactionsAllowed:  PaySuperDeclineCodeOnly3DSecureTransactionsAllowed,
		CardPayDeclineCode3DSecureAvailabilityIsUnknown:    PaySuperDeclineCode3DSecureAvailabilityIsUnknown,
		CardPayDeclineCodeLimitReached:                     PaySuperDeclineCodeLimitReached,
		CardPayDeclineCodeRequestedOperationIsNotSupported: PaySuperDeclineCodeRequestedOperationIsNotSupported,
		CardPayDeclineCodeDeclinedByBankWithoutReason:      PaySuperDeclineCodeDeclinedByBankWithoutReason,
		CardPayDeclineCodeCommonDeclineByBank:              PaySuperDeclineCodeCommonDeclineByBank,
		CardPayDeclineCodeInsufficientFunds:                PaySuperDeclineCodeInsufficientFunds,
		CardPayDeclineCodeCardLimitReached:                 PaySuperDeclineCodeCardLimitReached,
		CardPayDeclineCodeIncorrectCardData:                PaySuperDeclineCodeIncorrectCardData,
		CardPayDeclineCodeDeclinedByBankAntiFraud:          PaySuperDeclineCodeDeclinedByBankAntiFraud,
		CardPayDeclineCodeBanksMalfunction:                 PaySuperDeclineCodeBanksMalfunction,
		CardPayDeclineCodeConnectionProblem:                PaySuperDeclineCodeConnectionProblem,
		CardPayDeclineCodeNoPaymentWasReceived:             PaySuperDeclineCodeNoPaymentWasReceived,
		CardPayDeclineCodeWrongPaymentWasReceived:          PaySuperDeclineCodeWrongPaymentWasReceived,
		CardPayDeclineCodeConfirmationsPaymentTimeout:      PaySuperDeclineCodeConfirmationsPaymentTimeout,
	}

	orderRefundAllowedStatuses = map[int32]bool{
		recurringpb.OrderStatusPaymentSystemComplete: true,
		recurringpb.OrderStatusProjectInProgress:     true,
		recurringpb.OrderStatusProjectComplete:       true,
		recurringpb.OrderStatusProjectPending:        true,
	}

	orderStatusPublicMapping = map[int32]string{
		recurringpb.OrderStatusNew:                         recurringpb.OrderPublicStatusCreated,
		recurringpb.OrderStatusPaymentSystemCreate:         recurringpb.OrderPublicStatusCreated,
		recurringpb.OrderStatusPaymentSystemCanceled:       recurringpb.OrderPublicStatusCanceled,
		recurringpb.OrderStatusPaymentSystemRejectOnCreate: recurringpb.OrderPublicStatusRejected,
		recurringpb.OrderStatusPaymentSystemReject:         recurringpb.OrderPublicStatusRejected,
		recurringpb.OrderStatusProjectReject:               recurringpb.OrderPublicStatusRejected,
		recurringpb.OrderStatusPaymentSystemDeclined:       recurringpb.OrderPublicStatusRejected,
		recurringpb.OrderStatusPaymentSystemComplete:       recurringpb.OrderPublicStatusProcessed,
		recurringpb.OrderStatusProjectComplete:             recurringpb.OrderPublicStatusProcessed,
		recurringpb.OrderStatusRefund:                      recurringpb.OrderPublicStatusRefunded,
		recurringpb.OrderStatusChargeback:                  recurringpb.OrderPublicStatusChargeback,
		recurringpb.OrderStatusItemReplaced:                recurringpb.OrderPublicStatusProcessed,
		recurringpb.OrderStatusProjectInProgress:           recurringpb.OrderPublicStatusPending,
		recurringpb.OrderStatusPaymentSystemAuthorized:     recurringpb.OrderPublicStatusAuthorized,
	}

	HomeRegions = map[string]string{
		TariffRegionAsia:         "Asia",
		TariffRegionEurope:       "Europe",
		TariffRegionLatAm:        "Latin America",
		TariffRegionRussiaAndCis: "Russia & CIS",
		TariffRegionWorldwide:    "Worldwide",
	}

	ProductNoPriceInCurrencyError = errors.New("No product price in requested currency")

	productNoPriceInCurrency           = "no price in currency %s"
	productNoNameInLanguage            = "no name in language %s"
	productNoDescriptionInLanguage     = "no description in language %s"
	productNoLongDescriptionInLanguage = "no long description in language %s"

	UserLocaleAllowed = map[string]bool{
		UserLocaleEn: true,
		UserLocaleRu: true,
	}
)
