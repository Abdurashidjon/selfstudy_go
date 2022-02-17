package main

import (
	"fmt"
	"time"

	"gitlab.com/Udevs/Article/models"
	"gitlab.com/Udevs/Article/storage"
)

func main() {

	articleStorage := storage.NewArticleStorage()

	var a1 models.Article
	a1.ID = 1
	a1.Title = "Lorem"
	a1.Body = "Lorem"
	var p models.Person = models.Person{
		Firstname: "Abdulla",
		Lastname:  "Anvarov",
	}

	a1.Author = p
	t := time.Now()
	a1.CreatedAt = &t

	err := articleStorage.Add(a1)
	if err != nil {
		fmt.Println(err)
	}
	// err = articleStorage.Add(a1)
	// if err != nil  {
	// 	fmt.Println(err)
	// }

	// 2-si
	var a2 models.Article
	a2.ID = 2
	a2.Title = "1111"
	a2.Body = "Uchuch"
	var p1 models.Person = models.Person{
		Firstname: "Aziz",
		Lastname:  "Bahoromov",
	}

	a2.Author = p1
	t1 := time.Now()
	a2.CreatedAt = &t1

	err = articleStorage.Add(a2)
	if err != nil {
		fmt.Println(err)
	}
	// 3-si
	var a3 models.Article
	a3.ID = 3
	a3.Title = "Uchuch"
	a3.Body = "Uchuch"
	var p3 models.Person = models.Person{
		Firstname: "Qobil",
		Lastname:  "Karimberdiyev",
	}

	a3.Author = p3
	t2 := time.Now()
	a3.CreatedAt = &t2

	err = articleStorage.Add(a3)
	if err != nil {
		fmt.Println(err)
	}

	// Id bo'yicha qidirish
	getId, err := articleStorage.GetByID(1)
	if err != nil {
		panic(err)
	}

	// get list
	//getList := articleStorage.GetList()

	// delete by id
	delete := articleStorage.Delete(10)

	// get list hamma danlilarni ob kelish
	getList := articleStorage.GetList()

	// update qilish
	update := articleStorage.Update(models.Article{})

	// search qilish
	search,count := articleStorage.Search("Uchuch")
	if count == 0 {
        panic("Not found search")
	}

	fmt.Println()
	fmt.Println("Get List    : ", getList)
	fmt.Println()
	fmt.Println("Create storage: ", articleStorage)
	fmt.Println()
	fmt.Println("Get by ID: ", getId) // id boyicha olib kelish
	fmt.Println()
	fmt.Println("Delete by ID: ", delete) // delete by id
	fmt.Println()
	fmt.Println("update by ID: ", update)
	fmt.Println()
	fmt.Println("Search by string:  ",count,search)

}
