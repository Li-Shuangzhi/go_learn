package day1127

import "container/list"

type LRUCache struct {
	capacity    int
	dLinkedList *list.List
	cash        map[int]*list.Element
}

type DLinkNode struct {
	key, value int
}

func (d DLinkNode) Value() interface{} {
	return d.value
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:    capacity,
		dLinkedList: list.New(),
		cash:        map[int]*list.Element{},
	}
}

func (l *LRUCache) Get(key int) int {
	element, ok := l.cash[key]
	if !ok {
		return -1
	}
	l.dLinkedList.MoveToFront(element)
	return element.Value.(*DLinkNode).value
}

func (l *LRUCache) Put(key int, value int) {
	element, ok := l.cash[key]
	//如果key已存在
	if ok {
		node := element.Value.(*DLinkNode)
		node.value = value
		l.dLinkedList.MoveToFront(element)
	}
	//如果key不存在
	if len(l.cash) >= l.capacity {
		back := l.dLinkedList.Back()
		l.dLinkedList.Remove(back)
		delete(l.cash, back.Value.(*DLinkNode).key)
	}
	node := &DLinkNode{
		key:   key,
		value: value,
	}
	front := l.dLinkedList.PushFront(node)
	l.cash[key] = front
}
