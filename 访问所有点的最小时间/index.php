<?php

/**
 * @param Integer[][] $points
 * @return Integer
 */
function minTimeToVisitAllPoints($points) {
    $num = 0;

    for ($i = 0; $i < count($points) - 1; $i++) {
        $point = $points[$i];
        $point2 = $points[$i + 1];

        $num += max(abs($point[0] - $point2[0]), abs($point[1] - $point2[1]));
    }

    return $num;
}

echo minTimeToVisitAllPoints([
    [1, 1],
    [3, 4],
    [-1, 0]
]);
