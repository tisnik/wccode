package main

import "fmt"
import "math"

type Vector struct {
	X, Y float64
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var v *Vector
	fmt.Println(v.Length())
}
