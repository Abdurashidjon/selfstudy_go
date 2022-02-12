// break or continue a loop
package main

import "fmt"

func main() {
	sum := 0
	var i int
	for i = 1; i < 1000; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
		if 2550 == sum { // exit the loop
			break
		}
	}
	fmt.Println(i)   // 100
	fmt.Println(sum) // 2550 (2+ 4 + ... + 100)
}
