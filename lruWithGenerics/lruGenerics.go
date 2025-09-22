package main

import (
	"fmt"
	"sync"
	"time"
)

// Problem: Thread-Safe LRU Cache with TTL
// Implement a thread-safe LRU (Least Recently Used) cache with Time-To-Live (TTL) expiration that supports concurrent access and automatic cleanup.

// Requirements:
// 1. LRU eviction when capacity is exceeded
// 2. TTL-based expiration for entries
// 3. Thread-safe concurrent access
// 4. Automatic cleanup of expired entries
// 5. Support for generic key-value types

type Cache[K comparable, V any] interface {
	// Get retrieves a value by key, returns value and true if found
	Get(key K) (V, bool)

	// Set stores a key-value pair with TTL
	Set(key K, value V, ttl time.Duration)

	// Delete removes a key from cache
	Delete(key K) bool

	// Size returns current number of entries
	Size() int

	// Clear removes all entries
	Clear()

	// Stats returns cache statistics
	Stats() CacheStats
}

type CacheStats struct {
	Hits        int64
	Misses      int64
	Evictions   int64
	Expirations int64
	Size        int
	Capacity    int
}

type CacheEntry[V any] struct {
	Value     V
	ExpiresAt time.Time
	// TODO: Add fields for LRU tracking
}

type lruCache[K comparable, V any] struct {
	capacity int
	items    map[K]*node[K, V]

	// TODO: Implement doubly linked list for LRU
	// head, tail *node[K, V]

	// TODO: Add mutex for thread safety
	// mu sync.RWMutex

	// TODO: Add statistics tracking
	// stats CacheStats

	// TODO: Add cleanup mechanism
	// cleanupInterval time.Duration
	// stopCleanup     chan struct{}
}

type node[K comparable, V any] struct {
	key       K
	value     V
	expiresAt time.Time
	// TODO: Add prev/next pointers for doubly linked list
	// prev, next *node[K, V]
}

func NewCache[K comparable, V any](capacity int) Cache[K, V] {
	// TODO: Initialize cache with:
	// - map for O(1) key lookup
	// - doubly linked list for O(1) LRU operations
	// - background cleanup goroutine
	return nil
}

func (c *lruCache[K, V]) Get(key K) (V, bool) {
	// TODO: Implement Get with:
	// - Thread safety (read lock)
	// - TTL expiration check
	// - LRU update (move to front)
	// - Statistics update
	var zero V
	return zero, false
}

func (c *lruCache[K, V]) Set(key K, value V, ttl time.Duration) {
	// TODO: Implement Set with:
	// - Thread safety (write lock)
	// - Check if key exists (update vs insert)
	// - LRU eviction if at capacity
	// - Move to front of LRU list
	// - Statistics update
}

func (c *lruCache[K, V]) Delete(key K) bool {
	// TODO: Implement Delete with:
	// - Thread safety (write lock)
	// - Remove from map and linked list
	// - Statistics update
	return false
}

func (c *lruCache[K, V]) Size() int {
	// TODO: Thread-safe size check
	return 0
}

func (c *lruCache[K, V]) Clear() {
	// TODO: Remove all entries thread-safely
}

func (c *lruCache[K, V]) Stats() CacheStats {
	// TODO: Return current statistics
	return CacheStats{}
}

// Helper methods for doubly linked list operations
func (c *lruCache[K, V]) moveToFront(n *node[K, V]) {
	// TODO: Move node to front of LRU list
}

func (c *lruCache[K, V]) removeLRU() {
	// TODO: Remove least recently used item
}

func (c *lruCache[K, V]) removeNode(n *node[K, V]) {
	// TODO: Remove specific node from linked list
}

func (c *lruCache[K, V]) startCleanup() {
	// TODO: Start background cleanup goroutine
	// - Periodically scan for expired entries
	// - Remove expired items
	// - Handle graceful shutdown
}

func (c *lruCache[K, V]) isExpired(n *node[K, V]) bool {
	// TODO: Check if node has expired
	return false
}

func main() {
	// Test 1: Basic LRU functionality
	cache := NewCache[string, int](3)

	cache.Set("a", 1, 5*time.Second)
	cache.Set("b", 2, 5*time.Second)
	cache.Set("c", 3, 5*time.Second)

	// Should evict "a" (least recently used)
	cache.Set("d", 4, 5*time.Second)

	if _, found := cache.Get("a"); found {
		fmt.Println("ERROR: 'a' should have been evicted")
	}

	// Test 2: TTL expiration
	cache.Set("expire", 100, 1*time.Second)

	if val, found := cache.Get("expire"); !found || val != 100 {
		fmt.Println("ERROR: Should find 'expire' immediately")
	}

	time.Sleep(2 * time.Second)

	if _, found := cache.Get("expire"); found {
		fmt.Println("ERROR: 'expire' should have expired")
	}

	// Test 3: LRU order update on access
	cache.Clear()
	cache.Set("x", 1, 10*time.Second)
	cache.Set("y", 2, 10*time.Second)
	cache.Set("z", 3, 10*time.Second)

	// Access "x" to make it most recently used
	cache.Get("x")

	// Add new item, should evict "y" (now LRU)
	cache.Set("w", 4, 10*time.Second)

	if _, found := cache.Get("y"); found {
		fmt.Println("ERROR: 'y' should have been evicted")
	}

	// Test 4: Concurrent access
	testConcurrentAccess(cache)

	// Test 5: Statistics
	stats := cache.Stats()
	fmt.Printf("Cache Stats: %+v\n", stats)
}

func testConcurrentAccess(cache Cache[string, int]) {
	var wg sync.WaitGroup

	// Concurrent writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				cache.Set(key, j, 2*time.Second)
			}
		}(i)
	}

	// Concurrent readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key-%d-%d", id%10, j)
				cache.Get(key)
			}
		}(i)
	}

	wg.Wait()
}
