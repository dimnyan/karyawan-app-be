package utils

import "time"

func ParseDate(dateParam string) time.Time {
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		ErrorMessage("Cannot parse date")
		return time.Time{}
	}
	return date
}

func ParseDateTime(dateTimeParam string) int64 {
	dateTime, err := time.Parse("2006-01-02 15:04 MST", dateTimeParam)
	if err != nil {
		ErrorMessage("Cannot parse date time")
		return 0
	}
	return dateTime.Unix()
}
