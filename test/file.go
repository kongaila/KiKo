package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("D:\\data\\qq\\新建文件夹\\02.csv")
	b := bufio.NewReader(file)
	defer file.Close()

	var (
		err error
	)
	for ; err == nil; _, err = b.ReadString('\n') {
		//fmt.Print(line)
	}
	if err == io.EOF {
		//fmt.Print(line)
	} else {
		panic("read occur error!")
	}
	since := time.Since(start)
	fmt.Println("Segment finished in %dms", since) //Segment finished in xxms

}
