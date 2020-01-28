package main

import (
	"awesomeProject/dir1"
	"fmt"
)

const (WHITE  = iota
	BLACK
	BLUE
	RED
	YELLOW)

type Color byte


func main(){
	fmt.Println("entry1")
	dir1.Say1()
}