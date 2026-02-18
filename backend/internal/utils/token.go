package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateSessionToken(length int) (string, error) {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	token := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)

	return token, nil
}