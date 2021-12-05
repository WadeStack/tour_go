package main

import "fmt"

func swap(x,y string)(string, string, int){
	return y,x,1
}

func main() {
	a,b,c := swap("hello","hi")
	fmt.Println(a,b,c)
}