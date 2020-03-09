package main

import "fmt"

func main() {
	fmt.Printf("foo")
	fmt.Printf("foo", "bar")
	fmt.Printf("%d", "bar")
	fmt.Printf("%s %s", "bar")
	fmt.Printf("%d", 3.14)
	fmt.Printf("%p", nil)
}
