func numJewelsInStones(jewels string, stones string) int {
	hashmap := make(map[string]string)
	counter := 0
	for _, element := range jewels {
		hashmap[string(element)] = string(element)
	}

	for _, element := range stones {
		_, isExist := hashmap[string(element)]
		if isExist {
			counter++
		}
	}
	return counter
}