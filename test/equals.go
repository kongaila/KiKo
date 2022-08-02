package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("asda", ""))
	fmt.Println(strings.EqualFold("111", "111"))
	fmt.Println(strings.EqualFold("ff", "f"))
	fmt.Println(strings.EqualFold("\n\n", ""))
}
