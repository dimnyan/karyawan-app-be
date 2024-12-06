package utils

import (
	"fmt"
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(length int) string {
	var word string

	k := len(alphabet)

	for i := 0; i < 6; i++ {
		char := alphabet[rand.Intn(k)]
		word = fmt.Sprintf("%s%c", word, char)
	}

	return word
}
