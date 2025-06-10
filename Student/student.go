package Student

import "fmt"

type Student struct {
	ID    int
	Age   int
	Name  string
	Major string
}

func AddStudent(students *[]Student, id int, name string, age int, major string) {
	newStudent := Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Major: major,
	}
	*students = append(*students, newStudent)
}

func ListStudent(students []Student) {
	for _, s := range students { // for index, value := range array|slice|map
		//s là từng student trong slice
		fmt.Println(s.ID, s.Name, s.Age, s.Major)
	}
}
