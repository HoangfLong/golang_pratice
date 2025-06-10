package main

import (
	"fmt"
	"go_tutorials/Product"
	"go_tutorials/Student"
	"go_tutorials/basic"
)

// func add(x int, y int) int {
// 	return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// type Test1 struct {
// 	x, y int
// }

func main() {

	// fmt.Println("struct: ", Test1{1, 2})

	// fmt.Println("Hello World")

	// var intNum int
	// fmt.Println(intNum)

	// fmt.Println(add(12, 15))

	// var intNum int = 15

	// fmt.Println(intNum)

	// ho, he := swap("bomb", "ambato")
	// fmt.Println(ho, he)

	// var firstName, lastName, middleName string = "Long", "Duong", "Hoang"
	// // var lastName string = "Duong"
	// var intNum int = 10

	// var (
	// 	a int
	// 	b int    = 5
	// 	c string = "hehe"
	// )

	// pi := 2.5

	// fmt.Println(
	// 	firstName,
	// 	lastName,
	// 	middleName,
	// 	intNum,
	// 	pi,
	// )

	// fmt.Println(
	// 	a,
	// 	b,
	// 	c,
	// )

	// const bomb string = "ambatobomb"

	// fmt.Println("day la bomb", bomb)

	// fmt.Println(math.Pi)

	// //Array
	// //1
	// var arr1 = [3]int{1, 2, 3}

	// //2
	// arr2 := [3]float32{1.5, 5.5, 6.5}

	// fmt.Println("array 1:", arr1)

	// fmt.Println("array 2:", arr2)

	// //Slice
	// long_slice := []int{1, 5, 4, 5}

	// fmt.Println(len(long_slice))
	// fmt.Println(cap(long_slice))

	// //Loop
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// //Pointer
	// var ptr *int
	// x := 10
	// ptr = &x

	// var ptr2 *int

	// ptr2 = ptr

	// fmt.Println(ptr2)

	// fmt.Println("Value stored in variable x:", x) //10
	// // fmt.Println("Address of variable x", &x)         //address of x
	// fmt.Println("Address of variable ptr:", &ptr)    //address of x
	// fmt.Println("Value stored in variable ptr", ptr) //

	// func area(r float64) float64 {
	// 	return pi*r*r
	// }

	//student
	var students []Student.Student
	Student.AddStudent(&students, 1, "Alice", 20, "Computer Science")
	Student.AddStudent(&students, 2, "Hoang Long", 22, "Sofware Engineer")

	Student.ListStudent(students)

	//1. Tính tổng
	var sum int = 0

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum:", sum)

	//2. Đảo ngược chuỗi
	var s = basic.Reverse("hoanglong")
	fmt.Println("Reverse string", s)

	//Pointer
	var numb int = 5

	var ptr *int

	ptr = &numb

	*ptr = *ptr + 5

	fmt.Println("total:", *ptr)
	fmt.Println("variable")

	array := [4]int{1, 2, 3, 4}
	slices := array[1:3]
	fmt.Println("s:", slices)

	var products []Product.Product

	Product.AddProduct(&products, 1, "bomb", 5, 2.33)

	fmt.Println(products)

	//Control flow
	if 20 > 30 {
		fmt.Println("right:")
	} else {
		fmt.Println("false")
	}
	//Loop
	for i := 0; i < 10; i++ {
		fmt.Println(&i)
	}
	//range
	cars := [3]string{"bomb", "cas", "das"}
	for idx, val := range cars {
		fmt.Println(idx, val)
	}
	fmt.Print("hoho")
}
