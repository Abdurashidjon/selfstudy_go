package main

import (
	"fmt"
	"strconv"
)

func main() {
	// i := "23"
	// j := "132"
	// fmt.Println(i,j)
	// fmt.Printf("type i: %T\n", i)
	// fmt.Printf("type j: %T\n", j)
	// sum1 := i + j
	// fmt.Println("i + j = ",sum1)

	// ii, _ := strconv.Atoi(i)
	// ij, _ := strconv.Atoi(j)
	// fmt.Printf("ii type %T\n", ii)
	// fmt.Printf("ij type %T\n", ij)
	// sum2 := ii + ij
	// fmt.Println("ii + ij = ",sum2)

	str1 := "20"

	int1, _ := strconv.ParseInt(str1, 0, 0)
	fmt.Printf("type int1: %T\n", int1)
}
