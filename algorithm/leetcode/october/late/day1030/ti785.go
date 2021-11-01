package main

import "fmt"

func main() {
	res := isBipartite([][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	})
	fmt.Println(res)
}

func isBipartite(graph [][]int) bool {
	length := len(graph)
	visited := make([]bool, length)
	flags := make([]int, length)

	//for i := 0; i < length; i++ {
	//	if len(graph[i]) != 0 {
	//		flags[i] = 1
	//		queue = append(queue, i)
	//		break
	//	}
	//}
	for k := 0; k < length; k++ {
		var queue []int
		if flags[k] == 0 {
			counter := 1
			queue = append(queue, k)
			for len(queue) > 0 {
				t := 0
				if counter%2 == 0 {
					t = 1
				} else {
					t = 2
				}
				counter++
				size := len(queue)
				for i := 0; i < size; i++ {
					node := queue[0]
					visited[node] = true
					for j := 0; j < len(graph[node]); j++ {
						if !visited[graph[node][j]] {
							queue = append(queue, graph[node][j])
							index := graph[node][j]
							if flags[index] == 0 || flags[index] == t {
								flags[index] = t
							} else {
								return false
							}
						}
					}
					queue = queue[1:]
				}
			}
		}
	}

	return true
}
