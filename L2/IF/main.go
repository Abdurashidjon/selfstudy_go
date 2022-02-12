package main

import "fmt"

func main() {
	x := 2
	if num := 10; num%x == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	x, y, z := 3, 5, 4
	max := z
	if a := f(x, y); a > z { // With init statement
		max = a
	}

	fmt.Printf("max of {%d,%d,%d} is %d\n", x, y, z, max)
	fmt.Printf("min of {%d,%d,%d} is %d\n", x, y, z, findMin(x, y, z))

	flag := isValidTriangle(x, y, z)
	if flag {
		fmt.Printf("{%d,%d,%d} is VALID triangle\n", x, y, z)
	} else {
		fmt.Printf("{%d,%d,%d} is NOT triangle\n", x, y, z)
	}
}

// Basic if syntax
func f(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Nested if statements
func findMin(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}

// If else statement
func isValidTriangle(a, b, c int) bool {
	if a+b <= c || a+c <= b || b+c <= a {
		return false
	} else {
		return true
	}
}