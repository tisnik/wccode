package main

import "fmt"
import "bytes"

func main() {
	a := []byte{1, 2, 3}
	b := []byte{3, 4, 5}

	c := bytes.Compare(a, b) == 0
	fmt.Println(c)

	c = bytes.Compare(a, a) == 0
	fmt.Println(c)
}
