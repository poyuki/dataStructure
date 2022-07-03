package pkg

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func (n Node) Next() *Node {
	return n.next
}
func (n Node) Prev() *Node {
	return n.prev
}
