package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0,0,0,1}
	fmt.Println("dominantIndex", dominantIndex(nums1))
}

/**
 * 在一个给定的数组nums中，总是存在一个最大元素
 * 查找数组中的最大元素是否至少是数组中每个其他数字的两倍
 * 如果是，则返回最大元素的索引，否则返回-1
 */
func dominantIndex(nums []int) int {
	numIndex := -1
	maxValue, maxIndex := arrMax(nums)
	maxBol := true
	for _, value := range nums {
		if (maxValue != value && maxValue < value * 2) {
			maxBol = false
		}
	}

	if (maxBol) {
		numIndex = maxIndex
	}

	return numIndex
}

/**
 * 返回数组的最大值, 最大值的索引
 */
func arrMax(nums []int) (int, int) {
	var maxValue int
	maxIndex := -1
	for index, value := range nums {
		if (value > maxValue) {
			maxValue = value
			maxIndex = index
		}
	}
	return maxValue, maxIndex
}
