package main

import (
	"fmt"
	"time"
)

func thread(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	go thread("t1")
	thread("t2")
}
