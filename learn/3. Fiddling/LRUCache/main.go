package main

import (
	"fmt"
)

// Node represents a single key-value pair in the cache
type Node struct {
	key   string
	value interface{}
	prev  *Node
	next  *Node
}

// LRUCache implements a basic LRU cache
type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	cache    map[string]*Node
}

// NewLRUCache creates a new LRUCache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node),
	}
}

// Get retrieves the value for a given key from the cache
func (c *LRUCache) Get(key string) interface{} {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.value
	}
	return nil
}

// Put adds or updates a key-value pair in the cache
func (c *LRUCache) Put(key string, value interface{}) {
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.moveToFront(node)
		return
	}

	// Evict least recently used item if cache is full
	if len(c.cache) == c.capacity {
		c.evictLeastRecentlyUsed()
	}

	// Add new key-value pair to cache
	newNode := &Node{key: key, value: value}
	c.addToFront(newNode)
	c.cache[key] = newNode
}

// moveToFront moves a node to the front of the linked list
func (c *LRUCache) moveToFront(node *Node) {
	if node == c.head {
		return
	}

	if node == c.tail {
		c.tail = node.prev
		c.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.prev = nil
	node.next = c.head
	c.head.prev = node
	c.head = node
}

// evictLeastRecentlyUsed removes the least recently used item from the cache
func (c *LRUCache) evictLeastRecentlyUsed() {
	if c.tail == nil {
		return
	}

	delete(c.cache, c.tail.key)
	c.tail = c.tail.prev
	if c.tail == nil {
		c.head = nil
	} else {
		c.tail.next = nil
	}
}

// addToFront adds a node to the front of the linked list
func (c *LRUCache) addToFront(node *Node) {
	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		node.next = c.head
		c.head.prev = node
		c.head = node
	}
}

func main() {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Get "a" again, promoting it to the front
	cache.Get("a")

	cache.Put("d", 4) // Evicts "b"

	// Iterate through the cache in order of least recently used
	for node := cache.head; node != nil; node = node.next {
		fmt.Printf("Key: %s, Value: %d\n", node.key, node.value)
	}
}
