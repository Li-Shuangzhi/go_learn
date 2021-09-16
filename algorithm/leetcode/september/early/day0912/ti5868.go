package main

func main() {

}

type Cp struct {
	x int
	y int
}

func interchangeableRectangles(rectangles [][]int) (res int64) {
	mp := make(map[Cp]int64)
	for i := 0; i < len(rectangles); i++ {
		x, y := getXY(rectangles[i][0], rectangles[i][1])
		cp := Cp{
			x: x,
			y: y,
		}
		if _, ok := mp[cp]; ok {
			res += mp[cp]
		}
		mp[cp]++
	}
	return
}

func getXY(x int, y int) (int, int) {
	t := gcd(x, y)
	return x / t, y / t
}
func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
