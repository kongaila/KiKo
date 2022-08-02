package main

import (
	"fmt"
)

type B struct {
	Name string
}

func (b B) Test1() {
	fmt.Printf("Test1 addr:%p\n", &b)
	fmt.Printf("Test1 name:%s\n", b.Name)
	b.Name = "john"
}
func (b *B) Test2() {
	fmt.Printf("Test2 addr:%p\n", b)
	fmt.Printf("Test2 name:%s\n", b.Name)
	b.Name = "john"
}
func main() {
	b := B{}
	b.Test1()
	b.Test1()
	b.Test2()
	b.Test2()
}
