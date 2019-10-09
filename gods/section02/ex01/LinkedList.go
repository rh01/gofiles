// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedlist

// Node class
type Node struct {
	val  int
	next *Node
}

//related to add, update, remove, and lookup on linked list and
//doubly linked list in the following section.

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// The AddToHead method adds the node to the start of the linked list.
func (l *LinkedList) AddToHead(val int) {
	// fmt.Println("AddTo Head Method")
	var node Node
	node.val = val
	if l.headNode != nil {
		node.next = l.headNode
	}
	l.headNode = &node
}

//The IterateList method of the LinkedList class iterates from
//the headNode property and prints the property of the current head node.
// func (l *LinkedList) IterateList() {
// 	for cur := l.headNode; cur != nil; cur = cur.next {
// 		fmt.Println(cur.val)
// 	}
// }

//The LastNode method of LinkedList returns the node at the end of the list.
func (l *LinkedList) LastNode() *Node {
	cur := l.headNode
	// 处理 cur = nil 的情况
	if cur == nil {
		return nil
	}
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (l *LinkedList) AddToEnd(val int) {
	var node = &Node{}
	node.val = val

	if l.headNode == nil {
		l.headNode = node
	} else {
		cur := l.headNode
		for ; cur.next != nil; cur = cur.next {
		}
		if cur.next == nil {
			cur.next = node
		}
	}
}

//NodeWithValue method returns Node given parameter property
func (l *LinkedList) NodeWithValue(targetValue int) *Node {
	for cur := l.headNode; cur != nil; cur = cur.next {
		if cur.val == targetValue {
			return cur
		}
	}
	return nil
}

//AddAfter method adds a node with nodeProperty after node with property
func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	nodeBefore := l.NodeWithValue(nodeProperty)
	// if not found, it will not be added.
	if nodeBefore != nil {
		var node = &Node{}
		node.val = property
		node.next = nodeBefore.next
		nodeBefore.next = node
	}
}
