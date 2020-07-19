
console.log("[1, 2, 3] 中心索引: ", pivotIndex([1, 2, 3]))
console.log("[1, 7, 3, 6, 5, 6] 中心索引: ", pivotIndex([1, 7, 3, 6, 5, 6]))

/**
 * 获取数组的中心索引
 *
 * 数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和, 否则返回 -1
 *
 */
function pivotIndex(arr) {
    let index = -1
    const len = arr.length

    if (len > 10000) {
		return index
    }
    const arrSum = sum(arr)

    for (let i = 0; i < len; i++) {
        const arrLeft = sum(arr.slice(0, i))

        // 右边部分的和 = 总和 - 左边部分 - 当前值
        const arrRight  = arrSum - arrLeft - arr[i]
        if (arrLeft === arrRight) {
            index = i;
            break;
        }
    }

    return index
}

function sum(arr) {
    let sum = 0
    for (let i = 0; i < arr.length; i++) {
        sum += arr[i]
    }
    return sum
}
