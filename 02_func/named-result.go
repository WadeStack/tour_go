package main

import "fmt"

func split(sum int)(x int){
	x = sum*4-9
	y := sum -x
	fmt.Println(x)
	fmt.Println(y)
	return
}

func main() {
	fmt.Println(split(11))
}
