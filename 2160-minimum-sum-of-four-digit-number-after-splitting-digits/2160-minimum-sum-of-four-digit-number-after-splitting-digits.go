func minimumSum(num int) int {
	arr := [4]int{}
	arr[0] = num % 10
	arr[1] = (num / 10) % 10
	arr[2] = (num / 100) % 10
	arr[3] = (num / 1000) % 10
	sort.Ints(arr[:])
	return (arr[0]*10 + arr[2]) + (arr[1]*10 + arr[3])
}
