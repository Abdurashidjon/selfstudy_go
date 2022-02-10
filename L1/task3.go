package main

import (
	"fmt"
)

func main() {
	var r float32 = 10.04

	// fmt.Scanf("%f", &r)
	fmt.Println("R =", r)

	fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float32) (area float32) {

	// WRITE YOUR CODE HERE
	const pi float32 = 3.1415
	area = (4 - pi) * r * r
	return area
	
}
