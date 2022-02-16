package main

import (
	"fmt"

	h1 "gitlab.com/Udevs/Bigint/bigint"
)

func main() {

	a, err := h1.Newint("21")
	if err != nil {
		panic(err)
	}
	b, err := h1.Newint("23466")
	if err != nil {
		panic(err)
	}

	// err = a.Set("2")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(h1.Add(a, b))      // bigintdagi Add funksiyasini chaqirdim ikki sondi qoshadi
	fmt.Println(h1.Sub(a,b))       // ikki sondi ayiradi
	fmt.Println(h1.Multiply(a, b)) // Multiply funksiyasini chaqirdim ikki stringni ko'paytiradi
	fmt.Println(h1.Mod(a,b))       // Mod funksiyasini chaqirib kattasining kichikiga bo'lib qoldiqni qaytarayabman
}
