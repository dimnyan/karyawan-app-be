package utils

import "time"

func ParseDate(dateParam string) time.Time {
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		ErrorResponse(err)
	}
	return date
}
