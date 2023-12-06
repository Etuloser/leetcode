package besttimetobuyandsellstock

func solution1(prices []int) (ans int) {
	// 需要获取从左往右最大利润
	// 只需要知道第i天前的最小值，即prices[0]到prices[i-1]的最小值minPrice
	// 然后遍历p[i:]获取p[i]-minPrice的最大值
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return
}

// min -.
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// max -.
func max(a, b int) int {
	if b < a {
		return a
	}
	return b
}
