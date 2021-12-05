package main

import "fmt"

type Student4 struct {
	id   int
	name string
}

var (
	student1 = Student4{1, "test1"}
	student2 = Student4{id: 5}
	student3 = Student4{}
	p        = &Student4{2, "test2"}
)

func main() {
	fmt.Println(student1, *p, student2, student3)
}
