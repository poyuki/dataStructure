package pkg

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (dl *DoubleLinkedList) Append(data interface{}) *DoubleLinkedList {
	node := &Node{
		data: data,
	}

	if dl.tail != nil {
		dl.tail.next = node
		node.prev = dl.tail
		dl.tail = node
	} else {
		dl.head = node
		dl.tail = node
	}

	return dl
}

func (dl *DoubleLinkedList) Prepend(data interface{}) *DoubleLinkedList {
	node := &Node{
		data: data,
	}

	if dl.head != nil {
		dl.head.prev = node
		node.next = dl.head
		dl.head = node
	} else {
		dl.head = node
		dl.tail = node
	}

	return dl
}
