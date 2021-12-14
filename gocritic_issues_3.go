package main

import "fmt"

func new(len int, cap int) ([]int, int) {
	vals := make([]int, len)
	for i := 0; cap > i; i = i + 1 {
		vals = append(vals, i)
		vals = append(vals, i*2)
	}
	return vals, len + cap
}

func main() {
	vals, _ := new(0, 010)
	fmt.Println(vals)
	for i := 0; i < len(vals); i++ {
		vals[i] = 0
	}

}
