package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 8, 4, 23, 23, 78}
	fmt.Println(primes)
}
