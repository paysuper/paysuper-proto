package billingpb

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	tools "github.com/paysuper/paysuper-tools/number"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MgoId struct {
	Id primitive.ObjectID `bson:"_id"`
}

type MgoMultiLang struct {
	Lang  string `bson:"lang"`
	Value string `bson:"value"`
}

type MgoCommission struct {
	Id struct {
		PaymentMethodId primitive.ObjectID `bson:"pm_id"`
		ProjectId       primitive.ObjectID `bson:"project_id"`
	} `bson:"_id"`
	PaymentMethodCommission float64   `bson:"pm_commission"`
	PspCommission           float64   `bson:"psp_commission"`
	ToUserCommission        float64   `bson:"total_commission_to_user"`
	StartDate               time.Time `bson:"start_date"`
}

type MgoCommissionBilling struct {
	Id                      primitive.ObjectID `bson:"_id"`
	PaymentMethodId         primitive.ObjectID `bson:"pm_id"`
	ProjectId               primitive.ObjectID `bson:"project_id"`
	PaymentMethodCommission float64            `bson:"pm_commission"`
	PspCommission           float64            `bson:"psp_commission"`
	TotalCommissionToUser   float64            `bson:"total_commission_to_user"`
	StartDate               time.Time          `bson:"start_date"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}

type MgoOrderProject struct {
	Id                      primitive.ObjectID       `bson:"_id"`
	MerchantId              primitive.ObjectID       `bson:"merchant_id"`
	Name                    []*MgoMultiLang          `bson:"name"`
	UrlSuccess              string                   `bson:"url_success"`
	UrlFail                 string                   `bson:"url_fail"`
	NotifyEmails            []string                 `bson:"notify_emails"`
	SecretKey               string                   `bson:"secret_key"`
	SendNotifyEmail         bool                     `bson:"send_notify_email"`
	UrlCheckAccount         string                   `bson:"url_check_account"`
	UrlProcessPayment       string                   `bson:"url_process_payment"`
	CallbackProtocol        string                   `bson:"callback_protocol"`
	UrlChargebackPayment    string                   `bson:"url_chargeback_payment"`
	UrlCancelPayment        string                   `bson:"url_cancel_payment"`
	UrlFraudPayment         string                   `bson:"url_fraud_payment"`
	UrlRefundPayment        string                   `bson:"url_refund_payment"`
	Status                  int32                    `bson:"status"`
	MerchantRoyaltyCurrency string                   `bson:"merchant_royalty_currency"`
	RedirectSettings        *ProjectRedirectSettings `bson:"redirect_settings"`
}

type MgoOrderPaymentMethod struct {
	Id              primitive.ObjectID   `bson:"_id"`
	Name            string               `bson:"name"`
	Handler         string               `bson:"handler"`
	ExternalId      string               `bson:"external_id"`
	Params          *PaymentMethodParams `bson:"params"`
	PaymentSystemId primitive.ObjectID   `bson:"payment_system_id"`
	Group           string               `bson:"group_alias"`
	Saved           bool                 `bson:"saved"`
	Card            *PaymentMethodCard   `bson:"card,omitempty"`
	Wallet          *PaymentMethodWallet `bson:"wallet,omitempty"`
	CryptoCurrency  *PaymentMethodCrypto `bson:"crypto_currency,omitempty"`
	RefundAllowed   bool                 `bson:"refund_allowed"`
}

type MgoOrderNotificationRefund struct {
	Amount        float64 `bson:"amount"`
	Currency      string  `bson:"currency"`
	Reason        string  `bson:"reason"`
	Code          string  `bson:"code"`
	ReceiptNumber string  `bson:"receipt_number"`
	ReceiptUrl    string  `bson:"receipt_url"`
}

type MgoOrder struct {
	Id                          primitive.ObjectID             `bson:"_id"`
	Uuid                        string                         `bson:"uuid"`
	Transaction                 string                         `bson:"pm_order_id"`
	Object                      string                         `bson:"object"`
	Status                      string                         `bson:"status"`
	PrivateStatus               int32                          `bson:"private_status"`
	Description                 string                         `bson:"description"`
	CreatedAt                   time.Time                      `bson:"created_at"`
	UpdatedAt                   time.Time                      `bson:"updated_at"`
	CanceledAt                  time.Time                      `bson:"canceled_at"`
	Canceled                    bool                           `bson:"canceled"`
	Cancellation                *OrderNotificationCancellation `bson:"cancellation"`
	Refunded                    bool                           `bson:"refunded"`
	RefundedAt                  time.Time                      `bson:"refunded_at"`
	ReceiptEmail                string                         `bson:"receipt_email"`
	ReceiptPhone                string                         `bson:"receipt_phone"`
	ReceiptNumber               string                         `bson:"receipt_number"`
	ReceiptUrl                  string                         `bson:"receipt_url"`
	AgreementVersion            string                         `bson:"agreement_version"`
	AgreementAccepted           bool                           `bson:"agreement_accepted"`
	NotifySale                  bool                           `bson:"notify_sale"`
	NotifySaleEmail             string                         `bson:"notify_sale_email"`
	Issuer                      *OrderIssuer                   `bson:"issuer"`
	TotalPaymentAmount          float64                        `bson:"total_payment_amount"`
	Currency                    string                         `bson:"currency"`
	User                        *OrderUser                     `bson:"user"`
	BillingAddress              *OrderBillingAddress           `bson:"billing_address"`
	Tax                         *OrderTax                      `bson:"tax"`
	PaymentMethod               *MgoOrderPaymentMethod         `bson:"payment_method"`
	Items                       []*MgoOrderItem                `bson:"items"`
	Refund                      *MgoOrderNotificationRefund    `bson:"refund"`
	Metadata                    map[string]string              `bson:"metadata"`
	PrivateMetadata             map[string]string              `bson:"private_metadata"`
	Project                     *MgoOrderProject               `bson:"project"`
	ProjectOrderId              string                         `bson:"project_order_id"`
	ProjectAccount              string                         `bson:"project_account"`
	ProjectLastRequestedAt      time.Time                      `bson:"project_last_requested_at"`
	ProjectParams               map[string]string              `bson:"project_params"`
	PaymentMethodOrderClosedAt  time.Time                      `bson:"pm_order_close_date"`
	IsJsonRequest               bool                           `bson:"created_by_json"`
	OrderAmount                 float64                        `bson:"private_amount"`
	PaymentMethodPayerAccount   string                         `bson:"pm_account"`
	PaymentMethodTxnParams      map[string]string              `bson:"pm_txn_params"`
	PaymentRequisites           map[string]string              `bson:"payment_requisites"`
	ExpireDateToFormInput       time.Time                      `bson:"expire_date_to_form_input"`
	UserAddressDataRequired     bool                           `bson:"user_address_data_required"`
	Products                    []string                       `bson:"products"`
	IsNotificationsSent         map[string]bool                `bson:"is_notifications_sent"`
	CountryRestriction          *CountryRestriction            `bson:"country_restriction"`
	ParentOrder                 *ParentOrder                   `bson:"parent_order"`
	ParentPaymentAt             time.Time                      `bson:"parent_payment_at"`
	Type                        string                         `bson:"type"`
	IsVatDeduction              bool                           `bson:"is_vat_deduction"`
	CountryCode                 string                         `bson:"country_code"`
	PlatformId                  string                         `bson:"platform_id"`
	ProductType                 string                         `bson:"product_type"`
	Keys                        []string                       `bson:"keys"`
	IsKeyProductNotified        bool                           `bson:"is_key_product_notified"`
	ReceiptId                   string                         `bson:"receipt_id"`
	IsBuyForVirtualCurrency     bool                           `bson:"is_buy_for_virtual_currency"`
	MccCode                     string                         `bson:"mcc_code"`
	OperatingCompanyId          string                         `bson:"operating_company_id"`
	IsHighRisk                  bool                           `bson:"is_high_risk"`
	ChargeCurrency              string                         `bson:"charge_currency"`
	ChargeAmount                float64                        `bson:"charge_amount"`
	PaymentIpCountry            string                         `bson:"payment_ip_country"`
	IsIpCountryMismatchBin      bool                           `bson:"is_ip_country_mismatch_bin"`
	BillingCountryChangedByUser bool                           `bson:"billing_country_changed_by_user"`
	IsRefundAllowed             bool                           `bson:"is_refund_allowed"`
	VatPayer                    string                         `bson:"vat_payer"`
	IsProduction                bool                           `bson:"is_production"`
	FormMode                    string                         `bson:"form_mode"`
}

type MgoOrderItem struct {
	Id          primitive.ObjectID `bson:"_id"`
	Object      string             `bson:"object"`
	Sku         string             `bson:"sku"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Amount      float64            `bson:"amount"`
	Currency    string             `bson:"currency"`
	Images      []string           `bson:"images"`
	Url         string             `bson:"url"`
	Metadata    map[string]string  `bson:"metadata"`
	Code        string             `bson:"code"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	PlatformId  string             `bson:"platform_id"`
}

type MgoNotification struct {
	Id         primitive.ObjectID          `bson:"_id"`
	Message    string                      `bson:"message"`
	MerchantId primitive.ObjectID          `bson:"merchant_id"`
	UserId     string                      `bson:"user_id"`
	IsSystem   bool                        `bson:"is_system"`
	IsRead     bool                        `bson:"is_read"`
	CreatedAt  time.Time                   `bson:"created_at"`
	UpdatedAt  time.Time                   `bson:"updated_at"`
	Statuses   *SystemNotificationStatuses `bson:"statuses"`
}

type MgoRefundOrder struct {
	Id   primitive.ObjectID `bson:"id"`
	Uuid string             `bson:"uuid"`
}

type MgoRefund struct {
	Id             primitive.ObjectID `bson:"_id"`
	OriginalOrder  *MgoRefundOrder    `bson:"original_order"`
	ExternalId     string             `bson:"external_id"`
	Amount         float64            `bson:"amount"`
	CreatorId      primitive.ObjectID `bson:"creator_id"`
	Currency       string             `bson:"currency"`
	Status         int32              `bson:"status"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	PayerData      *RefundPayerData   `bson:"payer_data"`
	SalesTax       float32            `bson:"sales_tax"`
	IsChargeback   bool               `bson:"is_chargeback"`
	CreatedOrderId primitive.ObjectID `bson:"created_order_id,omitempty"`
	Reason         string             `bson:"reason"`
}

type MgoCustomerIdentity struct {
	MerchantId primitive.ObjectID `bson:"merchant_id"`
	ProjectId  primitive.ObjectID `bson:"project_id"`
	Type       string             `bson:"type"`
	Value      string             `bson:"value"`
	Verified   bool               `bson:"verified"`
	CreatedAt  time.Time          `bson:"created_at"`
}

type MgoCustomerIpHistory struct {
	Ip        []byte    `bson:"ip"`
	CreatedAt time.Time `bson:"created_at"`
}

type MgoCustomerAddressHistory struct {
	Country    string    `bson:"country"`
	City       string    `bson:"city"`
	PostalCode string    `bson:"postal_code"`
	State      string    `bson:"state"`
	CreatedAt  time.Time `bson:"created_at"`
}

type MgoCustomerStringValueHistory struct {
	Value     string    `bson:"value"`
	CreatedAt time.Time `bson:"created_at"`
}

type MgoCustomer struct {
	Id                    primitive.ObjectID               `bson:"_id"`
	TechEmail             string                           `bson:"tech_email"`
	ExternalId            string                           `bson:"external_id"`
	Email                 string                           `bson:"email"`
	EmailVerified         bool                             `bson:"email_verified"`
	Phone                 string                           `bson:"phone"`
	PhoneVerified         bool                             `bson:"phone_verified"`
	Name                  string                           `bson:"name"`
	Ip                    []byte                           `bson:"ip"`
	Locale                string                           `bson:"locale"`
	AcceptLanguage        string                           `bson:"accept_language"`
	UserAgent             string                           `bson:"user_agent"`
	Address               *OrderBillingAddress             `bson:"address"`
	Identity              []*MgoCustomerIdentity           `bson:"identity"`
	IpHistory             []*MgoCustomerIpHistory          `bson:"ip_history"`
	AddressHistory        []*MgoCustomerAddressHistory     `bson:"address_history"`
	LocaleHistory         []*MgoCustomerStringValueHistory `bson:"locale_history"`
	AcceptLanguageHistory []*MgoCustomerStringValueHistory `bson:"accept_language_history"`
	Metadata              map[string]string                `bson:"metadata"`
	CreatedAt             time.Time                        `bson:"created_at"`
	UpdatedAt             time.Time                        `bson:"updated_at"`
	NotifySale            bool                             `bson:"notify_sale"`
	NotifySaleEmail       string                           `bson:"notify_sale_email"`
	NotifyNewRegion       bool                             `bson:"notify_new_region"`
	NotifyNewRegionEmail  string                           `bson:"notify_new_region_email"`
}

type MgoCountry struct {
	Id                      primitive.ObjectID   `bson:"_id"`
	IsoCodeA2               string               `bson:"iso_code_a2"`
	Region                  string               `bson:"region"`
	Currency                string               `bson:"currency"`
	PaymentsAllowed         bool                 `bson:"payments_allowed"`
	ChangeAllowed           bool                 `bson:"change_allowed"`
	VatEnabled              bool                 `bson:"vat_enabled"`
	VatCurrency             string               `bson:"vat_currency"`
	PriceGroupId            string               `bson:"price_group_id"`
	VatThreshold            *CountryVatThreshold `bson:"vat_threshold"`
	VatPeriodMonth          int32                `bson:"vat_period_month"`
	VatDeadlineDays         int32                `bson:"vat_deadline_days"`
	VatStoreYears           int32                `bson:"vat_store_years"`
	VatCurrencyRatesPolicy  string               `bson:"vat_currency_rates_policy"`
	VatCurrencyRatesSource  string               `bson:"vat_currency_rates_source"`
	PayerTariffRegion       string               `bson:"payer_tariff_region"`
	CreatedAt               time.Time            `bson:"created_at"`
	UpdatedAt               time.Time            `bson:"updated_at"`
	HighRiskPaymentsAllowed bool                 `bson:"high_risk_payments_allowed"`
	HighRiskChangeAllowed   bool                 `bson:"high_risk_change_allowed"`
}

type MgoPayoutCostSystem struct {
	Id                    primitive.ObjectID `bson:"_id"`
	IntrabankCostAmount   float64            `bson:"intrabank_cost_amount"`
	IntrabankCostCurrency string             `bson:"intrabank_cost_currency"`
	InterbankCostAmount   float64            `bson:"interbank_cost_amount"`
	InterbankCostCurrency string             `bson:"interbank_cost_currency"`
	IsActive              bool               `bson:"is_active"`
	CreatedAt             time.Time          `bson:"created_at"`
}

type MgoPaymentChannelCostSystem struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name"`
	Region             string             `bson:"region"`
	Country            string             `bson:"country"`
	Percent            float64            `bson:"percent"`
	FixAmount          float64            `bson:"fix_amount"`
	FixAmountCurrency  string             `bson:"fix_amount_currency"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
	IsActive           bool               `bson:"is_active"`
	MccCode            string             `bson:"mcc_code"`
	OperatingCompanyId string             `bson:"operating_company_id"`
}

type MgoPaymentChannelCostMerchant struct {
	Id                      primitive.ObjectID `bson:"_id"`
	MerchantId              primitive.ObjectID `bson:"merchant_id"`
	Name                    string             `bson:"name"`
	PayoutCurrency          string             `bson:"payout_currency"`
	MinAmount               float64            `bson:"min_amount"`
	Region                  string             `bson:"region"`
	Country                 string             `bson:"country"`
	MethodPercent           float64            `bson:"method_percent"`
	MethodFixAmount         float64            `bson:"method_fix_amount"`
	MethodFixAmountCurrency string             `bson:"method_fix_amount_currency"`
	PsPercent               float64            `bson:"ps_percent"`
	PsFixedFee              float64            `bson:"ps_fixed_fee"`
	PsFixedFeeCurrency      string             `bson:"ps_fixed_fee_currency"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
	IsActive                bool               `bson:"is_active"`
	MccCode                 string             `bson:"mcc_code"`
}

type MgoMoneyBackCostSystem struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name"`
	PayoutCurrency     string             `bson:"payout_currency"`
	UndoReason         string             `bson:"undo_reason"`
	Region             string             `bson:"region"`
	Country            string             `bson:"country"`
	DaysFrom           int32              `bson:"days_from"`
	PaymentStage       int32              `bson:"payment_stage"`
	Percent            float64            `bson:"percent"`
	FixAmount          float64            `bson:"fix_amount"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
	IsActive           bool               `bson:"is_active"`
	MccCode            string             `bson:"mcc_code"`
	OperatingCompanyId string             `bson:"operating_company_id"`
	FixAmountCurrency  string             `bson:"fix_amount_currency"`
}

type MgoMoneyBackCostMerchant struct {
	Id                primitive.ObjectID `bson:"_id"`
	MerchantId        primitive.ObjectID `bson:"merchant_id"`
	Name              string             `bson:"name"`
	PayoutCurrency    string             `bson:"payout_currency"`
	UndoReason        string             `bson:"undo_reason"`
	Region            string             `bson:"region"`
	Country           string             `bson:"country"`
	DaysFrom          int32              `bson:"days_from"`
	PaymentStage      int32              `bson:"payment_stage"`
	Percent           float64            `bson:"percent"`
	FixAmount         float64            `bson:"fix_amount"`
	FixAmountCurrency string             `bson:"fix_amount_currency"`
	IsPaidByMerchant  bool               `bson:"is_paid_by_merchant"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
	IsActive          bool               `bson:"is_active"`
	MccCode           string             `bson:"mcc_code"`
}

type MgoPriceTable struct {
	Id       primitive.ObjectID    `bson:"_id"`
	Currency string                `bson:"currency"`
	Ranges   []*MgoPriceTableRange `bson:"range"`
}

type MgoPriceTableRange struct {
	From     float64 `bson:"from"`
	To       float64 `bson:"to"`
	Position int32   `bson:"position"`
}

type MgoAccountingEntrySource struct {
	Id   primitive.ObjectID `bson:"id"`
	Type string             `bson:"type"`
}

type MgoAccountingEntry struct {
	Id                 primitive.ObjectID        `bson:"_id"`
	Object             string                    `bson:"object"`
	Type               string                    `bson:"type"`
	Source             *MgoAccountingEntrySource `bson:"source"`
	MerchantId         primitive.ObjectID        `bson:"merchant_id"`
	Amount             float64                   `bson:"amount"`
	Currency           string                    `bson:"currency"`
	Reason             string                    `bson:"reason"`
	Status             string                    `bson:"status"`
	Country            string                    `bson:"country"`
	OriginalAmount     float64                   `bson:"original_amount"`
	OriginalCurrency   string                    `bson:"original_currency"`
	LocalAmount        float64                   `bson:"local_amount"`
	LocalCurrency      string                    `bson:"local_currency"`
	CreatedAt          time.Time                 `bson:"created_at"`
	AvailableOn        time.Time                 `bson:"available_on"`
	OperatingCompanyId string                    `bson:"operating_company_id"`
}

type MgoRoyaltyReport struct {
	Id                 primitive.ObjectID    `bson:"_id"`
	MerchantId         primitive.ObjectID    `bson:"merchant_id"`
	CreatedAt          time.Time             `bson:"created_at"`
	UpdatedAt          time.Time             `bson:"updated_at"`
	PayoutDate         time.Time             `bson:"payout_date"`
	Status             string                `bson:"status"`
	PeriodFrom         time.Time             `bson:"period_from"`
	PeriodTo           time.Time             `bson:"period_to"`
	AcceptExpireAt     time.Time             `bson:"accept_expire_at"`
	AcceptedAt         time.Time             `bson:"accepted_at"`
	Totals             *RoyaltyReportTotals  `bson:"totals"`
	Currency           string                `bson:"currency"`
	Summary            *RoyaltyReportSummary `bson:"summary"`
	DisputeReason      string                `bson:"dispute_reason"`
	DisputeStartedAt   time.Time             `bson:"dispute_started_at"`
	DisputeClosedAt    time.Time             `bson:"dispute_closed_at"`
	IsAutoAccepted     bool                  `bson:"is_auto_accepted"`
	PayoutDocumentId   string                `bson:"payout_document_id"`
	OperatingCompanyId string                `bson:"operating_company_id"`
}

type MgoRoyaltyReportCorrectionItem struct {
	AccountingEntryId primitive.ObjectID `bson:"accounting_entry_id"`
	Amount            float64            `bson:"amount"`
	Currency          string             `bson:"currency"`
	Reason            string             `bson:"reason"`
	EntryDate         time.Time          `bson:"entry_date"`
}

type MgoRoyaltyReportChanges struct {
	Id              primitive.ObjectID `bson:"_id"`
	RoyaltyReportId primitive.ObjectID `bson:"royalty_report_id"`
	Source          string             `bson:"source"`
	Ip              string             `bson:"ip"`
	Hash            string             `bson:"hash"`
	CreatedAt       time.Time          `bson:"created_at"`
}

type MgoOrderViewPrivate struct {
	Id                                         primitive.ObjectID             `bson:"_id" json:"-"`
	Uuid                                       string                         `bson:"uuid" json:"uuid"`
	TotalPaymentAmount                         float64                        `bson:"total_payment_amount" json:"total_payment_amount"`
	Currency                                   string                         `bson:"currency" json:"currency"`
	Project                                    *MgoOrderProject               `bson:"project" json:"project"`
	CreatedAt                                  time.Time                      `bson:"created_at" json:"created_at"`
	Transaction                                string                         `bson:"pm_order_id" json:"transaction"`
	PaymentMethod                              *MgoOrderPaymentMethod         `bson:"payment_method" json:"payment_method"`
	CountryCode                                string                         `bson:"country_code" json:"country_code"`
	MerchantId                                 primitive.ObjectID             `bson:"merchant_id" json:"merchant_id"`
	Locale                                     string                         `bson:"locale" json:"locale"`
	Status                                     string                         `bson:"status" json:"status"`
	TransactionDate                            time.Time                      `bson:"pm_order_close_date" json:"transaction_date"`
	User                                       *OrderUser                     `bson:"user" json:"user"`
	BillingAddress                             *OrderBillingAddress           `bson:"billing_address" json:"billing_address"`
	Type                                       string                         `bson:"type" json:"type"`
	IsVatDeduction                             bool                           `bson:"is_vat_deduction" json:"is_vat_deduction"`
	PaymentGrossRevenueLocal                   *OrderViewMoney                `bson:"payment_gross_revenue_local" json:"payment_gross_revenue_local"`
	PaymentGrossRevenueOrigin                  *OrderViewMoney                `bson:"payment_gross_revenue_origin" json:"payment_gross_revenue_origin"`
	PaymentGrossRevenue                        *OrderViewMoney                `bson:"payment_gross_revenue" json:"payment_gross_revenue"`
	PaymentTaxFee                              *OrderViewMoney                `bson:"payment_tax_fee" json:"payment_tax_fee"`
	PaymentTaxFeeLocal                         *OrderViewMoney                `bson:"payment_tax_fee_local" json:"payment_tax_fee_local"`
	PaymentTaxFeeOrigin                        *OrderViewMoney                `bson:"payment_tax_fee_origin" json:"payment_tax_fee_origin"`
	PaymentTaxFeeCurrencyExchangeFee           *OrderViewMoney                `bson:"payment_tax_fee_currency_exchange_fee" json:"payment_tax_fee_currency_exchange_fee"`
	PaymentTaxFeeTotal                         *OrderViewMoney                `bson:"payment_tax_fee_total" json:"payment_tax_fee_total"`
	PaymentGrossRevenueFx                      *OrderViewMoney                `bson:"payment_gross_revenue_fx" json:"payment_gross_revenue_fx"`
	PaymentGrossRevenueFxTaxFee                *OrderViewMoney                `bson:"payment_gross_revenue_fx_tax_fee" json:"payment_gross_revenue_fx_tax_fee"`
	PaymentGrossRevenueFxProfit                *OrderViewMoney                `bson:"payment_gross_revenue_fx_profit" json:"payment_gross_revenue_fx_profit"`
	GrossRevenue                               *OrderViewMoney                `bson:"gross_revenue" json:"gross_revenue"`
	TaxFee                                     *OrderViewMoney                `bson:"tax_fee" json:"tax_fee"`
	TaxFeeCurrencyExchangeFee                  *OrderViewMoney                `bson:"tax_fee_currency_exchange_fee" json:"tax_fee_currency_exchange_fee"`
	TaxFeeTotal                                *OrderViewMoney                `bson:"tax_fee_total" json:"tax_fee_total"`
	MethodFeeTotal                             *OrderViewMoney                `bson:"method_fee_total" json:"method_fee_total"`
	MethodFeeTariff                            *OrderViewMoney                `bson:"method_fee_tariff" json:"method_fee_tariff"`
	PaysuperMethodFeeTariffSelfCost            *OrderViewMoney                `bson:"paysuper_method_fee_tariff_self_cost" json:"paysuper_method_fee_tariff_self_cost"`
	PaysuperMethodFeeProfit                    *OrderViewMoney                `bson:"paysuper_method_fee_profit" json:"paysuper_method_fee_profit"`
	MethodFixedFeeTariff                       *OrderViewMoney                `bson:"method_fixed_fee_tariff" json:"method_fixed_fee_tariff"`
	PaysuperMethodFixedFeeTariffFxProfit       *OrderViewMoney                `bson:"paysuper_method_fixed_fee_tariff_fx_profit" json:"paysuper_method_fixed_fee_tariff_fx_profit"`
	PaysuperMethodFixedFeeTariffSelfCost       *OrderViewMoney                `bson:"paysuper_method_fixed_fee_tariff_self_cost" json:"paysuper_method_fixed_fee_tariff_self_cost"`
	PaysuperMethodFixedFeeTariffTotalProfit    *OrderViewMoney                `bson:"paysuper_method_fixed_fee_tariff_total_profit" json:"paysuper_method_fixed_fee_tariff_total_profit"`
	PaysuperFixedFee                           *OrderViewMoney                `bson:"paysuper_fixed_fee" json:"paysuper_fixed_fee"`
	PaysuperFixedFeeFxProfit                   *OrderViewMoney                `bson:"paysuper_fixed_fee_fx_profit" json:"paysuper_fixed_fee_fx_profit"`
	FeesTotal                                  *OrderViewMoney                `bson:"fees_total" json:"fees_total"`
	FeesTotalLocal                             *OrderViewMoney                `bson:"fees_total_local" json:"fees_total_local"`
	NetRevenue                                 *OrderViewMoney                `bson:"net_revenue" json:"net_revenue"`
	PaysuperMethodTotalProfit                  *OrderViewMoney                `bson:"paysuper_method_total_profit" json:"paysuper_method_total_profit"`
	PaysuperTotalProfit                        *OrderViewMoney                `bson:"paysuper_total_profit" json:"paysuper_total_profit"`
	PaymentRefundGrossRevenueLocal             *OrderViewMoney                `bson:"payment_refund_gross_revenue_local" json:"payment_refund_gross_revenue_local"`
	PaymentRefundGrossRevenueOrigin            *OrderViewMoney                `bson:"payment_refund_gross_revenue_origin" json:"payment_refund_gross_revenue_origin"`
	PaymentRefundGrossRevenue                  *OrderViewMoney                `bson:"payment_refund_gross_revenue" json:"payment_refund_gross_revenue"`
	PaymentRefundTaxFee                        *OrderViewMoney                `bson:"payment_refund_tax_fee" json:"payment_refund_tax_fee"`
	PaymentRefundTaxFeeLocal                   *OrderViewMoney                `bson:"payment_refund_tax_fee_local" json:"payment_refund_tax_fee_local"`
	PaymentRefundTaxFeeOrigin                  *OrderViewMoney                `bson:"payment_refund_tax_fee_origin" json:"payment_refund_tax_fee_origin"`
	PaymentRefundFeeTariff                     *OrderViewMoney                `bson:"payment_refund_fee_tariff" json:"payment_refund_fee_tariff"`
	MethodRefundFixedFeeTariff                 *OrderViewMoney                `bson:"method_refund_fixed_fee_tariff" json:"method_refund_fixed_fee_tariff"`
	RefundGrossRevenue                         *OrderViewMoney                `bson:"refund_gross_revenue" json:"refund_gross_revenue"`
	RefundGrossRevenueFx                       *OrderViewMoney                `bson:"refund_gross_revenue_fx" json:"refund_gross_revenue_fx"`
	MethodRefundFeeTariff                      *OrderViewMoney                `bson:"method_refund_fee_tariff" json:"method_refund_fee_tariff"`
	PaysuperMethodRefundFeeTariffProfit        *OrderViewMoney                `bson:"paysuper_method_refund_fee_tariff_profit" json:"paysuper_method_refund_fee_tariff_profit"`
	PaysuperMethodRefundFixedFeeTariffSelfCost *OrderViewMoney                `bson:"paysuper_method_refund_fixed_fee_tariff_self_cost" json:"paysuper_method_refund_fixed_fee_tariff_self_cost"`
	MerchantRefundFixedFeeTariff               *OrderViewMoney                `bson:"merchant_refund_fixed_fee_tariff" json:"merchant_refund_fixed_fee_tariff"`
	PaysuperMethodRefundFixedFeeTariffProfit   *OrderViewMoney                `bson:"paysuper_method_refund_fixed_fee_tariff_profit" json:"paysuper_method_refund_fixed_fee_tariff_profit"`
	RefundTaxFee                               *OrderViewMoney                `bson:"refund_tax_fee" json:"refund_tax_fee"`
	RefundTaxFeeCurrencyExchangeFee            *OrderViewMoney                `bson:"refund_tax_fee_currency_exchange_fee" json:"refund_tax_fee_currency_exchange_fee"`
	PaysuperRefundTaxFeeCurrencyExchangeFee    *OrderViewMoney                `bson:"paysuper_refund_tax_fee_currency_exchange_fee" json:"paysuper_refund_tax_fee_currency_exchange_fee"`
	RefundTaxFeeTotal                          *OrderViewMoney                `bson:"refund_tax_fee_total" json:"refund_tax_fee_total"`
	RefundReverseRevenue                       *OrderViewMoney                `bson:"refund_reverse_revenue" json:"refund_reverse_revenue"`
	RefundFeesTotal                            *OrderViewMoney                `bson:"refund_fees_total" json:"refund_fees_total"`
	RefundFeesTotalLocal                       *OrderViewMoney                `bson:"refund_fees_total_local" json:"refund_fees_total_local"`
	PaysuperRefundTotalProfit                  *OrderViewMoney                `bson:"paysuper_refund_total_profit" json:"paysuper_refund_total_profit"`
	Issuer                                     *OrderIssuer                   `bson:"issuer"`
	Items                                      []*MgoOrderItem                `bson:"items"`
	MerchantPayoutCurrency                     string                         `bson:"merchant_payout_currency"`
	ParentOrder                                *ParentOrder                   `bson:"parent_order"`
	Refund                                     *MgoOrderNotificationRefund    `bson:"refund"`
	Cancellation                               *OrderNotificationCancellation `bson:"cancellation"`
	MccCode                                    string                         `bson:"mcc_code"`
	OperatingCompanyId                         string                         `bson:"operating_company_id"`
	IsHighRisk                                 bool                           `bson:"is_high_risk"`
	RefundAllowed                              bool                           `bson:"refund_allowed"`
	VatPayer                                   string                         `bson:"vat_payer"`
	IsProduction                               bool                           `bson:"is_production"`
	TaxRate                                    float64                        `bson:"tax_rate"`
	MerchantInfo                               *OrderViewMerchantInfo         `bson:"merchant_info"`
	OrderChargeBeforeVat                       *OrderViewMoney                `bson:"order_charge_before_vat"`
}

type MgoOrderViewPublic struct {
	Id                                      primitive.ObjectID             `bson:"_id"`
	Uuid                                    string                         `bson:"uuid"`
	TotalPaymentAmount                      float64                        `bson:"total_payment_amount"`
	Currency                                string                         `bson:"currency"`
	Project                                 *MgoOrderProject               `bson:"project"`
	CreatedAt                               time.Time                      `bson:"created_at"`
	Transaction                             string                         `bson:"pm_order_id"`
	PaymentMethod                           *MgoOrderPaymentMethod         `bson:"payment_method"`
	CountryCode                             string                         `bson:"country_code"`
	MerchantId                              primitive.ObjectID             `bson:"merchant_id"`
	Locale                                  string                         `bson:"locale"`
	Status                                  string                         `bson:"status"`
	TransactionDate                         time.Time                      `bson:"pm_order_close_date"`
	User                                    *OrderUser                     `bson:"user"`
	BillingAddress                          *OrderBillingAddress           `bson:"billing_address"`
	Type                                    string                         `bson:"type"`
	IsVatDeduction                          bool                           `bson:"is_vat_deduction"`
	GrossRevenue                            *OrderViewMoney                `bson:"gross_revenue"`
	TaxFee                                  *OrderViewMoney                `bson:"tax_fee"`
	TaxFeeCurrencyExchangeFee               *OrderViewMoney                `bson:"tax_fee_currency_exchange_fee"`
	TaxFeeTotal                             *OrderViewMoney                `bson:"tax_fee_total"`
	MethodFeeTotal                          *OrderViewMoney                `bson:"method_fee_total"`
	MethodFeeTariff                         *OrderViewMoney                `bson:"method_fee_tariff"`
	MethodFixedFeeTariff                    *OrderViewMoney                `bson:"method_fixed_fee_tariff"`
	PaysuperFixedFee                        *OrderViewMoney                `bson:"paysuper_fixed_fee"`
	FeesTotal                               *OrderViewMoney                `bson:"fees_total"`
	FeesTotalLocal                          *OrderViewMoney                `bson:"fees_total_local"`
	NetRevenue                              *OrderViewMoney                `bson:"net_revenue"`
	RefundGrossRevenue                      *OrderViewMoney                `bson:"refund_gross_revenue"`
	MethodRefundFeeTariff                   *OrderViewMoney                `bson:"method_refund_fee_tariff"`
	MerchantRefundFixedFeeTariff            *OrderViewMoney                `bson:"merchant_refund_fixed_fee_tariff"`
	RefundTaxFee                            *OrderViewMoney                `bson:"refund_tax_fee"`
	RefundTaxFeeCurrencyExchangeFee         *OrderViewMoney                `bson:"refund_tax_fee_currency_exchange_fee"`
	PaysuperRefundTaxFeeCurrencyExchangeFee *OrderViewMoney                `bson:"paysuper_refund_tax_fee_currency_exchange_fee"`
	RefundReverseRevenue                    *OrderViewMoney                `bson:"refund_reverse_revenue"`
	RefundFeesTotal                         *OrderViewMoney                `bson:"refund_fees_total"`
	RefundFeesTotalLocal                    *OrderViewMoney                `bson:"refund_fees_total_local"`
	Issuer                                  *OrderIssuer                   `bson:"issuer"`
	Items                                   []*MgoOrderItem                `bson:"items"`
	MerchantPayoutCurrency                  string                         `bson:"merchant_payout_currency"`
	ParentOrder                             *ParentOrder                   `bson:"parent_order"`
	Refund                                  *MgoOrderNotificationRefund    `bson:"refund"`
	Cancellation                            *OrderNotificationCancellation `bson:"cancellation"`
	OperatingCompanyId                      string                         `bson:"operating_company_id"`
	RefundAllowed                           bool                           `bson:"refund_allowed"`
	OrderCharge                             *OrderViewMoney                `bson:"order_charge"`
	OrderChargeBeforeVat                    *OrderViewMoney                `bson:"order_charge_before_vat"`
	PaymentIpCountry                        string                         `bson:"payment_ip_country"`
	IsIpCountryMismatchBin                  bool                           `bson:"is_ip_country_mismatch_bin"`
	BillingCountryChangedByUser             bool                           `bson:"billing_country_changed_by_user"`
	VatPayer                                string                         `bson:"vat_payer"`
	IsProduction                            bool                           `bson:"is_production"`
}

type MgoKey struct {
	Id           primitive.ObjectID  `bson:"_id"`
	Code         string              `bson:"code"`
	KeyProductId primitive.ObjectID  `bson:"key_product_id"`
	PlatformId   string              `bson:"platform_id"`
	OrderId      *primitive.ObjectID `bson:"order_id"`
	CreatedAt    time.Time           `bson:"created_at"`
	ReservedTo   time.Time           `bson:"reserved_to"`
	RedeemedAt   time.Time           `bson:"redeemed_at"`
}

type MgoMerchantBalance struct {
	Id             primitive.ObjectID `bson:"_id"`
	MerchantId     primitive.ObjectID `bson:"merchant_id"`
	Currency       string             `bson:"currency"`
	Debit          float64            `bson:"debit"`
	Credit         float64            `bson:"credit"`
	RollingReserve float64            `bson:"rolling_reserve"`
	Total          float64            `bson:"total"`
	CreatedAt      time.Time          `bson:"created_at"`
}

type MgoOperatingCompany struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name"`
	Country            string             `bson:"country"`
	RegistrationNumber string             `bson:"registration_number"`
	RegistrationDate   string             `bson:"registration_date"`
	VatNumber          string             `bson:"vat_number"`
	Email              string             `bson:"email"`
	Address            string             `bson:"address"`
	VatAddress         string             `bson:"vat_address"`
	SignatoryName      string             `bson:"signatory_name"`
	SignatoryPosition  string             `bson:"signatory_position"`
	BankingDetails     string             `bson:"banking_details"`
	PaymentCountries   []string           `bson:"payment_countries"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
}

type MgoPaymentMinLimitSystem struct {
	Id        primitive.ObjectID `bson:"_id"`
	Currency  string             `bson:"currency"`
	Amount    float64            `bson:"amount"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type MgoUserRole struct {
	Id         primitive.ObjectID  `bson:"_id"`
	MerchantId *primitive.ObjectID `bson:"merchant_id"`
	Role       string              `bson:"role"`
	Status     string              `bson:"status"`
	UserId     *primitive.ObjectID `bson:"user_id"`
	FirstName  string              `bson:"first_name"`
	LastName   string              `bson:"last_name"`
	Email      string              `bson:"email"`
	CreatedAt  time.Time           `bson:"created_at"`
	UpdatedAt  time.Time           `bson:"updated_at"`
}

func (m *MerchantBalance) MarshalBSON() ([]byte, error) {
	st := &MgoMerchantBalance{
		Currency:       m.Currency,
		Debit:          m.Debit,
		Credit:         m.Credit,
		RollingReserve: m.RollingReserve,
		Total:          m.Total,
	}
	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, errors.New(ErrorInvalidObjectId)
	}

	st.MerchantId = merchantOid

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *MerchantBalance) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoMerchantBalance)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.MerchantId = decoded.MerchantId.Hex()
	m.Currency = decoded.Currency
	m.Debit = decoded.Debit
	m.Credit = decoded.Credit
	m.RollingReserve = decoded.RollingReserve
	m.Total = decoded.Total

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *Country) MarshalBSON() ([]byte, error) {
	st := &MgoCountry{
		IsoCodeA2:               m.IsoCodeA2,
		Region:                  m.Region,
		Currency:                m.Currency,
		PaymentsAllowed:         m.PaymentsAllowed,
		ChangeAllowed:           m.ChangeAllowed,
		VatEnabled:              m.VatEnabled,
		PriceGroupId:            m.PriceGroupId,
		VatCurrency:             m.VatCurrency,
		VatThreshold:            m.VatThreshold,
		VatPeriodMonth:          m.VatPeriodMonth,
		VatStoreYears:           m.VatStoreYears,
		VatCurrencyRatesPolicy:  m.VatCurrencyRatesPolicy,
		VatCurrencyRatesSource:  m.VatCurrencyRatesSource,
		PayerTariffRegion:       m.PayerTariffRegion,
		HighRiskPaymentsAllowed: m.HighRiskPaymentsAllowed,
		HighRiskChangeAllowed:   m.HighRiskChangeAllowed,
	}
	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *Country) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoCountry)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.IsoCodeA2 = decoded.IsoCodeA2
	m.Region = decoded.Region
	m.Currency = decoded.Currency
	m.PaymentsAllowed = decoded.PaymentsAllowed
	m.ChangeAllowed = decoded.ChangeAllowed
	m.VatEnabled = decoded.VatEnabled
	m.PriceGroupId = decoded.PriceGroupId
	m.VatCurrency = decoded.VatCurrency
	m.VatThreshold = decoded.VatThreshold
	m.VatPeriodMonth = decoded.VatPeriodMonth
	m.VatDeadlineDays = decoded.VatDeadlineDays
	m.VatStoreYears = decoded.VatStoreYears
	m.VatCurrencyRatesPolicy = decoded.VatCurrencyRatesPolicy
	m.VatCurrencyRatesSource = decoded.VatCurrencyRatesSource
	m.PayerTariffRegion = decoded.PayerTariffRegion
	m.HighRiskPaymentsAllowed = decoded.HighRiskPaymentsAllowed
	m.HighRiskChangeAllowed = decoded.HighRiskChangeAllowed

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *Commission) MarshalBSON() ([]byte, error) {
	paymentMethodOid, err := primitive.ObjectIDFromHex(m.PaymentMethodId)

	if err != nil {
		return nil, err
	}

	projectOid, err := primitive.ObjectIDFromHex(m.ProjectId)

	if err != nil {
		return nil, err
	}

	st := &MgoCommissionBilling{
		PaymentMethodId:         paymentMethodOid,
		ProjectId:               projectOid,
		PaymentMethodCommission: m.PaymentMethodCommission,
		PspCommission:           m.PspCommission,
		TotalCommissionToUser:   m.TotalCommissionToUser,
	}

	t, err := ptypes.Timestamp(m.StartDate)

	if err != nil {
		return nil, err
	}

	st.StartDate = t

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *Commission) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoCommissionBilling)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.PaymentMethodId = decoded.PaymentMethodId.Hex()
	m.ProjectId = decoded.ProjectId.Hex()
	m.PaymentMethodCommission = decoded.PaymentMethodCommission
	m.PspCommission = decoded.PspCommission
	m.TotalCommissionToUser = decoded.TotalCommissionToUser

	m.StartDate, err = ptypes.TimestampProto(decoded.StartDate)

	if err != nil {
		return err
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	return err
}

func (m *Order) MarshalBSON() ([]byte, error) {
	projectOid, err := primitive.ObjectIDFromHex(m.Project.Id)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.Project.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoOrder{
		Uuid:               m.Uuid,
		Transaction:        m.Transaction,
		Object:             "order",
		Status:             m.GetPublicStatus(),
		PrivateStatus:      m.PrivateStatus,
		Description:        m.Description,
		Canceled:           m.PrivateStatus == recurringpb.OrderStatusPaymentSystemCanceled,
		Cancellation:       m.Cancellation,
		Refunded:           m.PrivateStatus == recurringpb.OrderStatusRefund,
		ReceiptEmail:       m.GetReceiptUserEmail(),
		ReceiptPhone:       m.GetReceiptUserPhone(),
		ReceiptNumber:      m.ReceiptNumber,
		ReceiptUrl:         m.ReceiptUrl,
		AgreementVersion:   m.AgreementVersion,
		AgreementAccepted:  m.AgreementAccepted,
		NotifySale:         m.NotifySale,
		NotifySaleEmail:    m.NotifySaleEmail,
		Issuer:             m.Issuer,
		TotalPaymentAmount: m.TotalPaymentAmount,
		Currency:           m.Currency,
		User:               m.User,
		BillingAddress:     m.BillingAddress,
		Tax:                m.Tax,
		Items:              []*MgoOrderItem{},
		Metadata:           m.Metadata,
		PrivateMetadata:    m.PrivateMetadata,
		Project: &MgoOrderProject{
			Id:                      projectOid,
			MerchantId:              merchantOid,
			UrlSuccess:              m.Project.UrlSuccess,
			UrlFail:                 m.Project.UrlFail,
			NotifyEmails:            m.Project.NotifyEmails,
			SendNotifyEmail:         m.Project.SendNotifyEmail,
			SecretKey:               m.Project.SecretKey,
			UrlCheckAccount:         m.Project.UrlCheckAccount,
			UrlProcessPayment:       m.Project.UrlProcessPayment,
			CallbackProtocol:        m.Project.CallbackProtocol,
			UrlChargebackPayment:    m.Project.UrlChargebackPayment,
			UrlCancelPayment:        m.Project.UrlCancelPayment,
			UrlRefundPayment:        m.Project.UrlRefundPayment,
			UrlFraudPayment:         m.Project.UrlFraudPayment,
			Status:                  m.Project.Status,
			MerchantRoyaltyCurrency: m.Project.MerchantRoyaltyCurrency,
			RedirectSettings:        m.Project.RedirectSettings,
		},
		ProjectOrderId:              m.ProjectOrderId,
		ProjectAccount:              m.ProjectAccount,
		ProjectParams:               m.ProjectParams,
		IsJsonRequest:               m.IsJsonRequest,
		OrderAmount:                 m.OrderAmount,
		PaymentMethodPayerAccount:   m.PaymentMethodPayerAccount,
		PaymentMethodTxnParams:      m.PaymentMethodTxnParams,
		PaymentRequisites:           m.PaymentRequisites,
		UserAddressDataRequired:     m.UserAddressDataRequired,
		Products:                    m.Products,
		IsNotificationsSent:         m.IsNotificationsSent,
		CountryRestriction:          m.CountryRestriction,
		ParentOrder:                 m.ParentOrder,
		Type:                        m.Type,
		IsVatDeduction:              m.IsVatDeduction,
		CountryCode:                 m.GetCountry(),
		ProductType:                 m.ProductType,
		PlatformId:                  m.PlatformId,
		Keys:                        m.Keys,
		IsKeyProductNotified:        m.IsKeyProductNotified,
		ReceiptId:                   m.ReceiptId,
		IsBuyForVirtualCurrency:     m.IsBuyForVirtualCurrency,
		MccCode:                     m.MccCode,
		OperatingCompanyId:          m.OperatingCompanyId,
		IsHighRisk:                  m.IsHighRisk,
		ChargeCurrency:              m.ChargeCurrency,
		ChargeAmount:                m.ChargeAmount,
		PaymentIpCountry:            m.PaymentIpCountry,
		IsIpCountryMismatchBin:      m.IsIpCountryMismatchBin,
		BillingCountryChangedByUser: m.BillingCountryChangedByUser,
		IsRefundAllowed:             m.IsRefundAllowed,
		VatPayer:                    m.VatPayer,
		IsProduction:                m.IsProduction,
		FormMode:                    m.FormMode,
	}

	if m.Refund != nil {
		st.Refund = &MgoOrderNotificationRefund{
			Amount:        m.Refund.Amount,
			Currency:      m.Refund.Currency,
			Reason:        m.Refund.Reason,
			Code:          m.Refund.Code,
			ReceiptNumber: m.Refund.ReceiptNumber,
			ReceiptUrl:    m.Refund.ReceiptUrl,
		}
	}

	for _, v := range m.Items {
		item := &MgoOrderItem{
			Object:      v.Object,
			Sku:         v.Sku,
			Name:        v.Name,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
			Images:      v.Images,
			Url:         v.Url,
			Metadata:    v.Metadata,
			Code:        v.Code,
			PlatformId:  v.PlatformId,
		}

		if len(v.Id) <= 0 {
			item.Id = primitive.NewObjectID()
		} else {
			itemOid, err := primitive.ObjectIDFromHex(v.Id)

			if err != nil {
				return nil, errors.New(ErrorInvalidObjectId)
			}
			item.Id = itemOid
		}

		item.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		item.CreatedAt, _ = ptypes.Timestamp(v.UpdatedAt)
		st.Items = append(st.Items, item)
	}

	if m.PaymentMethod != nil {
		paymentMethodOid, err := primitive.ObjectIDFromHex(m.PaymentMethod.Id)

		if err != nil {
			return nil, err
		}

		paymentSystemOid, err := primitive.ObjectIDFromHex(m.PaymentMethod.PaymentSystemId)

		if err != nil {
			return nil, err
		}

		st.PaymentMethod = &MgoOrderPaymentMethod{
			Id:              paymentMethodOid,
			Name:            m.PaymentMethod.Name,
			ExternalId:      m.PaymentMethod.ExternalId,
			Params:          m.PaymentMethod.Params,
			PaymentSystemId: paymentSystemOid,
			Group:           m.PaymentMethod.Group,
			Saved:           m.PaymentMethod.Saved,
			Handler:         m.PaymentMethod.Handler,
			RefundAllowed:   m.PaymentMethod.RefundAllowed,
		}

		if m.PaymentMethod.Card != nil {
			st.PaymentMethod.Card = m.PaymentMethod.Card
		}
		if m.PaymentMethod.Wallet != nil {
			st.PaymentMethod.Wallet = m.PaymentMethod.Wallet
		}
		if m.PaymentMethod.CryptoCurrency != nil {
			st.PaymentMethod.CryptoCurrency = m.PaymentMethod.CryptoCurrency
		}
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	if m.ProjectLastRequestedAt != nil {
		t, err := ptypes.Timestamp(m.ProjectLastRequestedAt)

		if err != nil {
			return nil, err
		}

		st.ProjectLastRequestedAt = t
	}

	if m.PaymentMethodOrderClosedAt != nil {
		t, err := ptypes.Timestamp(m.PaymentMethodOrderClosedAt)

		if err != nil {
			return nil, err
		}

		st.PaymentMethodOrderClosedAt = t
	}

	if m.ExpireDateToFormInput != nil {
		t, err := ptypes.Timestamp(m.ExpireDateToFormInput)

		if err != nil {
			return nil, err
		}

		st.ExpireDateToFormInput = t
	} else {
		st.ExpireDateToFormInput = time.Now()
	}

	if m.Project != nil && len(m.Project.Name) > 0 {
		for k, v := range m.Project.Name {
			st.Project.Name = append(st.Project.Name, &MgoMultiLang{Lang: k, Value: v})
		}
	}

	if m.CanceledAt != nil {
		t, err := ptypes.Timestamp(m.CanceledAt)

		if err != nil {
			return nil, err
		}

		st.CanceledAt = t
	}

	if m.RefundedAt != nil {
		t, err := ptypes.Timestamp(m.RefundedAt)

		if err != nil {
			return nil, err
		}

		st.RefundedAt = t
	}

	if m.ParentPaymentAt != nil {
		t, err := ptypes.Timestamp(m.ParentPaymentAt)

		if err != nil {
			return nil, err
		}

		st.ParentPaymentAt = t
	}

	return bson.Marshal(st)
}

func (m *Order) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOrder)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Uuid = decoded.Uuid
	m.Transaction = decoded.Transaction
	m.Object = decoded.Object
	m.Status = decoded.Status
	m.PrivateStatus = decoded.PrivateStatus
	m.Description = decoded.Description
	m.Canceled = decoded.Canceled
	m.Cancellation = decoded.Cancellation
	m.Refunded = decoded.Refunded
	m.ReceiptEmail = decoded.ReceiptEmail
	m.ReceiptPhone = decoded.ReceiptPhone
	m.ReceiptNumber = decoded.ReceiptNumber
	m.ReceiptUrl = decoded.ReceiptUrl
	m.AgreementVersion = decoded.AgreementVersion
	m.AgreementAccepted = decoded.AgreementAccepted
	m.NotifySale = decoded.NotifySale
	m.NotifySaleEmail = decoded.NotifySaleEmail
	m.Issuer = decoded.Issuer
	m.TotalPaymentAmount = decoded.TotalPaymentAmount
	m.Currency = decoded.Currency
	m.User = decoded.User
	m.BillingAddress = decoded.BillingAddress
	m.Tax = decoded.Tax
	m.PaymentMethod = getPaymentMethodOrder(decoded.PaymentMethod)
	m.Items = []*OrderItem{}
	m.PlatformId = decoded.PlatformId
	m.ProductType = decoded.ProductType
	m.Keys = decoded.Keys
	m.IsKeyProductNotified = decoded.IsKeyProductNotified
	m.ReceiptId = decoded.ReceiptId
	m.IsBuyForVirtualCurrency = decoded.IsBuyForVirtualCurrency
	m.MccCode = decoded.MccCode
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.IsHighRisk = decoded.IsHighRisk
	m.ChargeCurrency = decoded.ChargeCurrency
	m.ChargeAmount = decoded.ChargeAmount
	m.PaymentIpCountry = decoded.PaymentIpCountry
	m.IsIpCountryMismatchBin = decoded.IsIpCountryMismatchBin
	m.BillingCountryChangedByUser = decoded.BillingCountryChangedByUser
	m.IsRefundAllowed = decoded.IsRefundAllowed
	m.FormMode = decoded.FormMode

	if decoded.Refund != nil {
		m.Refund = &OrderNotificationRefund{
			Amount:        decoded.Refund.Amount,
			Currency:      decoded.Refund.Currency,
			Reason:        decoded.Refund.Reason,
			Code:          decoded.Refund.Code,
			ReceiptNumber: decoded.Refund.ReceiptNumber,
			ReceiptUrl:    decoded.Refund.ReceiptUrl,
		}
	}
	m.Metadata = decoded.Metadata
	m.PrivateMetadata = decoded.PrivateMetadata
	m.Project = getOrderProject(decoded.Project)

	for _, v := range decoded.Items {
		item := &OrderItem{
			Id:          v.Id.Hex(),
			Object:      v.Object,
			Sku:         v.Sku,
			Name:        v.Name,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
			Images:      v.Images,
			Url:         v.Url,
			Metadata:    v.Metadata,
			Code:        v.Code,
			PlatformId:  v.PlatformId,
		}
		item.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		item.UpdatedAt, _ = ptypes.TimestampProto(v.UpdatedAt)
		m.Items = append(m.Items, item)
	}

	if decoded.Project != nil {
		nameLen := len(decoded.Project.Name)
		if nameLen > 0 {
			m.Project.Name = make(map[string]string, nameLen)

			for _, v := range decoded.Project.Name {
				m.Project.Name[v.Lang] = v.Value
			}
		}
	}

	m.ProjectOrderId = decoded.ProjectOrderId
	m.ProjectAccount = decoded.ProjectAccount
	m.ProjectParams = decoded.ProjectParams
	m.IsJsonRequest = decoded.IsJsonRequest
	m.OrderAmount = decoded.OrderAmount
	m.PaymentMethodPayerAccount = decoded.PaymentMethodPayerAccount
	m.PaymentMethodTxnParams = decoded.PaymentMethodTxnParams
	m.PaymentRequisites = decoded.PaymentRequisites
	m.UserAddressDataRequired = decoded.UserAddressDataRequired
	m.Products = decoded.Products
	m.IsNotificationsSent = decoded.IsNotificationsSent
	m.CountryRestriction = decoded.CountryRestriction
	m.ParentOrder = decoded.ParentOrder
	m.Type = decoded.Type
	m.IsVatDeduction = decoded.IsVatDeduction
	m.CountryCode = decoded.CountryCode
	m.VatPayer = decoded.VatPayer
	m.IsProduction = decoded.IsProduction

	m.PaymentMethodOrderClosedAt, err = ptypes.TimestampProto(decoded.PaymentMethodOrderClosedAt)
	if err != nil {
		return err
	}

	m.ParentPaymentAt, err = ptypes.TimestampProto(decoded.ParentPaymentAt)

	if err != nil {
		return err
	}

	m.ProjectLastRequestedAt, err = ptypes.TimestampProto(decoded.ProjectLastRequestedAt)
	if err != nil {
		return err
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}
	m.CanceledAt, err = ptypes.TimestampProto(decoded.CanceledAt)
	if err != nil {
		return err
	}

	m.RefundedAt, err = ptypes.TimestampProto(decoded.RefundedAt)
	if err != nil {
		return err
	}

	m.ExpireDateToFormInput, err = ptypes.TimestampProto(decoded.ExpireDateToFormInput)
	if err != nil {
		return err
	}

	return nil
}

func (m *Notification) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoNotification{
		Message:    m.Message,
		IsSystem:   m.IsSystem,
		IsRead:     m.IsRead,
		MerchantId: oid,
		UserId:     m.UserId,
		Statuses:   m.Statuses,
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err = primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
		m.CreatedAt, _ = ptypes.TimestampProto(st.CreatedAt)
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
		m.UpdatedAt, _ = ptypes.TimestampProto(st.UpdatedAt)
	}

	return bson.Marshal(st)
}

func (m *Notification) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoNotification)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Message = decoded.Message
	m.IsSystem = decoded.IsSystem
	m.IsRead = decoded.IsRead
	m.MerchantId = decoded.MerchantId.Hex()
	m.UserId = decoded.UserId
	m.Statuses = decoded.Statuses

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *Refund) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.OriginalOrder.Id)

	if err != nil {
		return nil, err
	}

	creatorOid, err := primitive.ObjectIDFromHex(m.CreatorId)

	if err != nil {
		return nil, err
	}

	st := &MgoRefund{
		OriginalOrder: &MgoRefundOrder{
			Id:   oid,
			Uuid: m.OriginalOrder.Uuid,
		},
		ExternalId:   m.ExternalId,
		Amount:       m.Amount,
		CreatorId:    creatorOid,
		Currency:     m.Currency,
		Status:       m.Status,
		PayerData:    m.PayerData,
		SalesTax:     m.SalesTax,
		IsChargeback: m.IsChargeback,
		Reason:       m.Reason,
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	if m.CreatedOrderId != "" {
		oid, err := primitive.ObjectIDFromHex(m.CreatedOrderId)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.CreatedOrderId = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *Refund) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRefund)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.OriginalOrder = &RefundOrder{
		Id:   decoded.OriginalOrder.Id.Hex(),
		Uuid: decoded.OriginalOrder.Uuid,
	}
	m.ExternalId = decoded.ExternalId
	m.Amount = decoded.Amount
	m.CreatorId = decoded.CreatorId.Hex()
	m.Currency = decoded.Currency
	m.Status = decoded.Status
	m.PayerData = decoded.PayerData
	m.SalesTax = decoded.SalesTax
	m.IsChargeback = decoded.IsChargeback
	m.Reason = decoded.Reason

	if !decoded.CreatedOrderId.IsZero() {
		m.CreatedOrderId = decoded.CreatedOrderId.Hex()
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *PaymentFormPaymentMethod) IsBankCard() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethod) IsBankCard() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsBankCard() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsCryptoCurrency() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBitcoin
}

func (m *Customer) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	st := &MgoCustomer{
		Id:                    oid,
		TechEmail:             m.TechEmail,
		ExternalId:            m.ExternalId,
		Email:                 m.Email,
		EmailVerified:         m.EmailVerified,
		Phone:                 m.Phone,
		PhoneVerified:         m.PhoneVerified,
		Name:                  m.Name,
		Ip:                    m.Ip,
		Locale:                m.Locale,
		AcceptLanguage:        m.AcceptLanguage,
		UserAgent:             m.UserAgent,
		Address:               m.Address,
		Metadata:              m.Metadata,
		Identity:              []*MgoCustomerIdentity{},
		IpHistory:             []*MgoCustomerIpHistory{},
		AddressHistory:        []*MgoCustomerAddressHistory{},
		LocaleHistory:         []*MgoCustomerStringValueHistory{},
		AcceptLanguageHistory: []*MgoCustomerStringValueHistory{},
		NotifySale:            m.NotifySale,
		NotifySaleEmail:       m.NotifySaleEmail,
		NotifyNewRegion:       m.NotifyNewRegion,
		NotifyNewRegionEmail:  m.NotifyNewRegionEmail,
	}

	for _, v := range m.Identity {
		merchantOid, err := primitive.ObjectIDFromHex(v.MerchantId)

		if err != nil {
			return nil, err
		}

		projectOid, err := primitive.ObjectIDFromHex(v.ProjectId)

		if err != nil {
			return nil, err
		}

		mgoIdentity := &MgoCustomerIdentity{
			MerchantId: merchantOid,
			ProjectId:  projectOid,
			Type:       v.Type,
			Value:      v.Value,
			Verified:   v.Verified,
		}

		mgoIdentity.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		st.Identity = append(st.Identity, mgoIdentity)
	}

	for _, v := range m.IpHistory {
		mgoIdentity := &MgoCustomerIpHistory{Ip: v.Ip}
		mgoIdentity.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		st.IpHistory = append(st.IpHistory, mgoIdentity)
	}

	for _, v := range m.AddressHistory {
		mgoIdentity := &MgoCustomerAddressHistory{
			Country:    v.Country,
			City:       v.City,
			PostalCode: v.PostalCode,
			State:      v.State,
		}
		mgoIdentity.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		st.AddressHistory = append(st.AddressHistory, mgoIdentity)
	}

	for _, v := range m.LocaleHistory {
		mgoIdentity := &MgoCustomerStringValueHistory{Value: v.Value}
		mgoIdentity.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		st.LocaleHistory = append(st.LocaleHistory, mgoIdentity)
	}

	for _, v := range m.AcceptLanguageHistory {
		mgoIdentity := &MgoCustomerStringValueHistory{Value: v.Value}
		mgoIdentity.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		st.AcceptLanguageHistory = append(st.AcceptLanguageHistory, mgoIdentity)
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *Customer) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoCustomer)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.TechEmail = decoded.TechEmail
	m.ExternalId = decoded.ExternalId
	m.Email = decoded.Email
	m.EmailVerified = decoded.EmailVerified
	m.Phone = decoded.Phone
	m.PhoneVerified = decoded.PhoneVerified
	m.Name = decoded.Name
	m.Ip = decoded.Ip
	m.Locale = decoded.Locale
	m.AcceptLanguage = decoded.AcceptLanguage
	m.UserAgent = decoded.UserAgent
	m.Address = decoded.Address
	m.Identity = []*CustomerIdentity{}
	m.IpHistory = []*CustomerIpHistory{}
	m.AddressHistory = []*CustomerAddressHistory{}
	m.LocaleHistory = []*CustomerStringValueHistory{}
	m.AcceptLanguageHistory = []*CustomerStringValueHistory{}
	m.Metadata = decoded.Metadata
	m.NotifySale = decoded.NotifySale
	m.NotifySaleEmail = decoded.NotifySaleEmail
	m.NotifyNewRegion = decoded.NotifyNewRegion
	m.NotifyNewRegionEmail = decoded.NotifyNewRegionEmail

	for _, v := range decoded.Identity {
		identity := &CustomerIdentity{
			MerchantId: v.MerchantId.Hex(),
			ProjectId:  v.ProjectId.Hex(),
			Type:       v.Type,
			Value:      v.Value,
			Verified:   v.Verified,
		}

		identity.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		m.Identity = append(m.Identity, identity)
	}

	for _, v := range decoded.IpHistory {
		identity := &CustomerIpHistory{Ip: v.Ip}
		identity.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		m.IpHistory = append(m.IpHistory, identity)
	}

	for _, v := range decoded.AddressHistory {
		identity := &CustomerAddressHistory{
			Country:    v.Country,
			City:       v.City,
			PostalCode: v.PostalCode,
			State:      v.State,
		}
		identity.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		m.AddressHistory = append(m.AddressHistory, identity)
	}

	for _, v := range decoded.LocaleHistory {
		identity := &CustomerStringValueHistory{Value: v.Value}
		identity.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		m.LocaleHistory = append(m.LocaleHistory, identity)
	}

	for _, v := range decoded.AcceptLanguageHistory {
		identity := &CustomerStringValueHistory{Value: v.Value}
		identity.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		m.AcceptLanguageHistory = append(m.AcceptLanguageHistory, identity)
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *PaymentChannelCostSystem) MarshalBSON() ([]byte, error) {
	st := &MgoPaymentChannelCostSystem{
		Name:               m.Name,
		Region:             m.Region,
		Country:            m.Country,
		Percent:            m.Percent,
		FixAmount:          m.FixAmount,
		FixAmountCurrency:  m.FixAmountCurrency,
		IsActive:           m.IsActive,
		MccCode:            m.MccCode,
		OperatingCompanyId: m.OperatingCompanyId,
	}
	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}
		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)
		if err != nil {
			return nil, err
		}
		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)
		if err != nil {
			return nil, err
		}
		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}
	return bson.Marshal(st)
}

func (m *PaymentChannelCostSystem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoPaymentChannelCostSystem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Name = decoded.Name
	m.Region = decoded.Region
	m.Country = decoded.Country
	m.Percent = decoded.Percent
	m.FixAmount = decoded.FixAmount
	m.FixAmountCurrency = decoded.FixAmountCurrency
	m.IsActive = decoded.IsActive
	m.MccCode = decoded.MccCode
	m.OperatingCompanyId = decoded.OperatingCompanyId

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *PaymentChannelCostMerchant) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoPaymentChannelCostMerchant{
		MerchantId:              oid,
		Name:                    m.Name,
		PayoutCurrency:          m.PayoutCurrency,
		MinAmount:               m.MinAmount,
		Region:                  m.Region,
		Country:                 m.Country,
		MethodPercent:           m.MethodPercent,
		MethodFixAmount:         m.MethodFixAmount,
		MethodFixAmountCurrency: m.MethodFixAmountCurrency,
		PsPercent:               m.PsPercent,
		PsFixedFee:              m.PsFixedFee,
		PsFixedFeeCurrency:      m.PsFixedFeeCurrency,
		IsActive:                m.IsActive,
		MccCode:                 m.MccCode,
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err = primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}
		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)
		if err != nil {
			return nil, err
		}
		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)
		if err != nil {
			return nil, err
		}
		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}
	return bson.Marshal(st)
}

func (m *PaymentChannelCostMerchant) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoPaymentChannelCostMerchant)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.MerchantId = decoded.MerchantId.Hex()
	m.Name = decoded.Name
	m.PayoutCurrency = decoded.PayoutCurrency
	m.MinAmount = decoded.MinAmount
	m.Region = decoded.Region
	m.Country = decoded.Country
	m.MethodPercent = decoded.MethodPercent
	m.MethodFixAmount = decoded.MethodFixAmount
	m.MethodFixAmountCurrency = decoded.MethodFixAmountCurrency
	m.PsPercent = decoded.PsPercent
	m.PsFixedFee = decoded.PsFixedFee
	m.PsFixedFeeCurrency = decoded.PsFixedFeeCurrency
	m.IsActive = decoded.IsActive
	m.MccCode = decoded.MccCode

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *MoneyBackCostSystem) MarshalBSON() ([]byte, error) {
	st := &MgoMoneyBackCostSystem{
		Name:               m.Name,
		PayoutCurrency:     m.PayoutCurrency,
		UndoReason:         m.UndoReason,
		Region:             m.Region,
		Country:            m.Country,
		DaysFrom:           m.DaysFrom,
		PaymentStage:       m.PaymentStage,
		Percent:            m.Percent,
		FixAmount:          m.FixAmount,
		IsActive:           m.IsActive,
		MccCode:            m.MccCode,
		OperatingCompanyId: m.OperatingCompanyId,
		FixAmountCurrency:  m.FixAmountCurrency,
	}
	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}
		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)
		if err != nil {
			return nil, err
		}
		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)
		if err != nil {
			return nil, err
		}
		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}
	return bson.Marshal(st)
}

func (m *MoneyBackCostSystem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoMoneyBackCostSystem)
	err := bson.Unmarshal(raw, decoded)
	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Name = decoded.Name
	m.PayoutCurrency = decoded.PayoutCurrency
	m.UndoReason = decoded.UndoReason
	m.Region = decoded.Region
	m.Country = decoded.Country
	m.DaysFrom = decoded.DaysFrom
	m.PaymentStage = decoded.PaymentStage
	m.Percent = decoded.Percent
	m.FixAmount = decoded.FixAmount
	m.IsActive = decoded.IsActive
	m.MccCode = decoded.MccCode
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.FixAmountCurrency = decoded.FixAmountCurrency

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *MoneyBackCostMerchant) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoMoneyBackCostMerchant{
		MerchantId:        oid,
		Name:              m.Name,
		PayoutCurrency:    m.PayoutCurrency,
		UndoReason:        m.UndoReason,
		Region:            m.Region,
		Country:           m.Country,
		DaysFrom:          m.DaysFrom,
		PaymentStage:      m.PaymentStage,
		Percent:           m.Percent,
		FixAmount:         m.FixAmount,
		FixAmountCurrency: m.FixAmountCurrency,
		IsPaidByMerchant:  m.IsPaidByMerchant,
		IsActive:          m.IsActive,
		MccCode:           m.MccCode,
	}
	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}
		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)
		if err != nil {
			return nil, err
		}
		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)
		if err != nil {
			return nil, err
		}
		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}
	return bson.Marshal(st)
}

func (m *MoneyBackCostMerchant) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoMoneyBackCostMerchant)
	err := bson.Unmarshal(raw, decoded)
	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.MerchantId = decoded.MerchantId.Hex()
	m.Name = decoded.Name
	m.PayoutCurrency = decoded.PayoutCurrency
	m.UndoReason = decoded.UndoReason
	m.Region = decoded.Region
	m.Country = decoded.Country
	m.DaysFrom = decoded.DaysFrom
	m.PaymentStage = decoded.PaymentStage
	m.Percent = decoded.Percent
	m.FixAmount = decoded.FixAmount
	m.FixAmountCurrency = decoded.FixAmountCurrency
	m.IsPaidByMerchant = decoded.IsPaidByMerchant
	m.IsActive = decoded.IsActive
	m.MccCode = decoded.MccCode

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *PayoutCostSystem) MarshalBSON() ([]byte, error) {
	st := &MgoPayoutCostSystem{
		IntrabankCostAmount:   m.IntrabankCostAmount,
		IntrabankCostCurrency: m.IntrabankCostCurrency,
		InterbankCostAmount:   m.InterbankCostAmount,
		InterbankCostCurrency: m.InterbankCostCurrency,
		IsActive:              m.IsActive,
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}
		st.Id = oid
	}

	t, err := ptypes.Timestamp(m.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	return bson.Marshal(st)
}

func (m *PayoutCostSystem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoPayoutCostSystem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}
	m.Id = decoded.Id.Hex()
	m.IntrabankCostAmount = decoded.IntrabankCostAmount
	m.IntrabankCostCurrency = decoded.IntrabankCostCurrency
	m.InterbankCostAmount = decoded.InterbankCostAmount
	m.InterbankCostCurrency = decoded.InterbankCostCurrency
	m.IsActive = decoded.IsActive

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *PriceTable) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	st := &MgoPriceTable{
		Id:       oid,
		Currency: m.Currency,
	}

	if len(m.Ranges) > 0 {
		for _, v := range m.Ranges {
			st.Ranges = append(st.Ranges, &MgoPriceTableRange{From: v.From, To: v.To, Position: v.Position})
		}
	}

	return bson.Marshal(st)
}

func (m *PriceTable) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoPriceTable)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Currency = decoded.Currency

	rangesLen := len(decoded.Ranges)

	if rangesLen > 0 {
		m.Ranges = make([]*PriceTableRange, rangesLen)

		for i, v := range decoded.Ranges {
			m.Ranges[i] = &PriceTableRange{
				From:     v.From,
				To:       v.To,
				Position: v.Position,
			}
		}
	}

	return nil
}

func (m *AccountingEntry) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	sourceOid, err := primitive.ObjectIDFromHex(m.Source.Id)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoAccountingEntry{
		Id:     oid,
		Object: m.Object,
		Type:   m.Type,
		Source: &MgoAccountingEntrySource{
			Id:   sourceOid,
			Type: m.Source.Type,
		},
		MerchantId:         merchantOid,
		Amount:             m.Amount,
		Currency:           m.Currency,
		OriginalAmount:     m.OriginalAmount,
		OriginalCurrency:   m.OriginalCurrency,
		LocalAmount:        m.LocalAmount,
		LocalCurrency:      m.LocalCurrency,
		Country:            m.Country,
		Reason:             m.Reason,
		Status:             m.Status,
		OperatingCompanyId: m.OperatingCompanyId,
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.AvailableOn != nil {
		t, err := ptypes.Timestamp(m.AvailableOn)

		if err != nil {
			return nil, err
		}

		st.AvailableOn = t
	}

	return bson.Marshal(st)
}

func (m *AccountingEntry) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoAccountingEntry)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Object = decoded.Object
	m.Type = decoded.Type
	m.Source = &AccountingEntrySource{
		Id:   decoded.Source.Id.Hex(),
		Type: decoded.Source.Type,
	}
	m.MerchantId = decoded.MerchantId.Hex()
	m.Amount = decoded.Amount
	m.Currency = decoded.Currency
	m.OriginalAmount = decoded.OriginalAmount
	m.OriginalCurrency = decoded.OriginalCurrency
	m.LocalAmount = decoded.LocalAmount
	m.LocalCurrency = decoded.LocalCurrency
	m.Country = decoded.Country
	m.Reason = decoded.Reason
	m.Status = decoded.Status
	m.OperatingCompanyId = decoded.OperatingCompanyId

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.AvailableOn, err = ptypes.TimestampProto(decoded.AvailableOn)

	if err != nil {
		return err
	}

	return nil
}

func (m *RoyaltyReport) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoRoyaltyReport{
		Id:                 oid,
		MerchantId:         merchantOid,
		Status:             m.Status,
		Totals:             m.Totals,
		Currency:           m.Currency,
		Summary:            m.Summary,
		DisputeReason:      m.DisputeReason,
		IsAutoAccepted:     m.IsAutoAccepted,
		PayoutDocumentId:   m.PayoutDocumentId,
		OperatingCompanyId: m.OperatingCompanyId,
	}

	if m.PayoutDate != nil {
		t, err := ptypes.Timestamp(m.PayoutDate)

		if err != nil {
			return nil, err
		}

		st.PayoutDate = t
	}

	t, err := ptypes.Timestamp(m.PeriodFrom)

	if err != nil {
		return nil, err
	}

	st.PeriodFrom = t
	t, err = ptypes.Timestamp(m.PeriodTo)

	if err != nil {
		return nil, err
	}

	st.PeriodTo = t
	t, err = ptypes.Timestamp(m.AcceptExpireAt)

	if err != nil {
		return nil, err
	}

	st.AcceptExpireAt = t

	if m.AcceptedAt != nil {
		t, err = ptypes.Timestamp(m.AcceptedAt)

		if err != nil {
			return nil, err
		}

		st.AcceptedAt = t
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	if m.DisputeStartedAt != nil {
		t, err := ptypes.Timestamp(m.DisputeStartedAt)

		if err != nil {
			return nil, err
		}

		st.DisputeStartedAt = t
	}

	if m.DisputeClosedAt != nil {
		t, err := ptypes.Timestamp(m.DisputeClosedAt)

		if err != nil {
			return nil, err
		}

		st.DisputeClosedAt = t
	}

	return bson.Marshal(st)
}

func (m *RoyaltyReport) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRoyaltyReport)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.MerchantId = decoded.MerchantId.Hex()
	m.Status = decoded.Status
	m.Totals = decoded.Totals
	m.Currency = decoded.Currency
	m.Summary = decoded.Summary
	m.DisputeReason = decoded.DisputeReason
	m.IsAutoAccepted = decoded.IsAutoAccepted
	m.PayoutDocumentId = decoded.PayoutDocumentId
	m.OperatingCompanyId = decoded.OperatingCompanyId

	m.PayoutDate, err = ptypes.TimestampProto(decoded.PayoutDate)
	if err != nil {
		return err
	}

	m.PeriodFrom, err = ptypes.TimestampProto(decoded.PeriodFrom)
	if err != nil {
		return err
	}

	m.PeriodTo, err = ptypes.TimestampProto(decoded.PeriodTo)
	if err != nil {
		return err
	}

	m.AcceptExpireAt, err = ptypes.TimestampProto(decoded.AcceptExpireAt)
	if err != nil {
		return err
	}

	m.AcceptedAt, err = ptypes.TimestampProto(decoded.AcceptedAt)
	if err != nil {
		return err
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	m.DisputeStartedAt, err = ptypes.TimestampProto(decoded.DisputeStartedAt)
	if err != nil {
		return err
	}

	m.DisputeClosedAt, err = ptypes.TimestampProto(decoded.DisputeClosedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *RoyaltyReportChanges) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	royaltyReportOid, err := primitive.ObjectIDFromHex(m.RoyaltyReportId)

	if err != nil {
		return nil, err
	}

	st := &MgoRoyaltyReportChanges{
		Id:              oid,
		RoyaltyReportId: royaltyReportOid,
		Source:          m.Source,
		Ip:              m.Ip,
		Hash:            m.Hash,
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *RoyaltyReportChanges) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRoyaltyReportChanges)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.RoyaltyReportId = decoded.RoyaltyReportId.Hex()
	m.Source = decoded.Source
	m.Ip = decoded.Ip
	m.Hash = decoded.Hash

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *RoyaltyReportCorrectionItem) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.AccountingEntryId)

	if err != nil {
		return nil, err
	}

	st := &MgoRoyaltyReportCorrectionItem{
		AccountingEntryId: oid,
		Amount:            m.Amount,
		Currency:          m.Currency,
		Reason:            m.Reason,
	}

	t, err := ptypes.Timestamp(m.EntryDate)

	if err != nil {
		return nil, err
	}

	st.EntryDate = t

	return bson.Marshal(st)
}

func (m *RoyaltyReportCorrectionItem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRoyaltyReportCorrectionItem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.AccountingEntryId = decoded.AccountingEntryId.Hex()
	m.Amount = decoded.Amount
	m.Currency = decoded.Currency
	m.Reason = decoded.Reason

	m.EntryDate, err = ptypes.TimestampProto(decoded.EntryDate)
	if err != nil {
		return err
	}

	return nil
}

func (m *OrderViewPrivate) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOrderViewPrivate)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Uuid = decoded.Uuid
	m.TotalPaymentAmount = decoded.TotalPaymentAmount
	m.Currency = decoded.Currency
	m.Project = getOrderProject(decoded.Project)
	m.Transaction = decoded.Transaction
	m.PaymentMethod = getPaymentMethodOrder(decoded.PaymentMethod)
	m.CountryCode = decoded.CountryCode
	m.MerchantId = decoded.MerchantId.Hex()
	m.Locale = decoded.Locale
	m.Status = decoded.Status
	m.User = decoded.User
	m.BillingAddress = decoded.BillingAddress
	m.Type = decoded.Type
	m.Issuer = decoded.Issuer
	m.MerchantPayoutCurrency = decoded.MerchantPayoutCurrency
	m.IsVatDeduction = decoded.IsVatDeduction
	m.ParentOrder = decoded.ParentOrder
	m.Cancellation = decoded.Cancellation

	if decoded.Refund != nil {
		m.Refund = &OrderNotificationRefund{
			Amount:        decoded.Refund.Amount,
			Currency:      decoded.Refund.Currency,
			Reason:        decoded.Refund.Reason,
			Code:          decoded.Refund.Code,
			ReceiptNumber: decoded.Refund.ReceiptNumber,
			ReceiptUrl:    decoded.Refund.ReceiptUrl,
		}
	}

	m.PaymentGrossRevenueLocal = getOrderViewMoney(decoded.PaymentGrossRevenueLocal)
	m.PaymentGrossRevenueOrigin = getOrderViewMoney(decoded.PaymentGrossRevenueOrigin)
	m.PaymentGrossRevenue = getOrderViewMoney(decoded.PaymentGrossRevenue)
	m.PaymentTaxFee = getOrderViewMoney(decoded.PaymentTaxFee)
	m.PaymentTaxFeeLocal = getOrderViewMoney(decoded.PaymentTaxFeeLocal)
	m.PaymentTaxFeeOrigin = getOrderViewMoney(decoded.PaymentTaxFeeOrigin)
	m.PaymentTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.PaymentTaxFeeCurrencyExchangeFee)
	m.PaymentTaxFeeTotal = getOrderViewMoney(decoded.PaymentTaxFeeTotal)
	m.PaymentGrossRevenueFx = getOrderViewMoney(decoded.PaymentGrossRevenueFx)
	m.PaymentGrossRevenueFxTaxFee = getOrderViewMoney(decoded.PaymentGrossRevenueFxTaxFee)
	m.PaymentGrossRevenueFxProfit = getOrderViewMoney(decoded.PaymentGrossRevenueFxProfit)
	m.GrossRevenue = getOrderViewMoney(decoded.GrossRevenue)
	m.TaxFee = getOrderViewMoney(decoded.TaxFee)
	m.TaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.TaxFeeCurrencyExchangeFee)
	m.TaxFeeTotal = getOrderViewMoney(decoded.TaxFeeTotal)
	m.MethodFeeTotal = getOrderViewMoney(decoded.MethodFeeTotal)
	m.MethodFeeTariff = getOrderViewMoney(decoded.MethodFeeTariff)
	m.PaysuperMethodFeeTariffSelfCost = getOrderViewMoney(decoded.PaysuperMethodFeeTariffSelfCost)
	m.PaysuperMethodFeeProfit = getOrderViewMoney(decoded.PaysuperMethodFeeProfit)
	m.MethodFixedFeeTariff = getOrderViewMoney(decoded.MethodFixedFeeTariff)
	m.PaysuperMethodFixedFeeTariffFxProfit = getOrderViewMoney(decoded.PaysuperMethodFixedFeeTariffFxProfit)
	m.PaysuperMethodFixedFeeTariffSelfCost = getOrderViewMoney(decoded.PaysuperMethodFixedFeeTariffSelfCost)
	m.PaysuperMethodFixedFeeTariffTotalProfit = getOrderViewMoney(decoded.PaysuperMethodFixedFeeTariffTotalProfit)
	m.PaysuperFixedFee = getOrderViewMoney(decoded.PaysuperFixedFee)
	m.PaysuperFixedFeeFxProfit = getOrderViewMoney(decoded.PaysuperFixedFeeFxProfit)
	m.FeesTotal = getOrderViewMoney(decoded.FeesTotal)
	m.FeesTotalLocal = getOrderViewMoney(decoded.FeesTotalLocal)
	m.NetRevenue = getOrderViewMoney(decoded.NetRevenue)
	m.PaysuperMethodTotalProfit = getOrderViewMoney(decoded.PaysuperMethodTotalProfit)
	m.PaysuperTotalProfit = getOrderViewMoney(decoded.PaysuperTotalProfit)
	m.PaymentRefundGrossRevenueLocal = getOrderViewMoney(decoded.PaymentRefundGrossRevenueLocal)
	m.PaymentRefundGrossRevenueOrigin = getOrderViewMoney(decoded.PaymentRefundGrossRevenueOrigin)
	m.PaymentRefundGrossRevenue = getOrderViewMoney(decoded.PaymentRefundGrossRevenue)
	m.PaymentRefundTaxFee = getOrderViewMoney(decoded.PaymentRefundTaxFee)
	m.PaymentRefundTaxFeeLocal = getOrderViewMoney(decoded.PaymentRefundTaxFeeLocal)
	m.PaymentRefundTaxFeeOrigin = getOrderViewMoney(decoded.PaymentRefundTaxFeeOrigin)
	m.PaymentRefundFeeTariff = getOrderViewMoney(decoded.PaymentRefundFeeTariff)
	m.MethodRefundFixedFeeTariff = getOrderViewMoney(decoded.MethodRefundFixedFeeTariff)
	m.RefundGrossRevenue = getOrderViewMoney(decoded.RefundGrossRevenue)
	m.RefundGrossRevenueFx = getOrderViewMoney(decoded.RefundGrossRevenueFx)
	m.MethodRefundFeeTariff = getOrderViewMoney(decoded.MethodRefundFeeTariff)
	m.PaysuperMethodRefundFeeTariffProfit = getOrderViewMoney(decoded.PaysuperMethodRefundFeeTariffProfit)
	m.PaysuperMethodRefundFixedFeeTariffSelfCost = getOrderViewMoney(decoded.PaysuperMethodRefundFixedFeeTariffSelfCost)
	m.MerchantRefundFixedFeeTariff = getOrderViewMoney(decoded.MerchantRefundFixedFeeTariff)
	m.PaysuperMethodRefundFixedFeeTariffProfit = getOrderViewMoney(decoded.PaysuperMethodRefundFixedFeeTariffProfit)
	m.RefundTaxFee = getOrderViewMoney(decoded.RefundTaxFee)
	m.RefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.RefundTaxFeeCurrencyExchangeFee)
	m.PaysuperRefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.PaysuperRefundTaxFeeCurrencyExchangeFee)
	m.RefundTaxFeeTotal = getOrderViewMoney(decoded.RefundTaxFeeTotal)
	m.RefundReverseRevenue = getOrderViewMoney(decoded.RefundReverseRevenue)
	m.RefundFeesTotal = getOrderViewMoney(decoded.RefundFeesTotal)
	m.RefundFeesTotalLocal = getOrderViewMoney(decoded.RefundFeesTotalLocal)
	m.PaysuperRefundTotalProfit = getOrderViewMoney(decoded.PaysuperRefundTotalProfit)
	m.Items = getOrderViewItems(decoded.Items)
	m.MccCode = decoded.MccCode
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.IsHighRisk = decoded.IsHighRisk
	m.RefundAllowed = decoded.RefundAllowed
	m.VatPayer = decoded.VatPayer
	m.IsProduction = decoded.IsProduction
	m.TaxRate = decoded.TaxRate
	m.MerchantInfo = decoded.MerchantInfo
	m.OrderChargeBeforeVat = decoded.OrderChargeBeforeVat

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.TransactionDate, err = ptypes.TimestampProto(decoded.TransactionDate)
	if err != nil {
		return err
	}

	return nil
}

func (m *OrderViewPublic) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOrderViewPublic)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Uuid = decoded.Uuid
	m.TotalPaymentAmount = decoded.TotalPaymentAmount
	m.Currency = decoded.Currency
	m.Project = getOrderProject(decoded.Project)
	m.Transaction = decoded.Transaction
	m.PaymentMethod = getPaymentMethodOrder(decoded.PaymentMethod)
	m.CountryCode = decoded.CountryCode
	m.MerchantId = decoded.MerchantId.Hex()
	m.Locale = decoded.Locale
	m.Status = decoded.Status
	m.User = decoded.User
	m.BillingAddress = decoded.BillingAddress
	m.Type = decoded.Type
	m.Issuer = decoded.Issuer
	m.MerchantPayoutCurrency = decoded.MerchantPayoutCurrency
	m.IsVatDeduction = decoded.IsVatDeduction
	m.ParentOrder = decoded.ParentOrder
	m.Cancellation = decoded.Cancellation

	if decoded.Refund != nil {
		m.Refund = &OrderNotificationRefund{
			Amount:        decoded.Refund.Amount,
			Currency:      decoded.Refund.Currency,
			Reason:        decoded.Refund.Reason,
			Code:          decoded.Refund.Code,
			ReceiptNumber: decoded.Refund.ReceiptNumber,
			ReceiptUrl:    decoded.Refund.ReceiptUrl,
		}
	}

	m.GrossRevenue = getOrderViewMoney(decoded.GrossRevenue)
	m.TaxFee = getOrderViewMoney(decoded.TaxFee)
	m.TaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.TaxFeeCurrencyExchangeFee)
	m.TaxFeeTotal = getOrderViewMoney(decoded.TaxFeeTotal)
	m.MethodFeeTotal = getOrderViewMoney(decoded.MethodFeeTotal)
	m.MethodFeeTariff = getOrderViewMoney(decoded.MethodFeeTariff)
	m.MethodFixedFeeTariff = getOrderViewMoney(decoded.MethodFixedFeeTariff)
	m.PaysuperFixedFee = getOrderViewMoney(decoded.PaysuperFixedFee)
	m.FeesTotal = getOrderViewMoney(decoded.FeesTotal)
	m.FeesTotalLocal = getOrderViewMoney(decoded.FeesTotalLocal)
	m.NetRevenue = getOrderViewMoney(decoded.NetRevenue)
	m.RefundGrossRevenue = getOrderViewMoney(decoded.RefundGrossRevenue)
	m.MethodRefundFeeTariff = getOrderViewMoney(decoded.MethodRefundFeeTariff)
	m.MerchantRefundFixedFeeTariff = getOrderViewMoney(decoded.MerchantRefundFixedFeeTariff)
	m.RefundTaxFee = getOrderViewMoney(decoded.RefundTaxFee)
	m.RefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.RefundTaxFeeCurrencyExchangeFee)
	m.PaysuperRefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.PaysuperRefundTaxFeeCurrencyExchangeFee)
	m.RefundReverseRevenue = getOrderViewMoney(decoded.RefundReverseRevenue)
	m.RefundFeesTotal = getOrderViewMoney(decoded.RefundFeesTotal)
	m.RefundFeesTotalLocal = getOrderViewMoney(decoded.RefundFeesTotalLocal)
	m.Items = getOrderViewItems(decoded.Items)
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.RefundAllowed = decoded.RefundAllowed
	m.OrderCharge = decoded.OrderCharge
	m.PaymentIpCountry = decoded.PaymentIpCountry
	m.IsIpCountryMismatchBin = decoded.IsIpCountryMismatchBin
	m.BillingCountryChangedByUser = decoded.BillingCountryChangedByUser
	m.VatPayer = decoded.VatPayer
	m.IsProduction = decoded.IsProduction

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.TransactionDate, err = ptypes.TimestampProto(decoded.TransactionDate)
	if err != nil {
		return err
	}

	return nil
}

func getPaymentMethodOrder(in *MgoOrderPaymentMethod) *PaymentMethodOrder {
	if in == nil {
		return nil
	}

	result := &PaymentMethodOrder{
		Id:              in.Id.Hex(),
		Name:            in.Name,
		ExternalId:      in.ExternalId,
		Params:          in.Params,
		PaymentSystemId: in.PaymentSystemId.Hex(),
		Group:           in.Group,
		Saved:           in.Saved,
		RefundAllowed:   in.RefundAllowed,
		Handler:         in.Handler,
	}

	if in.Card != nil {
		result.Card = in.Card
	}
	if in.Wallet != nil {
		result.Wallet = in.Wallet
	}
	if in.CryptoCurrency != nil {
		result.CryptoCurrency = in.CryptoCurrency
	}

	return result
}

func getOrderProject(in *MgoOrderProject) *ProjectOrder {
	project := &ProjectOrder{
		Id:                      in.Id.Hex(),
		MerchantId:              in.MerchantId.Hex(),
		UrlSuccess:              in.UrlSuccess,
		UrlFail:                 in.UrlFail,
		NotifyEmails:            in.NotifyEmails,
		SendNotifyEmail:         in.SendNotifyEmail,
		SecretKey:               in.SecretKey,
		UrlCheckAccount:         in.UrlCheckAccount,
		UrlProcessPayment:       in.UrlProcessPayment,
		UrlChargebackPayment:    in.UrlChargebackPayment,
		UrlCancelPayment:        in.UrlCancelPayment,
		UrlRefundPayment:        in.UrlRefundPayment,
		UrlFraudPayment:         in.UrlFraudPayment,
		CallbackProtocol:        in.CallbackProtocol,
		Status:                  in.Status,
		MerchantRoyaltyCurrency: in.MerchantRoyaltyCurrency,
		RedirectSettings:        in.RedirectSettings,
	}

	if len(in.Name) > 0 {
		project.Name = make(map[string]string)

		for _, v := range in.Name {
			project.Name[v.Lang] = v.Value
		}
	}

	return project
}

func getOrderViewMoney(in *OrderViewMoney) *OrderViewMoney {
	if in == nil {
		return &OrderViewMoney{}
	}

	return &OrderViewMoney{
		Amount:   tools.ToPrecise(in.Amount),
		Currency: in.Currency,
	}
}

func getOrderViewItems(in []*MgoOrderItem) []*OrderItem {
	var items []*OrderItem

	if len(in) <= 0 {
		return items
	}

	for _, v := range in {
		item := &OrderItem{
			Id:          v.Id.Hex(),
			Object:      v.Object,
			Sku:         v.Sku,
			Name:        v.Name,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
			Images:      v.Images,
			Url:         v.Url,
			Metadata:    v.Metadata,
			Code:        v.Code,
			PlatformId:  v.PlatformId,
		}

		item.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		item.CreatedAt, _ = ptypes.TimestampProto(v.UpdatedAt)

		items = append(items, item)
	}

	return items
}

func (m *Id) MarshalBSON() ([]byte, error) {
	st := &MgoId{}
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, errors.New(ErrorInvalidObjectId)
	}
	st.Id = oid
	return bson.Marshal(st)
}

func (m *Id) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoId)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	return nil
}

func (k *Key) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoKey)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	k.Id = decoded.Id.Hex()
	k.Code = decoded.Code
	k.KeyProductId = decoded.KeyProductId.Hex()
	k.PlatformId = decoded.PlatformId

	if decoded.OrderId != nil {
		k.OrderId = decoded.OrderId.Hex()
	}

	if k.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt); err != nil {
		return err
	}

	if k.RedeemedAt, err = ptypes.TimestampProto(decoded.RedeemedAt); err != nil {
		return err
	}

	if k.ReservedTo, err = ptypes.TimestampProto(decoded.ReservedTo); err != nil {
		return err
	}

	return nil
}

func (m *Key) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	keyProductOid, err := primitive.ObjectIDFromHex(m.KeyProductId)

	if err != nil {
		return nil, err
	}

	st := &MgoKey{
		Id:           oid,
		PlatformId:   m.PlatformId,
		KeyProductId: keyProductOid,
		Code:         m.Code,
	}

	if m.OrderId != "" {
		oid, err := primitive.ObjectIDFromHex(m.OrderId)

		if err != nil {
			return nil, err
		}

		orderId := oid
		st.OrderId = &orderId
	}

	if m.RedeemedAt != nil {
		if st.RedeemedAt, err = ptypes.Timestamp(m.RedeemedAt); err != nil {
			return nil, err
		}
	} else {
		st.RedeemedAt = time.Time{}
	}

	if m.ReservedTo != nil {
		if st.ReservedTo, err = ptypes.Timestamp(m.ReservedTo); err != nil {
			return nil, err
		}
	} else {
		st.ReservedTo = time.Time{}
	}

	if m.CreatedAt != nil {
		if st.CreatedAt, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return nil, err
		}
	} else {
		st.CreatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *OperatingCompany) MarshalBSON() ([]byte, error) {
	st := &MgoOperatingCompany{
		Name:               m.Name,
		Country:            m.Country,
		RegistrationNumber: m.RegistrationNumber,
		RegistrationDate:   m.RegistrationDate,
		VatNumber:          m.VatNumber,
		Email:              m.Email,
		Address:            m.Address,
		VatAddress:         m.VatAddress,
		SignatoryName:      m.SignatoryName,
		SignatoryPosition:  m.SignatoryPosition,
		BankingDetails:     m.BankingDetails,
		PaymentCountries:   m.PaymentCountries,
	}
	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	var err error

	if m.CreatedAt != nil {
		if st.CreatedAt, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return nil, err
		}
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		if st.UpdatedAt, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return nil, err
		}
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *OperatingCompany) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOperatingCompany)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Name = decoded.Name
	m.Country = decoded.Country
	m.RegistrationNumber = decoded.RegistrationNumber
	m.RegistrationDate = decoded.RegistrationDate
	m.VatNumber = decoded.VatNumber
	m.Email = decoded.Email
	m.Address = decoded.Address
	m.VatAddress = decoded.VatAddress
	m.SignatoryName = decoded.SignatoryName
	m.SignatoryPosition = decoded.SignatoryPosition
	m.BankingDetails = decoded.BankingDetails
	m.PaymentCountries = decoded.PaymentCountries

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *PaymentMinLimitSystem) MarshalBSON() ([]byte, error) {
	st := &MgoPaymentMinLimitSystem{
		Currency: m.Currency,
		Amount:   m.Amount,
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	var err error

	if m.CreatedAt != nil {
		if st.CreatedAt, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return nil, err
		}
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		if st.UpdatedAt, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return nil, err
		}
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *PaymentMinLimitSystem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoPaymentMinLimitSystem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Currency = decoded.Currency
	m.Amount = decoded.Amount

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserRole) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	st := &MgoUserRole{
		Id:        oid,
		Role:      m.Role,
		Status:    m.Status,
		Email:     m.Email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}

	oid, err = primitive.ObjectIDFromHex(m.MerchantId)

	if err == nil {
		hex := oid
		st.MerchantId = &hex
	}

	oid, err = primitive.ObjectIDFromHex(m.UserId)

	if err == nil {
		hex := oid
		st.UserId = &hex
	}

	if m.CreatedAt != nil {
		if st.CreatedAt, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return nil, err
		}
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		if st.UpdatedAt, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return nil, err
		}
	} else {
		st.UpdatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (k *UserRole) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoUserRole)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	k.Id = decoded.Id.Hex()
	k.Role = decoded.Role
	k.Status = decoded.Status
	k.Email = decoded.Email
	k.FirstName = decoded.FirstName
	k.LastName = decoded.LastName

	if decoded.MerchantId != nil {
		k.MerchantId = decoded.MerchantId.Hex()
	}

	if decoded.UserId != nil {
		k.UserId = decoded.UserId.Hex()
	}

	if k.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt); err != nil {
		return err
	}

	if k.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt); err != nil {
		return err
	}

	return nil
}

type CountryAndRegionItem struct {
	Country string `bson:"iso_code_a2"`
	Region  string `bson:"region"`
}

type CountryAndRegionItems struct {
	Items []*CountryAndRegionItem `json:"items"`
}
