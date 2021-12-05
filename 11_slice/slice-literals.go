package main

import "fmt"

func main() {
	q := []int{2, 4, 5, 6, 7, 76}
	fmt.Println(q)

	r := []bool{true, false, false, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{6, false},
		{11, true},
	}
	fmt.Println(s)
}
