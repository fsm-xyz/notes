package main

import (
	"fmt"
)

func main() {
	arr := []float64{20, 7, 2, 8, 8, 9, 2, 100, 2}
	fmt.Println(maxProfit(arr))
}

// 给定一个股票价格数组，只交易一次，求最大利润
func maxProfit(arr []float64) (float64, int, int) {
	if len(arr) < 2 {
		return 0, 0, 0
	}

	buy := 0
	sell := 0
	maxProfit := 0.0
	tempBuyIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[tempBuyIndex] {
			tempBuyIndex = i
		} else {
			// 不需要下表的时候只需要这句就可以
			// maxProfit = math.Max(maxProfit, arr[i]-arr[tempBuyIndex])
			profit := arr[i] - arr[tempBuyIndex]
			if profit > maxProfit {
				maxProfit = profit
				buy = tempBuyIndex
				sell = i
			}
		}
	}
	return maxProfit, buy, sell
}
