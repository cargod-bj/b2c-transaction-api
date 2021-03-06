package enums

import "sort"

const (
	CANCELLED           = 1
	COMPLETED           = 2
	CANCELLEDNOTARRIVED = 3
	ONPROCESSING        = 4
)

var appointmentStatusText = map[int]string{
	CANCELLED:           "CANCELLED",
	COMPLETED:           "COMPLETED",
	CANCELLEDNOTARRIVED: "CANCELLEDNOTARRIVED",
	ONPROCESSING:        "ONPROCESSING",
}

func AppointmentStatusText(code int) string {
	return appointmentStatusText[code]
}

func AppointmentStatusList() []int {
	return getKeys(appointmentStatusText)
}

func getKeys(maps map[int]string) []int {
	keys := make([]int, 0, len(maps))
	for mapKey, _ := range maps {
		keys = append(keys, mapKey)
	}
	sort.Ints(keys)

	return keys
}

var appointmentStatusCode = map[string]int{
	"CANCELLED":           CANCELLED,
	"COMPLETED":           COMPLETED,
	"CANCELLEDNOTARRIVED": CANCELLEDNOTARRIVED,
	"ONPROCESSING":        ONPROCESSING,
}

func AppointmentStatusCode(text string) int {
	return appointmentStatusCode[text]
}

const (
	LEAD   = 1
	WALKIN = 2
)

var leadTypeText = map[int]string{
	LEAD:   "LEAD",
	WALKIN: "WALKIN",
}

func LeadTypeText(code int) string {
	return leadTypeText[code]
}

func LeadTypeList() []int {
	return getKeys(leadTypeText)
}

var leadTypeCode = map[string]int{
	"LEAD":   LEAD,
	"WALKIN": WALKIN,
}

func LeadTypeCode(text string) int {
	return leadTypeCode[text]
}

const (
	PendingPay     = 1
	PendingDeliver = 2
	OrderCompleted = 3
	OrderClosed    = 4
	MoneyBack      = 5
	OrderCancelled = 6
	OrderCreated   = 99
)

var orderStatusText = map[int]string{
	PendingPay:     "Pending Pay",
	PendingDeliver: "Pending Deliver",
	OrderCompleted: "Order Completed",
	OrderClosed:    "Order Closed",
	MoneyBack:      "Money Back",
	OrderCancelled: "Order Cancelled",
	OrderCreated:   "Order Created",
}

func OrderStatusText(code int) string {
	return orderStatusText[code]
}

func OrderStatusList() []int {
	return getKeys(orderStatusText)
}

var orderStatusCode = map[string]int{
	"Pending Pay":     PendingPay,
	"Pending Deliver": PendingDeliver,
	"Order Completed": OrderCompleted,
	"Order Closed":    OrderClosed,
	"Money Back":      MoneyBack,
	"Order Cancelled": OrderCancelled,
	"Order Created":   OrderCreated,
}

func OrderStatusCode(text string) int {
	return orderStatusCode[text]
}

const (
	Logistics  = 1
	SelfPickUp = 2
)

var deliveryMethodText = map[int]string{
	Logistics:  "Logistics",
	SelfPickUp: "Self Pick-up",
}

func DeliveryMethodText(code int) string {
	return deliveryMethodText[code]
}

func DeliveryMethodList() []int {
	return getKeys(orderStatusText)
}

var deliveryMethodCode = map[string]int{
	"Logistics":    Logistics,
	"Self Pick-up": SelfPickUp,
}

func DeliveryMethodCode(text string) int {
	return deliveryMethodCode[text]
}

const (
	CarPrice  = 1
	Insurance = 2
	WarrantyA = 3
	WarrantyB = 4
)

var feeTypeText = map[int]string{
	CarPrice:  "CarPrice",
	Insurance: "Insurance",
	WarrantyA: "Warranty-A",
	WarrantyB: "Warranty-B",
}

func FeeTypeText(code int) string {
	return feeTypeText[code]
}

func FeeTypeList() []int {
	return getKeys(feeTypeText)
}

var feeTypeCode = map[string]int{
	"CarPrice":   CarPrice,
	"Insurance":  Insurance,
	"Warranty-A": WarrantyA,
	"Warranty-B": WarrantyB,
}

func FeeTypeCode(text string) int {
	return feeTypeCode[text]
}

const (
	BookingFee  = 1
	DownPayment = 2
	LoanPayment = 3
	FullPayment = 4
)

var paymentTypeText = map[int]string{
	BookingFee:  "Booking Fee",
	DownPayment: "Down Payment",
	LoanPayment: "Loan Payment",
	FullPayment: "Full Payment",
}

func PaymentTypeText(code int) string {
	return paymentTypeText[code]
}

func PaymentTypeList() []int {
	return getKeys(paymentTypeText)
}

var paymentTypeCode = map[string]int{
	"Booking Fee":  BookingFee,
	"Down Payment": DownPayment,
	"Loan Payment": LoanPayment,
	"Full Payment": FullPayment,
}

func PaymentTypeCode(text string) int {
	return paymentTypeCode[text]
}

const (
	Cash           = 1
	Cheque         = 2
	BankTransfer   = 3
	Loan           = 4
	DiscountCoupon = 5
	TradeIn        = 6
	CardPayment    = 7
	Online         = 8
)

var paymentMethodText = map[int]string{
	Cash:           "Cash",
	Cheque:         "Cheque",
	BankTransfer:   "Bank Transfer",
	Loan:           "Loan",
	DiscountCoupon: "Discount Coupon",
	TradeIn:        "TradeIn",
	CardPayment:    "CardPayment",
	Online:         "Online",
}

func PaymentMethodText(code int) string {
	return paymentMethodText[code]
}

func PaymentMethodList() []int {
	return getKeys(paymentMethodText)
}

var paymentMethodCode = map[string]int{
	"Cash":            Cash,
	"Cheque":          Cheque,
	"Bank Transfer":   BankTransfer,
	"Loan":            Loan,
	"Discount Coupon": DiscountCoupon,
	"TradeIn":         TradeIn,
	"CardPayment":     CardPayment,
	"Online":          Online,
}

func PaymentMethodCode(text string) int {
	return paymentMethodCode[text]
}

const (
	Payment = 1
	Refund  = 2
)

var recordTypeText = map[int]string{
	Payment: "Payment",
	Refund:  "Refund",
}

func RecordTypeText(code int) string {
	return recordTypeText[code]
}

func RecordTypeList() []int {
	return getKeys(recordTypeText)
}

var recordTypeCode = map[string]int{
	"Payment": Payment,
	"Refund":  Refund,
}

func RecordTypeCode(text string) int {
	return recordTypeCode[text]
}

const (
	SalesContract      = 1
	PUSPAKOMInspection = 2
	JPJVehicleTransfer = 3
	insurance          = 4
	RoadTax            = 5
	Warranty           = 6
	ReceivedConfirm    = 7
	Moneyback          = 8
)

var checkListTypeText = map[int]string{
	SalesContract:      "SalesContract",
	PUSPAKOMInspection: "PUSPAKOMInspection",
	JPJVehicleTransfer: "JPJVehicleTransfer",
	insurance:          "Insurance",
	RoadTax:            "RoadTax",
	Warranty:           "Warranty",
	ReceivedConfirm:    "ReceivedConfirm",
	Moneyback:          "Moneyback",
}

func CheckListTypeText(code int) string {
	return checkListTypeText[code]
}

func CheckListTypeList() []int {
	return getKeys(checkListTypeText)
}

var checkListTypeCode = map[string]int{
	"SalesContract":      SalesContract,
	"PUSPAKOMInspection": PUSPAKOMInspection,
	"JPJVehicleTransfer": JPJVehicleTransfer,
	"Insurance":          insurance,
	"RoadTax":            RoadTax,
	"Warranty":           Warranty,
	"ReceivedConfirm":    ReceivedConfirm,
	"Moneyback":          Moneyback,
}

func CheckListTypeCode(text string) int {
	return checkListTypeCode[text]
}

const (
	CouponTypeCash          = 1
	CouponTypeFullReduction = 2
	CouponTypeDiscount      = 3
)

var couponTypeText = map[int]string{
	CouponTypeCash:          "Cash",
	CouponTypeFullReduction: "FullReduction",
	CouponTypeDiscount:      "Discount",
}

func CouponTypeText(code int) string {
	return couponTypeText[code]
}

func CouponTypeList() []int {
	return getKeys(couponTypeText)
}

var CouponTypeCode = map[string]int{
	"Cash":          CouponTypeCash,
	"FullReduction": CouponTypeFullReduction,
	"Discount":      CouponTypeDiscount,
}

const (
	C2BTradeIn             = 1
	B2CClientReferral      = 2
	OnlineShopVoucherShop  = 3
	CarSomeEmployeeVoucher = 4
	BusinessCampaign       = 5
)

var couponBizTypeText = map[int]string{
	C2BTradeIn:             "C2B Trade-in",
	B2CClientReferral:      "B2C Client Referral",
	OnlineShopVoucherShop:  "Online Shop Voucher_Shopee&Lazada",
	CarSomeEmployeeVoucher: "CarSome Employee Voucher",
	BusinessCampaign:       "Business Campaign",
}

func CouponBizTypeText(code int) string {
	return couponBizTypeText[code]
}

var couponBizTypeCodeText = map[int]string{
	C2BTradeIn:             "T",
	B2CClientReferral:      "C",
	OnlineShopVoucherShop:  "O",
	CarSomeEmployeeVoucher: "E",
	BusinessCampaign:       "B",
}

func CouponBizTypeCodeText(code int) string {
	return couponBizTypeCodeText[code]
}

func CouponBizTypeList() []int {
	return getKeys(couponBizTypeText)
}

var CouponBizTypeCode = map[string]int{
	"C2B Trade-in":                      C2BTradeIn,
	"B2C Client Referral":               B2CClientReferral,
	"Online Shop Voucher_Shopee&Lazada": OnlineShopVoucherShop,
	"CarSome Employee Voucher":          CarSomeEmployeeVoucher,
	"Business Campaign":                 BusinessCampaign,
}

const (
	Active   = 1
	InActive = 2
)

var statusText = map[int]string{
	Active:   "Active",
	InActive: "InActive",
}

func StatusText(code int) string {
	return statusText[code]
}

func StatusList() []int {
	return getKeys(statusText)
}

var StatusCode = map[string]int{
	"Active":   Active,
	"InActive": InActive,
}

const (
	UnUsed = 1
	Used   = 2
)

var couponUsedStatusText = map[int]string{
	UnUsed: "UnUsed",
	Used:   "Used",
}

func CouponUsedStatusText(code int) string {
	return couponUsedStatusText[code]
}

func CouponUsedStatusList() []int {
	return getKeys(couponUsedStatusText)
}

var CouponUsedStatusCode = map[string]int{
	"UnUsed": UnUsed,
	"Used":   Used,
}

const (
	HomeTestDrive   = 1
	RetailTestDrive = 2
)

var testDriveTypeText = map[int]string{
	HomeTestDrive:   "Home Test Drive",
	RetailTestDrive: "Retail Test Drive",
}

func TestDriveTypeText(code int) string {
	return testDriveTypeText[code]
}

func TestDriveTypeList() []int {
	return getKeys(couponUsedStatusText)
}

var TestDriveTypeCode = map[string]int{
	"Home Test Drive":   HomeTestDrive,
	"Retail Test Drive": RetailTestDrive,
}

const (
	ChangeCar                             = 1
	RejectedByBank                        = 2
	UnableToProvideDocumentsRequestByBank = 3
	HigherInterestRate                    = 4
	NotEnoughDownPayment                  = 5
	ChangeMind                            = 6
	OtherReasons                          = 7
	OnlinePaymentTimeout                  = 8
	CustomerCancel                        = 9
	BookingFailedCancel                   = 10
)

var cancelReasonText = map[int]string{
	ChangeCar:                             "Change car",
	RejectedByBank:                        "Rejected by bank",
	UnableToProvideDocumentsRequestByBank: "Unable to provide documents request by bank",
	HigherInterestRate:                    "Higher interest rate",
	NotEnoughDownPayment:                  "Not enough Down Payment",
	ChangeMind:                            "Change mind",
	OtherReasons:                          "Other reasons",
	OnlinePaymentTimeout:                  "Online Payment Timeout",
	CustomerCancel:                        "Customer Cancel",
	BookingFailedCancel:                   "Booking Failed Cancel",
}

func CancelReasonText(code int) string {
	return cancelReasonText[code]
}

func CancelReasonList() []int {
	return getKeys(cancelReasonText)
}

var CancelReasonCode = map[string]int{
	"Change car":       ChangeCar,
	"Rejected by bank": RejectedByBank,
	"Unable to provide documents request by bank": UnableToProvideDocumentsRequestByBank,
	"Higher interest rate":                        HigherInterestRate,
	"Not enough Down Payment":                     NotEnoughDownPayment,
	"Change mind":                                 ChangeMind,
	"Other reasons":                               OtherReasons,
	"Online Payment Timeout":                      OnlinePaymentTimeout,
	"Customer Cancel":                             CustomerCancel,
	"Booking Failed Cancel":                       BookingFailedCancel,
}
