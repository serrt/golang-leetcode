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
