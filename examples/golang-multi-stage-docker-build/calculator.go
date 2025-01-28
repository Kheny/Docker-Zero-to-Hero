package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi Abhishek.Veeramalla, I am a calculator app ....")

	for {
		// Read input from the user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter any calculation (Example: 1 + 2 (or) 2 * 5 -> Please maintain spaces as shown in example): ")
		text, _ := reader.ReadString('\n')

		// Trim the newline character from the input
		text = strings.TrimSpace(text)

		// Check if the user entered "exit" to quit the program
		if text == "exit" {
			break
		}

		// Split the input into two parts: the left operand, operator and the right operand
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Convert the operands to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Perform the calculation based on the operator
		result, err := calculate(left, parts[1], right)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Print the result
		fmt.Printf("Result: %d\n", result)
	}
}

func calculate(left int, operator string, right int) (int, error) {
	switch operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return left / right, nil
	default:
		return 0, errors.New("invalid operator")
	}
}