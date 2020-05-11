/**
 * 在一个给定的数组nums中，总是存在一个最大元素
 * 查找数组中的最大元素是否至少是数组中每个其他数字的两倍
 * 如果是，则返回最大元素的索引，否则返回-1
 * 
 * @param {number[]} nums
 * @return {number}
 */
var dominantIndex = function(nums) {
    var max = arrayMax(nums)
    var maxValue = max[0]
    var maxIndex = max[1]
    var maxBol = true
    for (let i = 0; i < nums.length; i++) {
        if (maxValue != nums[i] && maxValue < nums[i] * 2) {
            maxBol = false
        }
    }
    return maxBol ? maxIndex : -1
}

var arrayMax = function (nums) {
    var maxIndex = -1
    var maxValue = 0
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] > maxValue) {
            maxValue = nums[i]
            maxIndex = i
        }
    }

    return [maxValue, maxIndex]
}
