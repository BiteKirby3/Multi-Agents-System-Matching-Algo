package main

import (
	"fmt"
	"sync"
)

type synchronzedInt struct {
	sync.Mutex
	i int
}

var n synchronzedInt

func f() {
	n.Lock()
	defer n.Unlock()
	n.i++
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n.i)
}
