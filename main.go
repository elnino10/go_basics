package main

import (
	"errors"
	"fmt"
)

func main() {

}

func getUserInput(prompt string) (string, error) {
	var input string

	fmt.Print(prompt)
	fmt.Scan(&input)

	if input == "" {
		return "", errors.New("Invalid input")
	}

	return input, nil
}
