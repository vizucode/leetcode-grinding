func smallestEvenMultiple(n int) int {
	//     2,5,10
	arr := []int{}
	for i := 1; i <= n*2; i++ {
		if i%2 == 0 && i%n == 0 {
			arr = append(arr, i)
		}
	}
	return arr[0]
}