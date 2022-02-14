package main

import (
	"fmt"
)

func main() {

	// Bu yerda biz arrayni oldin slice qilib keyin append qilishni yozib ketganmiz
	// arr := [...]int{1, 2, 3,4}
	// fmt.Println(arr)
	// arr1  := append(arr[:],10)
	// arr1 = append(arr1, 20)
	// fmt.Println(arr1)

	// Bu codda arrayni slice qilib kesib olishni ko'rganmiz uni avval [:] belgilari orqali slice qilib ketganmiz
	// arr := [6]int{3, 6, 2, 8, 1, 11}
	// slic := arr[:4]
	// fmt.Println(slic)
	// slic = append(slic,175)
	// fmt.Println(slic)

	// Bu codda slice ni arraydan farqlab olamiz
	arr := [...]int{1, 2, 4, 8} // bu array olchami aniq emas bolgan array slice emas
	fmt.Println(arr)
	// arr = append(arr,56) biz bunaqa qilib arrayga append qila olmemiz chunki array tepada 4 o'lcham bolib oldi
	slice := []int{1, 2, 4, 8} // endi bu	slice boldi chunki [] ichi bow demak bu slice ekanini bildiradi va bunga append qilish mumkin
	slice = append(slice, 12)
	fmt.Println(slice)

	// Endi esa length va capacitylar xaqida tanishamiz
	slice1 := []int{1, 2, 4}
	fmt.Println("Slice 1:  ", slice1)
	fmt.Printf("length: %v capacity: %v\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 12)
	fmt.Printf("length: %v capacity: %v\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 15)
	fmt.Printf("length: %v capacity: %v\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 15)
	fmt.Printf("length: %v capacity: %v\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 15)
	fmt.Printf("length: %v capacity: %v\n", len(slice1), cap(slice1))
	// Bu yerda length bilan capacityni farqini qande farqlash mumkin degan savolga
	// length bn capacity xam length faqat ikki xildir. Capacity length bilan tenglashgan payti
	// ikki barobarga o'zini uzunligini o'zgartirib oladi. Yana len = cap tenglashgan payti
	// ikki barobar ozini orttirib yuboradi

	slice2 := []int{3}
	fmt.Println(slice2)
	fmt.Printf("length: %v capacity: %v\n",len(slice2),cap(slice2))
	slice2 = append(slice2, 21)
	fmt.Println(slice2)
	fmt.Printf("length: %v capacity: %v\n",len(slice2),cap(slice2))
	slice2 = append(slice2, 91)
	fmt.Println(slice2)
	fmt.Printf("length: %v capacity: %v\n",len(slice2),cap(slice2))



}
