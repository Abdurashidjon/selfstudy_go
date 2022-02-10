package main

import "fmt"

func main() {
	var a, b int = 10, 12

	// fmt.Scanf("%d", &a)
	// fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Println(a, "is odd: ", isOdd(a))
	fmt.Println(b, "is even: ", isEven(b))
}

func isEven(num int) bool {

	// WRITE YOUR CODE HERE
	if (num%2 == 0) {
		return true
	}
	return false
}

func isOdd(num int) bool {
	
	// WRITE YOUR CODE HERE
	if (num%2 == 1){
		return  true
	}
	return false
}
