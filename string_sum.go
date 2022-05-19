package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	arr := make([]rune, 0)
	operands := 0
	for _, v := range input {
		val := string(v)
		if !strings.Contains("0123456789+- ", val) {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
		if v != ' ' {
			arr = append(arr, v)
		}
		if v == '+' || v == '-' {
			operands++
		}
	}

	i := 0
	if arr[0] == '+' || arr[0] == '-' {
		i = 1
	}

	for i < len(arr) && !(arr[i] == '+' || arr[i] == '-') {
		i++
	}

	first, ok := strconv.ParseInt(string(arr[:i]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	second, ok := strconv.ParseInt(string(arr[i:]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	return strconv.FormatInt(first+second, 10), nil
}
