package main

import (
	"fmt"
)

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph1(node *Node) *Node {
	var visited = make(map[int]bool)
	var valToNode = make(map[int]*Node)
	if node == nil {
		return nil
	}
	res := &Node{Val: node.Val}
	queue := make([]*Node, 0)
	queue = append(queue, node)
	list := make([]*Node, 0)
	list = append(list, res)
	for len(queue) > 0 {
		node = queue[0]
		lNode := list[0]
		fmt.Println(node.Val)
		visited[node.Val] = true
		lNode.Val = node.Val
		valToNode[node.Val] = lNode
		c := node.Neighbors
		for i := 0; i < len(c); i++ {
			if _, has := visited[c[i].Val]; !has {
				t := &Node{Val: c[i].Val}
				valToNode[c[i].Val] = t
				//t.Neighbors = append(t.Neighbors, lNode)
				lNode.Neighbors = append(lNode.Neighbors, t)
				list = append(list, t)
				queue = append(queue, c[i])
				visited[c[i].Val] = true
			} else {
				lNode.Neighbors = append(lNode.Neighbors, valToNode[c[i].Val])
			}
		}
		queue = queue[1:]
		list = list[1:]
	}
	return res
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := make(map[*Node]*Node)
	queue := make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		c := n.Neighbors
		for i := 0; i < len(c); i++ {
			if _, ok := visited[c[i]]; !ok {
				visited[c[i]] = &Node{c[i].Val, []*Node{}}
				queue = append(queue, c[i])
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[c[i]])
		}
	}
	return visited[node]
}
