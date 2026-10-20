package main

import (
	"bufio"
	"calculator/calculation"
	"calculator/serialize"
	"fmt"
	"os"
)

type Human struct {
	name string
	age  int8
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Math Shell")
	fmt.Println("------------------------------------------------------")

	for {
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Input error !")
			return
		}

		numbers, opreators, err := serialize.ParseInput(input)

		result, err := calculation.Calculate(numbers, opreators)

		if err != nil {
			fmt.Println("Error: Invalid Expression")
		}

		fmt.Println("= ", result)
		
	}
}
