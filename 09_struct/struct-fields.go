package main

import "fmt"

type Student2 struct {
	id   int
	name string
}

func main() {
	student := Student2{1, "www"}
	fmt.Println(student)
	student.name = "xxx"
	fmt.Println(student.name)
}
