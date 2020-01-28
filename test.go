package main

import (
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func printBook(book *Books){
	println(book)
	fmt.Printf("Book title is : %s\n", book.title)
	fmt.Printf("Book author is : %s\n", book.author)
	fmt.Printf("Book subject is : %s\n", book.subject)
	fmt.Printf("Book id is : %d\n", book.book_id)
}

func main()  {
	//var book1 Books
	//var book2 Books
	//
	//book1.title = "Go 言語"
	//book1.author= "xxxx"
	//book1.subject = "program"
	//book1.book_id = 123456
	//
	//book2.title = "Go2 言語"
	//book2.author = "xxxx2"
	//book2.subject = "program2"
	//book2.book_id = 567890
	//
	//printBook(&book1)
	//printBook(&book2)
	//
	//shuzu()
	//slice()

	type ages int
	type money float32
	var a ages = 10
	var b money =15.32
	fmt.Printf("%d, %f", a, b)
}

func shuzu() {
	var balance [10] int
	balance[0] = 1
	balance[1] = 2
	balance[2] = 3

	var a1 = [5]int{}
	a1[1] = 1

	var a2 = []int{1}
	a2[0]=2

	s :=  []int{1,2,3}
	s[2] = 4

	fmt.Println(balance)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(s)

}

func slice(){
	arr := [...]int{1,2,3,4}

	s:=arr[0:3]

	println(&arr)
	println(&s)
	println(s)
}