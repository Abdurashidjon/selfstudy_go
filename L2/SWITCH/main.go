// https://yourbasic.org/golang/switch-statement/
package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic switch with default
	switch today := time.Now().Weekday(); today {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}

	// No condition
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Case list
	str := "Hello, World\n"
	countWhiteSpace := 0
	for i := 0; i < len(str); i++ {
		if WhiteSpace(rune(str[i])) {
			countWhiteSpace++
		}
	}
	fmt.Printf(str)
	fmt.Println("countWhiteSpace = ", countWhiteSpace)

	// Fallthrough
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough // A fallthrough statement transfers control to the next case.
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

	// Exit with break
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break // A break statement terminates execution of the innermost for, switch, or select statement
		case '\n': // break at newline
			break Loop // If you need to break out of a surrounding loop, not the switch, you can put a label on the loop and break to that label
		default:
			fmt.Printf("%c\n", ch)
		}
	}

	// Execution order
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
}

// Case list
func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

// Foo prints and returns n.
func Foo(n int) int {
	fmt.Println(n)
	return n
}
