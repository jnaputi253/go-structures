package list

import "fmt"

type node struct {
	next *node
	prev *node
	data int
}

func emptyNode() *node {
	return &node{
		next: nil,
		prev: nil,
		data: -1,
	}
}

func newNode(data int) *node {
	return &node{
		next: nil,
		prev: nil,
		data: data,
	}
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func New() *LinkedList {
	linkedList := &LinkedList{
		head: emptyNode(),
		tail: emptyNode(),
		size: 0,
	}

	linkedList.head.next = linkedList.tail
	linkedList.tail.prev = linkedList.head

	return linkedList
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Clear() {
	if l.size == 0 {
		return
	}

	for tempNode := l.head.next; tempNode != l.tail; tempNode = l.head.next {
		tempNode.next.prev = tempNode.prev
		tempNode.prev.next = tempNode.next
		tempNode.next = nil
		tempNode.prev = nil
		l.size--
	}
}

func (l *LinkedList) AddToFront(data int) {
	node := newNode(data)
	node.next = l.head.next
	node.prev = l.head
	node.next.prev = node.prev
	node.prev.next = node.next

	l.size++
}

func (l *LinkedList) PrintForward() {
	if l.size == 0 {
		fmt.Println("Empty")
		return
	}

	fmt.Println("Set")

	for currNode := l.head.next; currNode != l.tail; currNode = currNode.next {
		fmt.Println(currNode.data)
	}

	fmt.Println("Done")
}
