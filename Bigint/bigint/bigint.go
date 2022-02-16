package bigint

import (
	"errors"
	"regexp"
	"strconv"
)

type Bigint struct {
	Value string
}

var ErrorBadInput = errors.New("bad input, please input only number")

// This function checks the input input to the given regexp
func checkInput(num string) (bool, error) {
	if match, err := regexp.MatchString(`^[+-]?[0-9]*$`, num); err != nil {
		return false, err
	} else {
		return match, nil
	}
}

// Return a new value
func Newint(num string) (Bigint, error) {
	var f bool
	var err error
	if f, err = checkInput(num); err != nil {
		return Bigint{}, err
	}
	if !f {
		return Bigint{}, ErrorBadInput
	}
	return Bigint{Value: num}, nil
}

// Set function a dagi sondi borib o'zgatirib keluvchi funksiya
func (z *Bigint) Set(num string) error {
	var f bool
	var err error

	if f, err = checkInput(num); err != nil {
		return err
	}
	if !f {
		return ErrorBadInput
	}
	z.Value = num
	return nil
}

// Sum two numbers
func Add(a, b Bigint) Bigint {
	// my code
	int1, _ := strconv.Atoi(a.Value) // validatsiyadan o'tkan sondi integerga o'tkazayabmiz
	int2, _ := strconv.Atoi(b.Value)
	Sum := int1 + int2
	unpars := strconv.Itoa(Sum) // hosil bo'lgann integerni qaytadan stringa o'tkazayabmiz
	return Bigint{Value: unpars}
}

func Sub(a, b Bigint) Bigint {
	// my code
	int1, _:= strconv.Atoi(a.Value)
	int2,_ := strconv.Atoi(b.Value)
	sub := int1 - int2
	unpars := strconv.Itoa(sub)
	return Bigint{Value: unpars}
}

func Multiply(a, b Bigint) Bigint {
	int1, _ := strconv.Atoi(a.Value)
	int2, _ := strconv.Atoi(b.Value)
	mult := int1 * int2
	unpars := strconv.Itoa(mult)
	return Bigint{Value: unpars}
}

func Mod(a, b Bigint) Bigint {
	int1, _ := strconv.Atoi(a.Value)
	int2, _ := strconv.Atoi(b.Value)
	var mod int
	if int1 > int2 {
		mod = int1 % int2
	} else {
		mod = int2 % int1
	}
	unpars := strconv.Itoa(mod)
	return Bigint{Value: unpars}
}

// func Sum(a, b int) (int, error) {
// 	return a + b, nil
// }
