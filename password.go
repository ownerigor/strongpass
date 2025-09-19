package main

import (
	"crypto/rand"
	"math/big"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?")

func randomInt(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

func GeneratePassword(length int) string {
	var password []rune
	for i := 0; i < length; i++ {
		idx := randomInt(int64(len(charset)))
		password = append(password, charset[idx])
	}
	return string(password)
}
