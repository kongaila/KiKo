package main

import (
	"KiKo/tools"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(tools.GenerateUUID())
	}
}
