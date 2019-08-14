package main

import (
	"fmt"
	"stark/gophercises/normalize/normalize"
)

func main() {
	numbers := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, num := range numbers {
		fmt.Println(normalize.Normalize(num))
	}
}
