package utils

import (
	"crypto/rand"
)

func CreateOtpCode(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	for i := range length {
		buffer[i] = (buffer[i] % 10) + '0'
	}
	return string(buffer), nil
}
