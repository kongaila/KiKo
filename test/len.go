package main

import "fmt"

func main() {
	topSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(topSlice) >= 10 {
		topSlice = topSlice[5:10]
	}
	fmt.Println(topSlice)

}
