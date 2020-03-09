package main

import "fmt"

func main() {
	fmt.Sprintf("foo")
	fmt.Sprintf("foo", "bar")
	fmt.Sprintf("%d", "bar")
	fmt.Sprintf("%s %s", "bar")
	fmt.Sprintf("%d", 3.14)
	fmt.Sprintf("%p", nil)
}
