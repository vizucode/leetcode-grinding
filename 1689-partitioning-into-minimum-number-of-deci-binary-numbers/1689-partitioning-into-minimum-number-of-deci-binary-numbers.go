func minPartitions(n string) int {
	max := 0

	for _, element := range n {
		num := int(element)
		if max <= num {
			max = num
		}
	}

	num, _ := strconv.Atoi(string(max))
	return num
}
