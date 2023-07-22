package leetcode

import (
	"strconv"
)

func reverseInteger(x int) int {
	if (x) == 0 || x > 2147483647 || x < -2147483647 {
		return 0
	}

	neg := false

	if x < 0 {
		x = x * -1
		neg = true
	}

	xStr := strconv.Itoa(x)

	i := len(xStr) - 1
	compiledStr := ""

	for i >= 0 {
		if len(compiledStr) == 0 && string(xStr[i]) == "0" {
			i--
			continue
		}
		compiledStr += string(xStr[i])

		i--
	}

	if compiledStr == "" {
		return 0
	}

	converted, _ := strconv.Atoi(compiledStr)

	if converted > 2147483647 || converted < -2147483647 {
		return 0
	}

	if neg {
		return converted * -1
	}

	return converted
}
