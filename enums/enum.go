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

var orderStatusText = map[int]string{
	CANCELLED:    "CANCELLED",
	COMPLETED:    "COMPLETED",
	ONPROCESSING: "ONPROCESSING",
}

func OrderStatusText(code int) string {
	return orderStatusText[code]
}

func OrderStatusList() []int {
	return getKeys(appointmentStatusText)
}

var orderStatusCode = map[string]int{
	"CANCELLED":    CANCELLED,
	"COMPLETED":    COMPLETED,
	"ONPROCESSING": ONPROCESSING,
}

func OrderStatusCode(text string) int {
	return orderStatusCode[text]
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
