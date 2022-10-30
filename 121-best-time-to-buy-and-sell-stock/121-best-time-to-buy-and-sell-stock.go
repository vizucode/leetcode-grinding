func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		// today price - minimum price
		tempProfit := price - minPrice
		if tempProfit > profit {
			profit = tempProfit
		}
	}

	return profit
}