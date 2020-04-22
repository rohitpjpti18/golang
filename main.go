package main

import (
	"fmt"
	"github.com/rohitpjpti18/golang-practices/basics"
	)

func main(){
	arr := [...]int{1, 2, 3}


	fmt.Println("Running from the main method")
	basics.BasicArray()
	basics.CompareArrayExample()
	basics.PassArrayByPointer(&arr)
}
