package utils

import "time"

func ExtractDate(dateTime time.Time) string {
	var dateStr string
	dateStr = dateTime.String()
	dateStr = dateStr[0:10]

	return dateStr
}
