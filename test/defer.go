package main

import (
	"fmt"
)

func main() {
	var slice *[]int
	defer func() {
		fmt.Println("defer: ", *slice)
	}()
	slice = &[]int{1, 2, 3}
	fmt.Println(*slice)
}
