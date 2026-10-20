package serialize

import (
	"fmt"
	"strings"
	"unicode"
	"errors"
)

// got a input and returns two array of numbers and opreators in this order -> numbers, opreators
func ParseInput(input string) ([]string, []string, error) {

	if len(input) < 3 {
		err := errors.New("Error: Please Enter Your Expresion")
		fmt.Println(err)
		return nil, nil, err
	}
	var numbers []string
	var operators []string
	var number strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			number.WriteRune(char)
			continue
		}

		if number.Len() > 0 {
			numbers = append(numbers, number.String())
			number.Reset()
		}

		if unicode.IsSpace(char) {
			continue
		}

		// if its not one of this states its a opreator
		operators = append(operators, string(char))
	}

	if number.Len() > 0 {
		numbers = append(numbers, number.String())
	}

	if len(numbers) < 2 {
		err := errors.New("Error: please enter at least two number")
		fmt.Println(err)
	}

	return numbers, operators, nil
}
