package main

import "fmt"

type User struct {
	Surname string
	Name string
	Birthday string
	Email string
	Phone string
	Level string
	Passpord string
	Staff_photo string
}

func main() {
	user := User {
		Surname: "Gaybullaev",
		Name: "Abdurashid",
		Birthday: "03.12.1999",
		Email: "abdurashidgaybullayev@gmail.com",
		Phone: "+998941967799",
		Level: "Junior backend developer",
		Passpord: "I have got passpord",
		Staff_photo: "No photo",
	}
	fmt.Println("Welcome to Udevs ",user)
}