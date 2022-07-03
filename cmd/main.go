package main

import "dataStructure/pkg"

func main() {

	list := pkg.LinkedList{}
	dlist := pkg.DoubleLinkedList{}
	stack := pkg.Stack{}

	list.Append(1)
	dlist.Append(1)
	list.Append(2)
	dlist.Append(2)
	list.Append(3)
	dlist.Append(3)
	list.Append(4)
	dlist.Append(4)

	list.Prepend(5)
	dlist.Prepend(5)
	list.Prepend(6)
	dlist.Prepend(6)
	list.Prepend(7)
	dlist.Prepend(7)
	stack.Push(1)
	stack.Push(2)
	t := stack.Pop()
	stack.Push(t)

}
