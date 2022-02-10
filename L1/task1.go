package main

import "fmt"

func main() {
	var a, b int = 3, 4

	// fmt.Scanf("%d", &a)
	// fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	//
	// WRITE YOUR CODE HERE
	//

	// First way
	var c int
	c = a
	a = b
	b = c
	fmt.Printf("a = %d, b = %d\n", a, b)

	// Second way
	fmt.Println()
	fmt.Printf("a = %d, b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d, b = %d\n", a, b)

	// Third way
	fmt.Println()
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
