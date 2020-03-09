package main

import "fmt"
import "bytes"

func main() {
	a := []byte{1, 2, 3}
	b := []byte{3, 4, 5}

	c := bytes.Equal(a, b)
	fmt.Println(c)

	c = bytes.Equal(a, a)
	fmt.Println(c)
}
