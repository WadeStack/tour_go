package main

import "fmt"

func main() {
	i, j := 34, 432
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)
}
