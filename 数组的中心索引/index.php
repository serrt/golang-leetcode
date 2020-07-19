<?php

/**
 * 获取数组的中心索引
 * 
 * 数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和, 否则返回 -1
 * 
 * @param Integer[] $nums
 * @return Integer
 */
function pivotIndex($nums) {
    $index = -1;
    $len = count($nums);
    if ($len > 10000) {
        return $index;
    }
    $sum = array_sum($nums);

    foreach ($nums as $key => $value) {
        $num_left = array_sum(array_slice($nums, 0, $key));
        
        // 右边部分的和 = 总和 - 左边部分 - 当前值
        $num_right = $sum - $num_left - $value;
        if ($num_left == $num_right) {
            $index = $key;
            break;
        }
    }
    return $index;
}

echo "[1, 7, 3, 6, 5, 6] 中心索引: " . pivotIndex([1, 7, 3, 6, 5, 6]);
echo "[1, 2, 3] 中心索引: " . pivotIndex([1, 2, 3]);
