<?php

// 5
echo largestSumAfterKNegations([4,2,3], 1) . PHP_EOL;
// 6
echo largestSumAfterKNegations([3,-1,0,2], 3) . PHP_EOL;
// 13
echo largestSumAfterKNegations([2,-3,-1,5,-4], 2) . PHP_EOL;

/**
 * @param Integer[] $A
 * @param Integer $K
 * @return Integer
 */
function largestSumAfterKNegations($A, $K) {
    sort($A);

    $sum = 0;
    $min = null;
    for ($i = 0; $i < count($A); $i++) {
        if ($K > 0 && $A[$i] < 0) {
            $K--;
            $A[$i] = 0 - $A[$i];
        }
        $sum += $A[$i];

        $min = $min === null || $A[$i] < $min ? $A[$i] : $min;
    }

    if ($K % 2 === 1) {
        $sum -= $min * 2;
    }

    return $sum;
}
