// Node 双端列表节点
type Node struct {
	prev, next *Node
	key, value int
	cnt        int // 访问频率
}

// remove移除掉当前node
func (node *Node) remove() {
	node.prev.next = node.next
    	node.next.prev = node.prev

    node.cnt = 0 
}

// addNodeBefore 在node节点前面添加新的节点newNode
func (node *Node) addNodeBefore(newNode *Node) {
	newNode.next = node
	newNode.prev = node.prev
	node.prev.next = newNode
	node.prev = newNode
}

// LFUCache 最近最少使用结构体的定义
type LFUCache struct {
	head, tail *Node
	m          map[int]*Node
	capacity   int
}

func (this *LFUCache) setHeader(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

// Constructor : 构造一个LFUCache
func Constructor(capacity int) LFUCache {
	cache := LFUCache{
		head:     new(Node),
		tail:     new(Node),
		m:        make(map[int]*Node),
		capacity: capacity,
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// Get : 从Cache中获取指定的key
func (this *LFUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
        res := node.value
		node.cnt++
		// 遍历，找到合适的节点位置 -- O(N)
		cur := node
		for cur.prev != this.head && cur.prev.cnt <= node.cnt {
			cur = cur.prev
		}
		node.remove()
		cur.addNodeBefore(node)
		return res
	}
	return -1
}

// Put : 从Cache中添加指定的key
func (this *LFUCache) Put(key int, value int) {
	cnt := 1
	if node, ok := this.m[key]; ok {
		delete(this.m, key)
		cnt = node.cnt
		node.remove()
	} else {
		if len(this.m) == this.capacity {
			node := this.tail.prev
			k := node.key
			delete(this.m, k)
			node.remove()
		}
	}

	cur := this.tail
	for cur.prev != this.head && cur.prev.cnt <= cnt {
		cur = cur.prev
	}
	node := &Node{nil, nil, key, value, cnt}
	cur.addNodeBefore(node)
	this.m[key] = node
}