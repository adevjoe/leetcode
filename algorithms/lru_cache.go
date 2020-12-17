// https://leetcode.com/problems/lru-cache/

package algorithms

type LRUCache struct {
	values   map[int]*lruDoubleListNode
	head     *lruDoubleListNode
	tail     *lruDoubleListNode
	capacity int
}

type lruDoubleListNode struct {
	key   int
	value int
	pre   *lruDoubleListNode
	next  *lruDoubleListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		values:   map[int]*lruDoubleListNode{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.values[key]
	if !ok {
		return -1
	}
	this.moveAhead(v)
	return v.value
}

func (this *LRUCache) moveAhead(node *lruDoubleListNode) {
	if this.head == this.tail {
		this.head = node
		this.tail = node
		return
	}
	if node == this.head {
		return
	}
	if node == this.tail {
		this.tail = node.pre
	}
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}
	node.pre = nil
	node.next = this.head
	this.head.pre = node
	this.head = node
}

func (this *LRUCache) putHead(node *lruDoubleListNode) {
	if this.head == nil {
		this.head = node
		if this.tail == nil {
			this.tail = node
		}
		return
	}
	this.head.pre = node
	node.next = this.head
	this.head = node
}

func (this *LRUCache) deleteTail() {
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
		return
	}
	this.tail = this.tail.pre
	this.tail.next = nil
}

func (this *LRUCache) Put(key int, value int) {
	// update value
	if v, ok := this.values[key]; ok {
		this.values[key].value = value
		this.moveAhead(v)
		return
	}

	if len(this.values) == this.capacity {
		// evict
		delete(this.values, this.tail.key)
		this.deleteTail()
	}

	// put
	node := &lruDoubleListNode{
		key:   key,
		value: value,
	}
	this.putHead(node)
	this.values[key] = node
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
