package day1106

type StockSpanner struct {
	nums  []int
	stack []int
}

func Constructor() StockSpanner {
	spanner := StockSpanner{stack: []int{}, nums: []int{}}
	return spanner
}

func (s *StockSpanner) Next(price int) (res int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, len(s.nums))
		s.nums = append(s.nums, price)
		return 1
	}
	for len(s.stack) > 0 && price >= s.nums[s.stack[len(s.stack)-1]] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	if len(s.stack) > 0 {
		res = len(s.nums) - s.stack[len(s.stack)-1]
	} else {
		res = len(s.nums) + 1
	}
	s.stack = append(s.stack, len(s.nums))
	s.nums = append(s.nums, price)
	return
}
