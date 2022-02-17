package main

import (
	"fmt"
	"math"
)

var sum int

func adder() func(int) int {
	return func(x int) int {
		sum += x
		return sum
	}
}

func Sqrt(a float64) func(float64, float64) float64 {
	return func(f1, f2 float64) float64 {
		return a*math.Sqrt(f1*f1 + f2*f2)
	}
}

func main() {
	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }

	sqrt := Sqrt(10)
	fmt.Println(sqrt(3,4))
}
