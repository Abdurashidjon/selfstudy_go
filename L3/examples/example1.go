package main

import "fmt"

func main() {
	var a int = 100
	var b *int = &a

	fmt.Printf("Value of a: %v\nValue of b: %v\n", a, *b)
	fmt.Printf("Adress of a: %v\nAdress of b: %v\n", &a, b)

	*b = 260
	fmt.Println("New a", a)  // & shu o'rinda aytib ketish joizki ampersand belgisi pointer ya'ni o'zida ko'rsatilgan adersni saqlab qoladi
	fmt.Println("New b", *b) // * bu pointer esa ko'rsatilgan adresdagi qiymatni ushlab oladi yangisida o'zlashtiradi ya'ni o'sha adresdagi qiymatni oladi

	c := *b
	fmt.Println(c)
}
