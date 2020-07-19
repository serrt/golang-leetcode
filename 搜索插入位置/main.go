package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	index := searchInsert(nums, target)
	fmt.Println(index)
}

/**
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 */
func searchInsert(nums []int, target int) int {
	targetIndex := -1
	for index, value := range nums {
		if (value > target) {
			targetIndex = index
			break
		}
		if (value < target) {
			targetIndex = index + 1
		}
		if (value == target) {
			targetIndex = index
			break
		}
	}
	return targetIndex
}
