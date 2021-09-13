package main

import (
	"fmt"
)

func pair() {
	for i := 1; i < 1001; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	pair()
}
