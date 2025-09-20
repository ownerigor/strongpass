package generator

import (
	"testing"
)

func TestPasswordLength(t *testing.T) {
	length := 20
	password := GeneratePassword(length, true, true)

	if len(password) != length {
		t.Errorf("expected length %d, got %d", length, len(password))
	}
}

func TestPasswordCharset(t *testing.T) {
	length := 50
	password := GeneratePassword(length, true, true)

	charset := append([]rune{}, Lowercase...)
	charset = append(charset, Uppercase...)
	charset = append(charset, Numbers...)
	charset = append(charset, Symbols...)

	for _, char := range password {
		if !containsRune(charset, char) {
			t.Errorf("found invalid character: %q", char)
		}
	}
}

func containsRune(set []rune, r rune) bool {
	for _, c := range set {
		if c == r {
			return true
		}
	}
	return false
}

func TestPasswordRandomness(t *testing.T) {
	pass1 := GeneratePassword(16, true, true)
	pass2 := GeneratePassword(16, true, true)

	if pass1 == pass2 {
		t.Errorf("expected different passwords, got the same: %s", pass1)
	}
}

func TestPasswordMinLength(t *testing.T) {
	pass := GeneratePassword(1, true, true)
	if len(pass) != 1 {
		t.Errorf("expected length 1, got %d", len(pass))
	}
}

func TestPasswordLargeLength(t *testing.T) {
	length := 1000
	pass := GeneratePassword(length, true, true)
	if len(pass) != length {
		t.Errorf("expected length %d, got %d", length, len(pass))
	}
}
