package day1126

type LRUCache struct {
	capacity   int
	head, tail *DNode
	cashMap    map[int]*DNode
}

type DNode struct {
	prev       *DNode
	next       *DNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
		cashMap:  map[int]*DNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cashMap[key]
	if !ok {
		return -1
	}
	l.removeToHead(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	//存在
	if node, ok := l.cashMap[key]; ok {
		node.value = value
		l.removeToHead(node)
		return
	}
	//不存在
	if len(l.cashMap) >= l.capacity {
		removed := l.removeTail()
		delete(l.cashMap, removed.key)
	}
	node := initNode(key, value)
	l.cashMap[key] = node
	l.addNodeToHead(node)
}

func initNode(key, value int) *DNode {
	return &DNode{
		key:   key,
		value: value,
	}
}

func (l LRUCache) addNodeToHead(node *DNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l LRUCache) removeTail() *DNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func (l LRUCache) removeToHead(node *DNode) {
	l.removeNode(node)
	l.addNodeToHead(node)
}

func (l LRUCache) removeNode(node *DNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
