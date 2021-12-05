package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex2{
	"a": {10.01521, 44.522},
	"b": {18.01521, 45.522},
}

func main() {
	fmt.Println(m2)
	fmt.Println(m2["b"])
}
