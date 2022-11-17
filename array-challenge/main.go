package main

import (
	"fmt"
)

func ArrayChallenge(arr []int) int {

	max := arr[0]
	sum := 0
	// take the largest number
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	// check if sum == max
	for _, value := range arr {
		if value != max {
			sum += value
		}
	}

	if sum >= max {
		return 1
	}

	return 0
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(ArrayChallenge([]int{3,5,-1,8,12}))

}
