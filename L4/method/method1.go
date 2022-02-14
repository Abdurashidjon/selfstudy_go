package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Bu funksiya orqali:", Gipaten(3, 4))
	fmt.Println("Bu method orqali  :", Gipatenuza{1, 1}.Tomon())
	fmt.Println("Bu method funcga num berganmiz:",Gipatenuza{15,20}.Tomon1(10))
}

// biz method va funksiyani ko'rib chiqamiz method biror declare bilan elon qilinadi
// oddiy funksiya uchun misol
func Gipaten(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

// tepadagi funksiyani method orqali yozamiz
type Gipatenuza struct {
	a float64
	b float64
}

func (gip Gipatenuza) Tomon() float64 {
	return math.Sqrt(gip.a*gip.a + gip.b*gip.b)
}

// bu yozgan kodimizda Tomon funskisayini oziga xam ozgaruvchi berib ketganmiz
func (gip Gipatenuza) Tomon1(num float64) float64 {
	return math.Sqrt(gip.a*gip.a+gip.b*gip.b) + num
}
