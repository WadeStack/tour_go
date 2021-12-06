package main

import "golang.org/x/tour/reader"

// import "github.com/dupoxy/go-tour-fr/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (int, error) {
	// b = b[:cap(b)]
	for i := range b {
		b[i] = 'A'
	}
	return cap(b), nil
}
func main() {
	reader.Validate(MyReader{})
}
