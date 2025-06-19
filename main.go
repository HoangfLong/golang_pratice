package main

import (
	"fmt"
	"go_tutorials/array"
	"go_tutorials/linkedlist"
	"go_tutorials/queues"
	"go_tutorials/stacks"
)

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func main() {

	// myChan := make(chan int)

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		time.Sleep(time.Second * 1)
	// 		myChan <- i
	// 	}
	// }()

	// go receiveAndSend(myChan)
	// myChan <- 1

	// fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)

	// ch := make(chan int, 10)

	// go func() {
	// 	for i := 1; i <= 100; i++ {
	// 		ch <- i
	// 	}
	// }()

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(<-ch)
	// 	time.Sleep(time.Second * 1)
	// }

	// var ptr *int
	// n := 5 // &n 0052156x
	// ptr = &n
	// fmt.Print(*ptr)

	//Array Bubble sort
	arr := []int{7, 12, 9, 4, 11}
	fmt.Print(&arr[0], &arr[1])
	sorted := array.BubbleSort(arr)
	fmt.Println(sorted)

	fmt.Println("hoho")

	//Stack
	myStack := stacks.Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(500)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

	fmt.Println(myStack.Peek())
	//Queues
	myQueues := queues.Queues{}
	myQueues.Enqueues(200)
	myQueues.Enqueues(51200)
	myQueues.Enqueues(151)
	fmt.Println("My queues: ", myQueues)
	myQueues.Dequeues()
	fmt.Println("My queues: ", myQueues)

	//Linkedlist
	myList := linkedlist.LinkedList{}
	node1 := &linkedlist.Node{Data: 48}
	node2 := &linkedlist.Node{Data: 18}
	node3 := &linkedlist.Node{Data: 58}
	node4 := &linkedlist.Node{Data: 58}
	node5 := &linkedlist.Node{Data: 58}

	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Prepend(node4)
	myList.Prepend(node5)
	myList.PrintListData()
}
