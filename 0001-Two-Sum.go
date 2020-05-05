package main

import (
	"fmt"
	"log"
)

/**
 * 给你两个二进制串，求其和的二进制串
 */

func main() {
	a, b := "11", "1"

	fmt.Println("Example 1: ", a, b, twoSum(a, b))

	a2, b2 := "1010", "1011"

	fmt.Println("Example 1: ", a2, b2, twoSum(a2, b2))
}

func twoSum(str1 string, str2 string) string {
	num1, num2 := bin2Dec(str1), bin2Dec(str2)

	numSum := num1 + num2

	return dec2Bin(numSum)
}

/**
 * 二进制 - 十进制
 */
func bin2Dec(str string) (num int) {
	l := len(str)
	for i := l - 1; i >= 0; i-- {
		num += (int(str[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

/**
 * 十进制 - 二进制
 */
func dec2Bin(n int) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}
