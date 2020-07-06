package enums

import "sort"

const (
	CANCELLED           = 3
	COMPLETED           = 4
	CANCELLEDNOTARRIVED = 5
	ONPROCESSING        = 6
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
