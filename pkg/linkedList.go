package pkg

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data interface{}) *LinkedList {
	node := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = node
	} else {

		next := l.head

		if next.next != nil {
			next = next.next
		}
		next.next = node
	}

	return l
}

func (l *LinkedList) Prepend(data interface{}) *LinkedList {
	node := &Node{
		data: data,
	}

	if l.head != nil {
		node.next = l.head
	}
	l.head = node

	return l
}

func (l LinkedList) Next() *Node {
	return l.head
}
