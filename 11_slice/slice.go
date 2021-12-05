package main

import "fmt"

func main() {
	primes := [6]int{1, 6, 5, 4, 3, 2}
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(primes[:])
	//fmt.Println(primes[:7]) //提示越界
	fmt.Println(primes[:6])
	fmt.Println(primes[:5])
	fmt.Println(primes[0:6])
	fmt.Println(primes[1:6])
	fmt.Println(primes[0:5])
	fmt.Println(primes[0:1])
	fmt.Println(primes[0:0])
}
