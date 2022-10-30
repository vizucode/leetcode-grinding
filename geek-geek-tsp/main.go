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

func calculateDistance(cities [][]int, tour []int) int {
	distance := 0
	k := 0
	for _, v := range tour {
		distance += cities[k][v]
		k = v
	}
	distance += cities[k][0]
	return distance
}

func tsp(cities [][]int) int {
	// populate tour
	newTour := []int{}
	for i := 0; i < len(cities); i++ {
		newTour = append(newTour, i)
	}

	// create a permutation tour
	tourChance := [][]int{}
	generateIntPermutations(newTour, len(newTour), &tourChance)

	// get the best minimum path of tour
	minDistance := math.MaxInt

	for _, tour := range tourChance {
		min := calculateDistance(cities, tour)
		if minDistance > min {
			minDistance = min
		}
	}

	return minDistance
}

func main() {
	cities := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	fmt.Println(tsp(cities))
}
