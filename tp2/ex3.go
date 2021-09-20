package main

import (
	"fmt"
	"sync"
)

var n = 0
var l sync.Mutex

func afficher(i int) {
	fmt.Print(i)
	fmt.Print(", ")
}

func main() {
	/* Sleep
	for i := 5; i > 0; i-- {
		go afficher(i)
		time.Sleep(time.Second)
	}

	fmt.Print("Bonne Année !")
	fmt.Scanln()
	*/

	/*	After
		for i := 5; i > 0; i-- {
			afficher(i)
			<-time.After(time.Second)
		}
		fmt.Println("Bonne Année !")
	*/

	/* Tick
	c := time.Tick(time.Second) //attention aux fuites de mémoires
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i != 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("...")
		}
		<-c
	}
	fmt.Println("Bonne Année !")
	*/
}
