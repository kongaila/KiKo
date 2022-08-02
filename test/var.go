package main

import (
	"KiKo/datamodels/domain"
	"fmt"
)

func main() {
	c1 := domain.TbUser{}
	var c2 domain.TbUser
	c3 := new(domain.TbUser)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}
