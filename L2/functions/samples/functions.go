package samples

import (
	"math/rand"
	"time"
)

// Ikki sondi qo'shish
func Add(a, b int) int {
	return a + b
}

// Ikki sondi ayirish
func Sub(x, y int) int {
	return x - y
}

// Sonlarni oshirish
func Inc(x, y, z int) (a, b, c int) {
	a = x + 1
	b = y + 1
	c = z + 1
	return
}

// Random orqali 3 ta sondi qaytarish
func Threerandom() (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	z := rand.Intn(10)
	return x, y, z
}

// Arraydagilarni yig'indisi
func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// O'zini o'zi chaqiruvchi funksiya
func Fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Fact(n-1)
}
