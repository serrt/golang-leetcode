package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("[1, 2, 3] 中心索引: ", pivotIndex(nums))
	nums2 := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("[1, 7, 3, 6, 5, 6] 中心索引: ", pivotIndex(nums2))
}

/**
 * 获取数组的中心索引
 *
 * 数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和, 否则返回 -1
 *
 */
func pivotIndex(nums []int) int {
	index := -1
	numsLen := len(nums)
	if (numsLen > 10000) {
		return index
	}
	numSum := sum(nums)
	for key, value := range nums {
		numsLeft := sum(nums[0:key])

		// 右边部分的和 = 总和 - 左边部分 - 当前值
		numsRight := numSum - numsLeft - value
		if (numsLeft == numsRight) {
			index = key
			break
		}
    }
	return index
}

func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
