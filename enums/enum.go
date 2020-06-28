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
