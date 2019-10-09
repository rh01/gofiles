// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublyLinkedList

type Node struct {
	property int
	prevNode *Node
	nextNode *Node
}

type DoubleLinkedList struct {
	headNode *Node
}

// The following sections explain doubly linked list methods,
// such as the NodeBetweenValues, AddToHead, AddAfter, AddToEnd, and main methods.

// The NodeBetweenValues method of the LinkedList class returns the node that
// has a property lying between the firstProperty and secondProperty values.
func (d *DoubleLinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var nodeWith *Node
	for cur := d.headNode; cur != nil; cur = cur.nextNode {
		if cur.prevNode != nil && cur.nextNode != nil {
			if cur.prevNode.property == firstProperty && cur.nextNode.property == secondProperty {
				nodeWith = cur
				break
			}
		}
	}
	return nodeWith
}

//The AddToHead method of the doubly LinkedList class
//sets the previousNode property of the current headNode of the
//linked list to the node that's added with property.
func (d *DoubleLinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	if d.headNode != nil {
		node.nextNode = d.headNode
		d.headNode.prevNode = node
	}
	d.headNode = node
}

//AddAfter method of LinkedList
func (d *DoubleLinkedList) AddAfter(nodeProperty int, property int) {
	for cur := d.headNode; cur != nil; cur = cur.nextNode {
		if cur.property == nodeProperty {
			var node = &Node{}
			node.property = property
			// main logic
			node.nextNode = cur.nextNode
			node.prevNode = cur
			cur.nextNode.prevNode = node
			cur.nextNode = node
			break //这里需要break???
		}
	}
}
