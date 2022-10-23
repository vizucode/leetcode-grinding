package main

import "fmt"

/*
	Given an array of integers nums, return the number of good pairs.

	A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/

/*
	Input: nums = [1,2,3,1,1,3]
	Output: 4
	Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
*/

func numIdenticalPairs(nums []int) int {
	hashmap := make(map[int]int)
	count := 0

	for _, element := range nums {
		count += hashmap[element]
		hashmap[element]++
	}
	return count
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}
