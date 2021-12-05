package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m1 = map[string]Vertex1{
	"a": Vertex1{10.01521, 44.522},
	"b": Vertex1{18.01521, 45.522},
}

func main() {
	fmt.Println(m1)
}
