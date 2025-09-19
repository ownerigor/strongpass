package main

import (
	"flag"
	"fmt"
)

func main() {
	length := flag.Int("length", 16, "Length of the password")
	useNumbers := flag.Bool("numbers", true, "Include numbers in the password")
	useSymbols := flag.Bool("symbols", true, "Include symbols in the password")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	for i := 0; i < *count; i++ {
		password := GeneratePassword(*length, *useNumbers, *useSymbols)
		fmt.Println(password)
	}
}
