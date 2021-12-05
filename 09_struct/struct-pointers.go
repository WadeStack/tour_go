package main

import "fmt"

type Student3 struct {
	id   int
	name string
}

func main() {
	student := Student3{1, "www"}
	p := &student
	fmt.Println(*p)
	student.name = "xxx"
	fmt.Println(student)
}
