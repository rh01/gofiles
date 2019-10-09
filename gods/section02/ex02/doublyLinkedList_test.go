package doublyLinkedList

import "testing"

func TestAddToHead(t *testing.T) {
	var dll = DoubleLinkedList{}
	headNode := dll.headNode
	if headNode != nil {
		t.Errorf("AddToHead method failed, expect %v but got %v", nil, headNode)
	}

	dll.AddToHead(3)
	headNode = dll.headNode
	if headNode.property != 3 {
		t.Errorf("AddToHead method failed, expect %d but got %d", 3, headNode.property)
	}

	dll.AddToHead(4)
	headNode = dll.headNode
	if headNode.property != 4 {
		t.Errorf("AddToHead method failed, expect %d but got %d", 4, headNode.property)
	}
}

func TestNodeBetweenValues(t *testing.T) {
	var dll = DoubleLinkedList{}
	headNode := dll.headNode
	if headNode != nil {
		t.Errorf("AddToHead method failed, expect %v but got %v", nil, headNode)
	}

	dll.AddToHead(3)
	headNode = dll.headNode
	if headNode.property != 3 {
		t.Errorf("AddToHead method failed, expect %d but got %d", 3, headNode.property)
	}

	dll.AddToHead(4)
	headNode = dll.headNode
	if headNode.property != 4 {
		t.Errorf("AddToHead method failed, expect %d but got %d", 4, headNode.property)
	}

	headNode = dll.NodeBetweenValues(3, 4)
	if headNode != nil {
		t.Errorf("NodeBetweenValues method failed, expect %v but got %v", nil, headNode)
	}

	dll.AddToHead(5)
	headNode = dll.NodeBetweenValues(5, 3)
	if headNode.property != 4 {
		t.Errorf("NodeBetweenValues method failed, expect %d but got %d", 4, headNode.property)
	}

}
