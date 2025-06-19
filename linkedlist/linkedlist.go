package linkedlist

import "fmt"

type Node struct {
	Data int16
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf(" %d", toPrint.Data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
