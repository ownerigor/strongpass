package main

import (
	"flag"
	"fmt"

	"github.com/ownerigor/strongpass/internal/generator"
)

func main() {
	length := flag.Int("length", 16, "Length of the password")
	useNumbers := flag.Bool("numbers", false, "Include numbers in the password")
	useSymbols := flag.Bool("symbols", false, "Include symbols in the password")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	for i := 0; i < *count; i++ {
		password := generator.GeneratePassword(*length, *useNumbers, *useSymbols)
		fmt.Println(password)
	}
}
