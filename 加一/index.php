<?php

/**
 * @param Integer[] $digits
 * @return Integer[]
 */
function plusOne($digits) {
    for ($i = count($digits) - 1; $i >= 0; $i--) {
        $subNum = $digits[$i];
        $subNum++;
        if ($subNum >= 10) {
            $digits[$i] = 0;
            if ($i === 0) {
                array_unshift($digits, 1);
            }
        } else {
            $digits[$i] = $subNum;
            break;
        }
    }
    return $digits;
}

var_dump(plusOne([9]));
