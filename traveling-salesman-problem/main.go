package main

import (
	"fmt"
	"math"
)

func generateIntPermutations(array []int, n int, result *[][]int) {
	if n == 1 {
		dst := make([]int, len(array))
		copy(dst, array[:])
		*result = append(*result, dst)
	} else {
		for i := 0; i < n; i++ {
			generateIntPermutations(array, n-1, result)
			if n%2 == 0 {
				// Golang allow us to do multiple assignments
				array[0], array[n-1] = array[n-1], array[0]
			} else {
				array[i], array[n-1] = array[n-1], array[i]
			}
		}
	}
}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
}

func calculateDistance(cities [][]int, tour []int) int {
	totalDistance := 0
	for i := 1; i < len(cities); i++ {
		x1 := cities[tour[i-1]][0]
		y1 := cities[tour[i-1]][1]
		x2 := cities[tour[i]][0]
		y2 := cities[tour[i]][1]
		d := distance(float64(x1), float64(y1), float64(x2), float64(y2))
		totalDistance += int(d)
	}
	x1 := cities[len(tour)-1][0]
	y1 := cities[len(tour)-1][1]
	x2 := cities[0][0]
	y2 := cities[0][0]
	d := distance(float64(x1), float64(y1), float64(x2), float64(y2))
	totalDistance += int(d)

	return totalDistance
}

func tsp(cities [][]int) int {
	// populate tour
	populateTour := []int{}
	for i := 0; i < len(cities); i++ {
		populateTour = append(populateTour, i)
	}

	// populate permutation
	tourCombination := [][]int{}
	generateIntPermutations(populateTour, len(populateTour), &tourCombination)

	// get minimal tour distance
	minimalTourDistance := math.MaxInt
	for _, tour := range tourCombination {
		min := calculateDistance(cities, tour)
		if minimalTourDistance > min {
			minimalTourDistance = min
		}
	}

	return minimalTourDistance
}

func main() {
	cities := [][]int{
		{3, 3},
		{7, 4},
		{9, 2},
		{3, 9},
		{9, 8},
	}
	fmt.Println(tsp(cities))
}
