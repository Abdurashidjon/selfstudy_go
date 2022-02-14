package main

import (
	"fmt"
	"math"
)

type Qiymatlar struct {
	A float64
	B float64
}
func (m Qiymatlar) Abs() float64 {
	return math.Sqrt(m.A*m.A + m.B*m.B)
}
// func (m *Qiymatlar) Oshir(num float64) {
// 	m.A = m.A * num
// 	m.B = m.B * num
// }
func Oshir(m *Qiymatlar,num float64) {
	m.A = m.A * num
	m.B = m.B * num 
}

func main() {
	v := Qiymatlar{3,4}
	Oshir(&v,10)
	fmt.Println(v.Abs())
}
