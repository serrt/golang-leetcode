<?php

/**
 * @param Integer[][] $grid
 * @param Integer $k
 * @return Integer[][]
 */
function shiftGrid($grid, $k) {
    $arr = array_values($grid);
    var_dump($arr);
}

shiftGrid([[1,2,3],[4,5,6],[7,8,9]], 1);
