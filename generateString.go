package randomStrings

import (
	"fmt"
	"math/rand"
)

func GenerateRandomString(lengthOfString int) (string, error) {

	if lengthOfString <= 0 {
		return "", fmt.Errorf("Length of the String to generate Random number has to be positive number")
	}
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, lengthOfString)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s), nil
}
