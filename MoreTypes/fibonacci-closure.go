package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 0
	return func() int {
		if x == 0 && y == 0 {
			y++
			return 0
		}
		sum := x + y
		x = y
		y = sum
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
