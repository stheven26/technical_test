package main

import "fmt"

var (
	beli, jual int
)

func MaxProfit(prices []int, i int) int {
	var max int
	if i < 2 {
		return -1
	} else {
		for i := range prices {
			for j := i; j < len(prices); j++ {
				if prices[j]-prices[i] > max {
					max = prices[j] - prices[i]
					beli = prices[i]
					jual = prices[j]
				}
			}
		}
	}

	return max
}

func main() {
	slice := []int{4, 11, 2, 20, 59, 80}
	fmt.Printf("beli: %v, jual: %v, max profit: %v\n", beli, jual, MaxProfit(slice, 2))
}
