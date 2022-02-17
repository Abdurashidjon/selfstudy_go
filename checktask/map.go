package main

import (
	"fmt"
	"math"
)

type Add struct {
	yosh float64
	buy  float64
}

func main() {
	var m1 map[string]Add
	m1 = map[string]Add{
		"Abdurashid": {22, 175},
	}
	fmt.Println(m1)
	fmt.Println(math.Pi)
}
