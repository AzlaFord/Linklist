package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func Linklis() {
	node1 := &Node{value: 10}
	node2 := &Node{value: 20}
	node3 := &Node{value: 30}
	head := LinkedList{head: node1}
	node1.next = node2
	node2.next = node3
	head.Append(40)
	head.Print()
}

func (l *LinkedList) Print() {
	if l.head == nil {
		return
	}

	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
