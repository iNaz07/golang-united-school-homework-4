package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
//
//func main() {
//	fmt.Println(StringSum("3+5"))
//	fmt.Println(StringSum("+ 35 + 54 +"))
//	fmt.Println(StringSum("-3-5"))
//	fmt.Println(StringSum("-5"))
//	fmt.Println(StringSum("--3+5"))
//	fmt.Println(StringSum("3+5+"))
//	fmt.Println(StringSum("   "))
//	fmt.Println(StringSum(""))
//}

func StringSum(input string) (output string, err error) {

	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("invalid input: %w", errorEmptyInput)
	}
	params := []string{}
	m, n := 0, 0
	for i, char := range input {
		n++
		if i != 0 && (char == 43 || char == 45) {
			params = append(params, input[m:n-1])
			m = i
		} else if n >= len(input) {
			params = append(params, input[m:])
			break
		}
	}

	if len(params) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if len(params) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	a, err := strconv.Atoi(strings.Trim(params[0], " "))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	b, err := strconv.Atoi(strings.Trim(params[1], " "))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	res := a + b
	output = strconv.Itoa(res)

	return output, nil
}
