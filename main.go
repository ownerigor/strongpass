package main

import "fmt"

func main() {
	password := GeneratePassword(16)
	fmt.Printf("Generated secure password: %s", password)
}
