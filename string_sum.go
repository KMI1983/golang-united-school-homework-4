package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorNotValidValue  = errors.New("input expression is not valid")
)

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) == 0 {
		return "", fmt.Errorf("Empty input error: %w", errorEmptyInput)
	}
	match, _ := regexp.MatchString(`^\D*[0-9]+\D+[0-9]+\D*$`, input)
	if !match {
		return "", fmt.Errorf("Not Two Operands error: %w", errorNotTwoOperands)
	}

	match, _ = regexp.MatchString(`^([\+\-]?[0-9]+[\+\-]{1}[0-9]+)$`, input)
	if !match {

		_, err := strconv.ParseInt(input, 10, 32)
		return "", fmt.Errorf("Incorrect input error: %w", err)
	}

	firstOperator := ""
	secondOperator := ""
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "+" || string(input[i]) == "-" {
			if i == 0 {
				firstOperator = string(input[i])
			} else {
				secondOperator = string(input[i])
			}
		}
	}

	reg := regexp.MustCompile(`[\+\-]`)
	operands := reg.Split(input, -1)
	var correctOperands []string
	for i := 0; i < len(operands); i++ {
		if operands[i] != "" {
			correctOperands = append(correctOperands, operands[i])
		}
	}

	firstOperand, err := strconv.Atoi(correctOperands[0])
	if err != nil {
		return "", fmt.Errorf("Parsing first operand error: %w", err)
	}

	secondOperand, err := strconv.Atoi(correctOperands[1])
	if err != nil {
		return "", fmt.Errorf("Parsing second operand error: %w", err)
	}

	if firstOperator == "-" {
		firstOperand = firstOperand * -1
	}
	sum := 0
	if secondOperator == "+" {
		sum = firstOperand + secondOperand
	} else {
		sum = firstOperand - secondOperand
	}

	return strconv.Itoa(sum), nil
}
