package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for ; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
