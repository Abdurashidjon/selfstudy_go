package main

import (
	"fmt"
)

func main() {
	fmt.Println(MySquareRoot(10))
}

func MySquareRoot(num uint) (result float64) {
	// DO NOT USE math.Sqrt!

	//
	// WRITE YOUR CODE HERE
	//
	var sr float64 = float64(num / 2)
	var temp float64
	for {
		temp = sr
		sr = (temp + (float64(num) / temp)) / 2
		if temp == sr {
			break
		}
	}
	return sr
}
