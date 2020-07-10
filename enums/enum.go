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
)

var orderStatusText = map[int]string{
	PendingPay:     "Pending Pay",
	PendingDeliver: "Pending Deliver",
	OrderCompleted: "Order Completed",
	OrderClosed:    "Order Closed",
	MoneyBack:      "Money Back",
	OrderCancelled: "Order Cancelled",
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
	Earnest        = 1
	DownPayment    = 2
	BalancePayment = 3
	FullPayment    = 4
)

var paymentTypeText = map[int]string{
	Earnest:        "Earnest",
	DownPayment:    "Down Payment",
	BalancePayment: "Balance Payment",
	FullPayment:    "Full Payment",
}

func PaymentTypeText(code int) string {
	return paymentTypeText[code]
}

func PaymentTypeList() []int {
	return getKeys(paymentTypeText)
}

var paymentTypeCode = map[string]int{
	"Earnest":         Earnest,
	"Down Payment":    DownPayment,
	"Balance Payment": BalancePayment,
	"Full Payment":    FullPayment,
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
)

var paymentMethodText = map[int]string{
	Cash:           "Cash",
	Cheque:         "Cheque",
	BankTransfer:   "Bank Transfer",
	Loan:           "Loan",
	DiscountCoupon: "Discount Coupon",
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
