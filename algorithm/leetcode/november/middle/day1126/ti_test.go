package day1126

import (
	"container/list"
	"testing"
)

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestTi(t *testing.T) {

	cash := Constructor(2)
	cash.Put(1, 1)
	cash.Put(2, 2)
	r := cash.Get(1)
	t.Log(r)
	cash.Put(3, 3)
	r = cash.Get(2)
	t.Log(r)
}

func TestContainer(t *testing.T) {
	l := list.New()
	//把4元素放在最后

	l.PushBack(&DNode{value: 10})
	front := l.Front()
	node := front.Value.(*DNode)
	t.Log(node.value)

}
