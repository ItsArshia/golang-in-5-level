package calculation

import (
	"fmt"
	"strconv"
)

func Calculate(numbers []string, operators []string) (float64, error) {
	if len(numbers) == 0 || len(numbers) != len(operators)+1 {
		return 0, fmt.Errorf("invalid expression")
	}

	// store first value as first
	first, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		return 0, err
	}

	total := 0.0
	term := first

	for i, op := range operators {
		next, err := strconv.ParseFloat(numbers[i+1], 64)
		if err != nil {
			return 0, err
		}

		switch op {
		case "+":
			total += term
			term = next

		case "-":
			total += term
			term = -next

		case "*":
			term *= next

		case "/":
			if next == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			term /= next

		default:
			return 0, fmt.Errorf("unknown operator: %s", op)
		}
	}

	return total + term, nil
}