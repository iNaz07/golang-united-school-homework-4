package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function 35 + 46
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

// func main() {
// 	fmt.Println(StringSum("3+5"))
// 	fmt.Println(StringSum("35 + 54"))
// 	fmt.Println(StringSum("-3+5"))
// 	fmt.Println(StringSum("-3-56"))
// 	fmt.Println(StringSum("--3+5"))
// 	fmt.Println(StringSum("3+5+"))
// 	fmt.Println(StringSum("   "))
// 	fmt.Println(StringSum(""))
// }

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", errorEmptyInput
	}
	var first, second, op1, op2 string
	for i, char := range input {
		if char == 32 {
			continue
		}
		if i == 0 && (char == 45 || char == 43) {
			op1 = string(char)
		} else if op2 == "" {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				if char == 45 || char == 43 {
					if op1 != "" && first == "" {
						return "", errorNotTwoOperands
					}
					op2 = string(char)
					continue
				}
				return "", fmt.Errorf("invalid input := %w", err)
			}
			first += string(char)
		} else {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				return "", fmt.Errorf("invalid input := %w", err)
			}
			second += string(char)
		}

	}
	if first == "" && second == "" && op1 == "" && op2 == "" {
		return "", errorEmptyInput
	}
	if op2 != "+" && op2 != "-" {
		return "", errorNotTwoOperands
	} else {
		a, err := strconv.Atoi(op1 + first)
		if err != nil {
			return "", fmt.Errorf("invalid input := %w", err)
		}
		b, err := strconv.Atoi(op2 + second)
		if err != nil {
			return "", fmt.Errorf("invalid input := %w", err)
		}
		output = strconv.Itoa(a + b)
	}
	return output, nil
}
