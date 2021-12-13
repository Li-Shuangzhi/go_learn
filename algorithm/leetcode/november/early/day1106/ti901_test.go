package day1106

import "testing"

func TestStockSpanner(t *testing.T) {
	obj := Constructor()
	res := obj.Next(100)
	t.Log(res)
	res = obj.Next(80)
	t.Log(res)
	res = obj.Next(60)
	t.Log(res)
	res = obj.Next(70)
	t.Log(res)
	res = obj.Next(60)
	t.Log(res)
	res = obj.Next(75)
	t.Log(res)
	res = obj.Next(85)
	t.Log(res)
}
