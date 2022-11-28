package leetcode

import (
	"fmt"
	"strconv"
)

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.



Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"


Constraints:

1 <= num1.length, num2.length <= 200
num1 and num2 consist of digits only.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
*/

func multiply(num1 string, num2 string) string {
	// Check if the parameters is 0, then return immediately. Because anything multiplied by zero, is zero
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// Define array for containing all multiplied operation results
	arr := [][]int{}

	// X for additional zero
	// Example: num1 -> 456, num -> 123.
	// Then:
	//  1368
	//  9120
	// 45600
	// Every new enter give new 0
	x := 0
	for i := len(num1) - 1; i >= 0; i-- {
		multiplier1, _ := strconv.Atoi(string(num1[i]))
		carry := 0
		arr2 := []int{}

		for k := x; k > 0; k-- {
			arr2 = append(arr2, 0)
		}

		for j := len(num2) - 1; j >= 0; j-- {
			multiplier2, _ := strconv.Atoi(string(num2[j]))
			multiplied := (multiplier1 * multiplier2) + carry
			carry = multiplied / 10
			arr2 = append(arr2, multiplied%10)
		}
		if carry != 0 {
			arr2 = append(arr2, carry)
		}

		arr = append(arr, arr2)
		x++
	}

	arrFinal := []string{}
	max := 0
	for _, arr2 := range arr {
		str := ""
		for _, x := range arr2 {
			str += strconv.Itoa(x)
		}
		if len(str) > max {
			max = len(str)
		}
		arrFinal = append(arrFinal, str)
	}

	carry := 0
	superFinalString := ""
	for i := 0; i <= max; i++ {
		tot := 0
		for _, final := range arrFinal {
			if i < len(final) {
				x, _ := strconv.Atoi(string(final[i]))
				tot += x
			}
		}
		tot += carry
		if tot == 0 && i == max {
			continue
		}
		carry = tot / 10
		tot = tot % 10
		superFinalString = fmt.Sprintf("%d%s", tot, superFinalString)
	}

	return superFinalString
}

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			res[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
			res[i+j] += res[i+j+1] / 10
			res[i+j+1] %= 10
		}
	}
	i := 0
	ans := ""
	for res[i] == 0 {
		i++
	}
	for i < len(res) {
		ans += strconv.Itoa(res[i])
		i++
	}
	return ans
}
