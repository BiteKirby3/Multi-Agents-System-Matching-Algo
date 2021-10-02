package main

import (
	"fmt"
	"strconv"
	"time"
)

func compte(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func compteMsg(n int, msg string) {
	for i := 0; i < n; i++ {
		fmt.Println(msg, i)
	}
}

func compteMsgFromTo(start int, end int, msg string) {
	for i := start; i < end; i++ {
		fmt.Println(msg, i)
	}
}

func main() {
	// Question 1
	compte(10)
	go compte(10)
	time.Sleep(250 * time.Millisecond)

	// Question 2
	compteMsg(10, "sans goroutine:")
	go compteMsg(10, "goroutine n°1:")
	go compteMsg(10, "goroutine n°2:")
	time.Sleep(250 * time.Millisecond)

	// Question 3-4
	for i := 0; i < 10; i++ {
		go func(i int) {
			compteMsgFromTo(i*10, (i+1)*10, "routine n°"+strconv.Itoa(i))
		}(i)
	}
	time.Sleep(250 * time.Millisecond)
}
