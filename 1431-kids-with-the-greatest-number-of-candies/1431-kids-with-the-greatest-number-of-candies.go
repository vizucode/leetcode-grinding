func kidsWithCandies(candies []int, extraCandies int) []bool {
	//     urutkan desc
	//     ambil index pertama -> dapat candy terbesar
	//     loop candies + extra candies
	//     bandingkan candies[i]+extra >= greatestBeforeExtra[0]

	//     5 -> greatestBeforeExtra
	//     5+3 = 8 true

	result := []bool{}
	greatestCandies := []int{}

	greatestCandies = append(greatestCandies, candies...)

	sort.Slice(greatestCandies, func(i int, j int) bool {
		return greatestCandies[i] > greatestCandies[j]
	})

	for _, candy := range candies {
		addExtraCandies := candy + extraCandies
		if addExtraCandies >= greatestCandies[0] {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}