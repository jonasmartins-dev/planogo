package main

import (
	"fmt"
)

func main() {

var a int = 7
var b *int
b = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
	
	*b++
	
	fmt.Println(a)
	
}