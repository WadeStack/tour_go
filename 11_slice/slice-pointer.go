package main

import "fmt"

func main() {
	names := [4]string{
		"wade",
		"hider",
		"stack",
		"jack",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)
}
