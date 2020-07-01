package enums

import "sort"

const (
	CANCELLED           = 1
	COMPLETED           = 2
	CANCELLEDNOTARRIVED = 3
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
	ORDER_CREATED   = 1
	ORDER_EARNEST   = 2
	ORDER_PAID      = 3
	ORDER_COMPLETED = 4
	ORDER_CANCELED  = 20
)

var orderStatusText = map[int]string{
	ORDER_CREATED:   "CANCELLED",
	ORDER_EARNEST:   "EARNEST",
	ORDER_PAID:      "PAID",
	ORDER_COMPLETED: "COMPLETED",
	ORDER_CANCELED:  "CANCELED",
}

func OrderStatusText(code int) string {
	return orderStatusText[code]
}

func OrderStatusList() []int {
	return getKeys(orderStatusText)
}

var orderStatusCode = map[string]int{
	"CANCELLED": ORDER_CREATED,
	"EARNEST":   ORDER_EARNEST,
	"PAID":      ORDER_PAID,
	"COMPLETED": ORDER_COMPLETED,
	"CANCELED":  ORDER_CANCELED,
}

func OrderStatusCode(text string) int {
	return orderStatusCode[text]
}
