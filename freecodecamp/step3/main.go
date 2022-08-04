package main

import (
	"fmt"
)

func main() {
	grades := [3]int{89, 89, 99}
	fmt.Printf("Grades: %v \n", grades)

	var students [3]string
	fmt.Printf("Studenst: %v \n", students)
	students[0] = "Lisa"
	students[1] = "YO"
	students[2] = "Jim"
	fmt.Printf("Studenst: %v \n", students)
	fmt.Printf("Num of students: %v \n", len(students))

	sliceTest := []int{1, 2, 3}
	fmt.Println(sliceTest)
	fmt.Printf("Length: %v \n", len(sliceTest))
	fmt.Printf("Capacity: %v \n", cap(sliceTest))

	sliceTest2 := append(sliceTest[:2])
	fmt.Println(sliceTest2)

}
