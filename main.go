package main

import (
	"fmt"
	"go_tutorials/array"
	"go_tutorials/leetcode"
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
	arr := []int{40, 7, 2, 3, 1}
	// fmt.Print(&arr[0], &arr[1])
	// bubbleSort := array.BubbleSort(arr)
	// fmt.Println(bubbleSort)
	//Selections Sort
	selectionSort := array.SelectionSort(arr)
	fmt.Println("Selection sort: ", selectionSort)

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

	numb := 5
	array := []int{4, 5, 6, 4, 7, 8, 9, 8, 4}
	// for index, arr1 := range array {
	// 	if arr1 == numb {
	// 		fmt.Println("phan tu cua mang tai: ", index, arr1)
	// 		break
	// 	}
	// }
	for i := 0; i < len(array); i++ {
		if array[i] == numb {
			fmt.Println("Phan tu bang tai array", i)
		}
	}

	nums := []int{3, 2, 4}
	target := 6

	twoSum := leetcode.TwoSum(nums, target)
	fmt.Println("geeg", twoSum)
}
