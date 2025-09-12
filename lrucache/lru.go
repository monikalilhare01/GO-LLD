package main

import "fmt"

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

// Constructor
func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{}, // dummy head
		tail:     &Node{}, // dummy tail
	}
	// link head and tail
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get value by key
func (lru *LRUCache) Get(key int) (int, bool) {
	if node, ok := lru.cache[key]; ok {
		// move accessed node to front
		lru.moveToFront(node)
		return node.value, true
	}
	return 0, false
}

// Set key -> value
func (lru *LRUCache) Set(key, val int) {
	if node, ok := lru.cache[key]; ok {
		// update existing value
		node.value = val
		lru.moveToFront(node)
	} else {
		// evict if full
		if len(lru.cache) >= lru.capacity {
			lru.removeFromLRU(lru.tail.prev) // remove least recently used
		}
		// insert new node
		newNode := &Node{key: key, value: val}
		lru.cache[key] = newNode
		lru.addToFront(newNode)
	}
}

// Move existing node to front (most recently used)
func (lru *LRUCache) moveToFront(node *Node) {
	lru.removeNode(node)
	lru.addToFront(node)
}

// Insert node right after head
func (lru *LRUCache) addToFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

// Remove node from linked list
func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// Remove least recently used node
func (lru *LRUCache) removeFromLRU(node *Node) {
	lru.removeNode(node)
	delete(lru.cache, node.key)
}

// Print current cache content (debug helper)
func (lru *LRUCache) GetCacheValues() {
	fmt.Print("Cache state: ")
	for node := lru.head.next; node != lru.tail; node = node.next {
		fmt.Printf("[%d:%d] ", node.key, node.value)
	}
	fmt.Println()
}

func main() {
	cache := NewLRUCache(2)

	cache.Set(1, 10)
	cache.Set(2, 20)
	cache.GetCacheValues() // [2:20] [1:10]

	cache.Set(3, 30)       // evicts key 1
	cache.GetCacheValues() // [3:30] [2:20]

	if _, found := cache.Get(1); found {
		fmt.Println("ERROR: 1 should have been evicted")
	}

	cache.Get(2) // access 2 → makes it most recent
	cache.GetCacheValues()
	cache.Set(4, 40)       // evicts key 3
	cache.GetCacheValues() // [4:40] [2:20]
}
