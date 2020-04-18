package helpers

import (
	"time"
)

func ConvertToTime(dates string) time.Time {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = dates + " 00:00:00"
	date, _ = time.Parse(layoutFormat, value)

	return date
}