package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {
	student := Student{1, "www"}
	fmt.Println(student)
	fmt.Println(student.name)
}
