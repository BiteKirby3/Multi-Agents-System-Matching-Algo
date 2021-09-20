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
		fmt.Println(msg)
		fmt.Println(i)
	}
}

func compteMsgFromTo(start int, end int, msg string) {
	for i := start; i < end; i++ {
		fmt.Println(msg)
		fmt.Println(i)
	}
}

func main() {
	//compteMsgFromTo(1, 10, "msg")
	//go compte(100)
	//go compteMsg(100, "routine1")
	//time.Sleep(time.Second * 10)
	//go compteMsg(100, "routine2")
	//time.Sleep(time.Second)
	//fmt.Scanln()

	for g := 0; g < 10; g++ {
		go compteMsgFromTo(g*10, g*10+10, "routine "+strconv.Itoa(g))
		time.Sleep(time.Second)
	}
	fmt.Scanln()
}
