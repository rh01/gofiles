type LRUCache struct {
	hashMap  map[int]*List
	fifo     *FifoList
	capacity int
}

type List struct {
	before *List
	next   *List
	value  []int
}
type FifoList struct {
	head   *List
	tail   *List
	length int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.capacity = capacity
	lru.hashMap = make(map[int]*List)
	lru.fifo = &FifoList{}
	return lru
}

func (this *FifoList) PushFront(node *List) {
	if node.before != nil {
		node.before.next = node.next
		if node.next != nil {
			node.next.before = node.before
		} else {
			this.tail = node.before
		}
		node.before = nil
		node.next = this.head
		this.head.before = node
		this.head = node
	}
}

func (this *FifoList) Add(node []int) *List {
	tmp := &List{nil, this.head, node}
	if this.head != nil {
		this.head.before = tmp
	} else {
		this.tail = tmp
	}
	this.head = tmp
	this.length += 1
	return tmp
}

func (this *FifoList) RemoveBack() *List {
	if this.tail != nil {
		tmp := this.tail
		this.tail = this.tail.before
		if this.tail == nil {
			this.head = nil
		} else {
			this.tail.next = nil
		}
		this.length -= 1
		return tmp
	}
	return nil
}

func (this *FifoList) Len() int {
	return this.length
}

func (this *LRUCache) Get(key int) int {
	if this.hashMap[key] != nil {
		this.fifo.PushFront(this.hashMap[key])

		return this.hashMap[key].value[1]
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.hashMap[key] != nil {
		this.hashMap[key].value[1] = value
		this.fifo.PushFront(this.hashMap[key])
	} else {
		//不线程安全
		tmp := this.fifo.Add([]int{key, value})
		this.hashMap[tmp.value[0]] = tmp
		//不信容量是0，容量是0时tail前面没东西,删掉需要维护head,
		if this.fifo.Len() == this.capacity+1 {
			tmp := this.fifo.RemoveBack()
			this.hashMap[tmp.value[0]] = nil
		}
	}
}