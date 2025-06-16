package main

import (
	"fmt"
	"go_tutorials/array"
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
	sorted := array.BubbleSort(arr)
	fmt.Println(sorted)

}
