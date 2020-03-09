package main

func main() {
	for x := 1; x < 10; x++ {
		for y := 1; y < 10; y++ {
			for z := 1; z < 10; z++ {
				if x > 5 && x < 8 && y > 5 && y < 8 && z > 5 && z < 8 {
					break
				}
			}
		}
	}
	for u := 1; u < 10; u++ {
		for v := 1; v < 10; v++ {
			for w := 1; w < 10; w++ {
				if u > 5 && u < 8 && v > 5 && v < 8 && w > 5 && w < 8 {
					break
				}
			}
		}
	}
}
