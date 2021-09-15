package main

func main() {

}

func maxProfit(prices []int) int {
	res := 0
	curMin := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < curMin {
			curMin = prices[i]
		}
		res = max(res, prices[i]-curMin)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
