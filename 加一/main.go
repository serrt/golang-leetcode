package main

import (
	"fmt"
)

func main() {
	num := []int{8,9,9,9}
	fmt.Println("", plusOne(num))
}

func plusOne(digits []int) []int {
	reverseArray := reverse(digits)
	for i := 0; i < len(digits); i++ {
		num := reverseArray[i] + 1
		if (num >= 10) {
			num = 0
			reverseArray[i] = num
			if (i == len(digits) - 1) {
                reverseArray = append(reverseArray, 1);
            }
		} else {
			reverseArray[i] = num
			break;
		}
	}
	return reverse(reverseArray)
}

func reverse(digits []int) []int {
	var reverseArray []int
	for i := len(digits) - 1; i >= 0; i-- {
		reverseArray = append(reverseArray, digits[i])
	}
	return reverseArray
}
