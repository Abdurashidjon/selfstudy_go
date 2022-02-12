package main

import (
	"fmt"

	"gitlab.com/Udevs/L2/functions/samples"
)

func main() {
	x := 4
	y := 5

	z := samples.Add(x, y)
	f := samples.Add

	fmt.Println(z)
	fmt.Println(f(x, y))

	fmt.Println(samples.Sub(x, y))
	fmt.Println(samples.Inc(x, y, z))
	fmt.Println(samples.Threerandom())

	sum := func(a, b, c int) int {
		return a + b + c
	}(3, 5, 7)
	fmt.Println(sum)                     // 15

	fmt.Println(samples.Sum(1,2,3))      // 6
	fmt.Println(samples.Sum(1,2,3,4,5))  // 15
	fmt.Println(samples.Sum(2,3,7,4,12)) // 28

	fmt.Println(samples.Fact(3))         // 6
	fmt.Println(samples.Fact(7))         // 5040
	fmt.Println(samples.Fact(10))        // 3628800
}
