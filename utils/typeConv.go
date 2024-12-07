package utils

import "strconv"

func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		ErrorResponse(err)
	}
	return num
}
