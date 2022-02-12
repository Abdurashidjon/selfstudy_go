package main

import (
	"fmt"
	"time"
)

func main() {
	dobStr := "03.12.1999" // Replace this date with your birthday
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %s\n", givenDate.Format("02.01.2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday string) {
	//
	// WRITE YOUR CODE HERE
	//
	switch weekDay := time.Date(1999,time.December,03,20,23,23,23,time.Local).Weekday(); weekDay {
	case time.Monday:
		fmt.Println("Dushanba bo'lgan")
	case time.Tuesday:
		fmt.Println("Seshanba bo'lgan")
	case time.Wednesday:
		fmt.Println("Chorshanba bo'lgan")
	case time.Thursday:
		fmt.Println("Payshanba bo'lgan")
	case time.Friday:
		fmt.Println("Juma bo'lgan")
	case time.Saturday:
		fmt.Println("Shanba bo'lgan")
	case time.Sunday:
		fmt.Println("Yakshanba bo'lgan")
	default:
		fmt.Println("Bunaqa xafta kuni yo'q")
	}
	return
}
