package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{2, 23, 432, 432, 65, 2, 1, 3 - 7, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	// z := <-c
	fmt.Println(x, y, x+y)
	// fmt.Println(z)

}
