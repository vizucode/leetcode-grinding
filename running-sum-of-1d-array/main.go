package main

import "fmt"

/*
	Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
	Return the running sum of nums.
*/

/*
	Input: nums = [1,2,3,4]
	Output: [1,3,6,10]
	Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/

func runningSum(nums []int) []int {
	count := 0
	for index, num := range nums {
		nums[index] = num + count
		count += num
	}
	return nums
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}
