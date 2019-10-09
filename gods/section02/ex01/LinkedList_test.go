// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedlist

import "testing"

func TestAddToHead(t *testing.T) {
	var linkedlist = LinkedList{}
	// 测试添加一个
	linkedlist.AddToHead(3)
	value := linkedlist.headNode.val
	if value != 3 {
		t.Errorf("AddToHead method failed, expect:%d but got %d", 3, value)
	}

	// 测试nil列表
	linkedlist = LinkedList{}
	var headNode = linkedlist.headNode
	if headNode != nil {
		t.Errorf("AddToHead method failed, expect:%T but got %d", nil, value)
	}

	linkedlist = LinkedList{}
	// 测试添加一个
	linkedlist.AddToHead(1)
	linkedlist.AddToHead(4)
	linkedlist.AddToHead(5)
	linkedlist.AddToHead(6)
	value = linkedlist.headNode.val
	if value != 6 {
		t.Errorf("AddToHead method failed, expect:%d but got %d", 6, value)
	}

}

func TestAddToLast(t *testing.T) {
	var linkedlist = LinkedList{}
	linkedlist.AddToEnd(3)
	linkedlist.AddToEnd(4)
	last := linkedlist.LastNode()
	if last.val != 4 {
		t.Errorf("AddToLast method failed, expect %d but got %d", 4, last.val)
	}

	linkedlist = LinkedList{}
	linkedlist.AddToEnd(3)
	linkedlist.AddToHead(4)
	last = linkedlist.LastNode()
	if last.val != 3 {
		t.Errorf("AddToLast method failed, expect %d but got %d", 3, last.val)
	}

	linkedlist = LinkedList{}
	lastNode := linkedlist.LastNode()
	if lastNode != nil {
		t.Errorf("AddToLast method failed, expect %v but got %v", nil, lastNode)
	}

}

func TestNodeWithValue(t *testing.T) {
	var linkedlist = LinkedList{}
	linkedlist.AddToEnd(3)
	linkedlist.AddToEnd(4)
	targetNode := linkedlist.NodeWithValue(3)
	if targetNode.val != 3 {
		t.Errorf("NodeWithValue method failed, expect %d but got %d", 3, targetNode.val)
	}

	targetNode = linkedlist.NodeWithValue(10)
	if targetNode != nil {
		t.Errorf("NodeWithValue method failed, expect %v but got %v", nil, targetNode)
	}
}

func TestAddAfter(t *testing.T) {
	var linkedlist = LinkedList{}
	linkedlist.AddToEnd(3)
	linkedlist.AddToEnd(4)
	linkedlist.AddAfter(3, 5)
	targetNode := linkedlist.NodeWithValue(5)
	if targetNode == nil {
		t.Errorf("TestAddAfter failed, expect %v, but got %v", targetNode, nil)
	}

	linkedlist = LinkedList{}
	linkedlist.AddToEnd(4)
	linkedlist.AddAfter(3, 5)
	targetNode = linkedlist.NodeWithValue(5)
	if targetNode != nil {
		t.Errorf("TestAddAfter failed, expect %v, but got %v", nil, targetNode)
	}
}
