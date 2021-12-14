package main

import "fmt"
import "strings"

func printMessages(Format string, message1 string, message2 string) {
	//fmt.Printf("%s %s\n", message1, message2)

	if len(message1) != 0 && len(message2) != 0 {
		fmt.Printf(Format, strings.Replace(message1, " ", "", -1), message2)
	}
}

func main() {
	const fmt = "%s %s\n"

	for i := 0; 10 > i; i = i + 1 {
		printMessages(fmt, "Hello ", "world")
	}
}
