package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//可以生成随机数，但是数值不会变。
	a := rand.Int()
	b := rand.Intn(100)
	fmt.Println(a, b)

	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	//生成指定范围内的随机数
	//生成[15，88]之间的随机数,括号左包含右不包含
	n := rand.Intn(73) + 15 //(88-15 )+15
	fmt.Println(n)

	nums := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11}
	knuthShuffle1(nums)
	fmt.Println(nums)

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(aa), func(i, j int) { aa[i], aa[j] = aa[j], aa[i] })
	fmt.Println(a)
}

func knuthShuffle1(nums []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(nums); i++ {
		r := rand.Intn(i + 1)
		nums[i], nums[r] = nums[r], nums[i]
	}
}
