package main

import (
	"fmt"
	"math"
)

// Square root calculates
func Sqrt(x float64) float64 {
	// initial guess for the answer
	answer := 1.0
	diff := math.Inf(1)

	// calculate until the difference between the answer squared and the original is less than 1*10^-12
	for diff >= math.Pow10(-12) {
		answer -= (answer*answer - x) / (2 * answer)
		fmt.Printf("Guess: %v\n", answer)
		diff = math.Abs(x - math.Pow(answer, 2))
	}
	return answer
}

func main() {
	fmt.Println(Sqrt(2))
}
