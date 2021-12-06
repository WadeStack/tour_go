package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"wade", 21}
	z := Person{"Hider", 20}
	fmt.Println(a, z)
	// println(a, z)
}
