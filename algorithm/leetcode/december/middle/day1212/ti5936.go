package day1212

var Bombs [][]int

func maximumDetonation(bombs [][]int) int {
	Bombs = bombs
	n := len(bombs)
	max := 0
	for i := 0; i < n; i++ {
		queue := make([]int, 0)
		queue = append(queue, i)
		flag := make([]bool, n)
		flag[i] = true
		for len(queue) > 0 {
			index := queue[0]
			for j := 0; j < n; j++ {
				if !flag[j] && effected(index, j) {
					queue = append(queue, j)
					flag[j] = true
				}
			}
			queue = queue[1:]
		}
		t := 0
		for _, value := range flag {
			if value {
				t++
			}
		}
		if t > max {
			max = t
		}
	}
	return max
}

func effected(index, j int) bool {
	return (Bombs[j][0]-Bombs[index][0])*(Bombs[j][0]-Bombs[index][0])+(Bombs[j][1]-Bombs[index][1])*(Bombs[j][1]-Bombs[index][1]) <= Bombs[index][2]*Bombs[index][2]
}
