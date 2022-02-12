package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(a int) {
	//
	// WRITE YOUR CODE HERE
	//
	if a%15 == 0 {
		fmt.Println(a," bu son FizzBuzz")
	} else if a%3 == 0 {
		fmt.Println(a," bu son Fizz")
	} else if a%5 == 0 {
		fmt.Println(a," bu son Buzz")
	}
}
