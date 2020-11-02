package main;

import (
	"fmt"
	"math"
)

func main() {
	points := [][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}

	fmt.Println(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	num := 0
	for i := 0; i < len(points) - 1; i++ {
		point := points[i]
		point2 := points[i + 1]

		deviation_x := math.Abs((float64)(point[0] - point2[0]))
		deviation_y := math.Abs((float64)(point[1] - point2[1]))
		num += int(math.Max(deviation_x, deviation_y))
	}
	return num
}
