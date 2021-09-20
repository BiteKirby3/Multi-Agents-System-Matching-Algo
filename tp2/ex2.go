package main

import (
	"fmt"
	"sync"
)

var n = 0
var l sync.Mutex

func f() {
	l.Lock()
	n++
	l.Unlock()
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
}
