package main

import "fmt"

func findDuplicateNumber(numbers []int) {
	tempMap := map[int]int{}
	dupMap := map[int]int{}
	result := []int{}

	for _, number := range numbers {
		_, isExist := tempMap[number]

		if isExist {
			dupMap[number] = number
		}

		tempMap[number] = number
	}

	for _, number := range dupMap {
		result = append(result, number)
	}

	fmt.Println(result)
}

func main() {
	findDuplicateNumber([]int{1, 1})                            // [1]
	findDuplicateNumber([]int{1, 2, 3, 4})                      // []
	findDuplicateNumber([]int{1, 2, 1, 2, 2, 3, 4, 5, 5, 5, 5}) // [1,2,5]
}
