package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string) (string, error) {
	salt := []byte("poopoo peepee")

	hashedPass, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)

	encodedStr := base64.StdEncoding.EncodeToString(hashedPass)

	return encodedStr, err
}
