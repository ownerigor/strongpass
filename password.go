package main

import (
	"crypto/rand"
	"math/big"
)

var (
	lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers   = []rune("0123456789")
	symbols   = []rune("!@#$%^&*()_+-=[]{}|;:,.<>?")
)

func randomInt(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

func GeneratePassword(length int, useNumbers bool, useSymbols bool) string {
	charset := append([]rune{}, lowercase...)
	charset = append(charset, uppercase...)

	if useNumbers {
		charset = append(charset, numbers...)
	}

	if useSymbols {
		charset = append(charset, symbols...)
	}

	password := make([]rune, length)
	for i := 0; i < length; i++ {
		idx := randomInt(int64(len(charset)))
		password[i] = charset[idx]
	}
	return string(password)
}
