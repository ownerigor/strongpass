package generator

import (
	"crypto/rand"
	"math/big"
)

var (
	Lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	Uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Numbers   = []rune("0123456789")
	Symbols   = []rune("!@#$%^&*()_+-=[]{}|;:,.<>?")
)

func randomInt(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

func GeneratePassword(length int, useNumbers bool, useSymbols bool) string {
	charset := append([]rune{}, Lowercase...)
	charset = append(charset, Uppercase...)

	if useNumbers {
		charset = append(charset, Numbers...)
	}

	if useSymbols {
		charset = append(charset, Symbols...)
	}

	password := make([]rune, length)
	for i := 0; i < length; i++ {
		idx := randomInt(int64(len(charset)))
		password[i] = charset[idx]
	}
	return string(password)
}
