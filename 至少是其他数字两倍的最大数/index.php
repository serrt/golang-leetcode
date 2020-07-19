<?php

echo dominantIndex([0,0,0,0]);

/**
 * 在一个给定的数组nums中，总是存在一个最大元素
 * 查找数组中的最大元素是否至少是数组中每个其他数字的两倍
 * 如果是，则返回最大元素的索引，否则返回-1
 * 
 * @param Integer[] $nums
 * @return Integer
 */
function dominantIndex($nums) {
    list($maxValue, $maxIndex) = arrayMax($nums);
    $isMax = true;
    foreach($nums as $value) {
        if ($maxValue != $value && $maxValue < $value * 2) {
            $isMax = false;
        }
    }
    return $isMax ? $maxIndex : -1;
}

function arrayMax($nums) {
    $maxIndex = -1;
    $maxValue = 0;
    foreach($nums as $key => $value) {
        if ($value > $maxValue) {
            $maxValue = $value;
            $maxIndex = $key;
        }
    }
    return [$maxValue, $maxIndex];
}
