package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Key       string
	Value     string
	ExpiresAt time.Time
}

type LRUCache struct {
	capacity int
	cache    map[string]*CacheItem
	order    *DoublyLinkedList
	mutex    sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*CacheItem),
		order:    NewDoublyLinkedList(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, exists := c.cache[key]
	if !exists || time.Now().After(item.ExpiresAt) {
		return "", false
	}
	c.order.MoveToFront(item)
	return item.Value, true
}

func (c *LRUCache) Set(key string, value string, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, exists := c.cache[key]; exists {
		item.Value = value
		item.ExpiresAt = time.Now().Add(ttl)
		c.order.MoveToFront(item)
		return
	}

	if len(c.cache) >= c.capacity {
		oldest := c.order.RemoveOldest()
		delete(c.cache, oldest.Key)
	}

	newItem := &CacheItem{
		Key:       key,
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
	c.cache[key] = newItem
	c.order.AddToFront(newItem)
}

func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, exists := c.cache[key]; exists {
		c.order.Remove(item)
		delete(c.cache, key)
	}
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	item *CacheItem
	next *Node
	prev *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) AddToFront(item *CacheItem) {
	node := &Node{item: item}
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
}

func (dll *DoublyLinkedList) MoveToFront(item *CacheItem) {
	node := dll.findNode(item)
	if node == dll.head {
		return
	}
	dll.removeNode(node)
	dll.AddToFront(item)
}

func (dll *DoublyLinkedList) Remove(item *CacheItem) {
	node := dll.findNode(item)
	dll.removeNode(node)
}

func (dll *DoublyLinkedList) RemoveOldest() *CacheItem {
	if dll.tail == nil {
		return nil
	}
	oldest := dll.tail.item
	dll.removeNode(dll.tail)
	return oldest
}

func (dll *DoublyLinkedList) findNode(item *CacheItem) *Node {
	for node := dll.head; node != nil; node = node.next {
		if node.item == item {
			return node
		}
	}
	return nil
}

func (dll *DoublyLinkedList) removeNode(node *Node) {
	if node == dll.head {
		dll.head = node.next
		if dll.head != nil {
			dll.head.prev = nil
		}
	} else if node == dll.tail {
		dll.tail = node.prev
		if dll.tail != nil {
			dll.tail.next = nil
		}
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}
