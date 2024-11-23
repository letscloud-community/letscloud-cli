package helpers

import (
	"crypto/rand"
	"log"
	"math/big"
)

// GenerateRandomPassword generates a secure random password with alphanumeric characters
func GenerateRandomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=<>?"
	const passwordLength = 16

	password := make([]byte, passwordLength)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Printf("Error generating password: %v", err)
			return "default-password"
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password)
}
