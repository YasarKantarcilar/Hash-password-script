package decodePassword

import (
	HashPassword "hashAlgorithm/HashPassword"
	"slices"
)

// DecodePassword takes a password and returns the decoded password
func DecodePassword(password string) string {
	if slices.Contains([]byte(HashPassword.Alph), 'a') {
		return "THIS IS NOT A HASHED PASSWORD"
	}
	passwordBytes := []byte(password)
	for i := 0; i < len(passwordBytes); i++ {
		for j := 0; j < len(HashPassword.Alph); j++ {
			if passwordBytes[i] == HashPassword.Alph[j] {
				passwordBytes[i] = HashPassword.Alph[j-1]
				break
			}
		}
	}

	return string(passwordBytes)
}
