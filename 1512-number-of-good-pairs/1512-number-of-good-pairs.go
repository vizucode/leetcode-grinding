func numIdenticalPairs(nums []int) int {
	hashmap := make(map[int]int)
	count := 0

	for _, element := range nums {
		count += hashmap[element]
		hashmap[element]++
	}
	return count
}