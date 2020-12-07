package main

import (
	"fmt"
	"time"
)

func main() {
	var t1 int64 = 1000
	t2 := time.Minute * 6
	fmt.Println(t1)
	fmt.Println(int64(t2) * 1000)
	fmt.Println(int64(time.Minute.Milliseconds() * 30 / 1000))
	fmt.Println(1800000000 / 1000 / 1000 / 60)

}
