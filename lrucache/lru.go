package main

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = nil
	lru.tail.prev = nil
	return lru
}

func (lru *LRUCache) Get(key int) (int, bool) {
	if node, ok := lru.cache[key]; ok {
		return node.value, true
	}
	return 0, false
}

func (lru *LRUCache) Set(key, val int) {
	if node, ok := lru.cache[key]; ok {
		node.value = val
		lru.moveToFront(node)
	} else {
		if len(lru.cache) >= lru.capacity {
			lru.removeFromLRU(lru.tail)
		}
		newNode := &Node{
			key:   key,
			value: val,
		}
		lru.cache[key] = newNode
		lru.moveToFront(newNode)
	}
}

func (lru *LRUCache) moveToFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node

}
func (lru *LRUCache) removeFromLRU(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func main() {

}
