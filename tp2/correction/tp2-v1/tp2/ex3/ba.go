package main

import (
	"fmt"
	"time"
)

func ba1() {
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i != 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Bonne année ! 🍾")
}

func ba2() {
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i != 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		<-time.After(time.Second) //Receive from ch
	}
	fmt.Println("Bonne année ! 🎉")
}

func ba3() {
	c := time.Tick(time.Second) // attention aux fuites de mémoires !

	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i != 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		<-c
	}
	fmt.Println("Bonne année ! 🥂")
}

func main() {
	// Sleep, After et Tick
	ba1()
	time.Sleep(250 * time.Millisecond)
	ba2()
	time.Sleep(250 * time.Millisecond)
	ba3()
}
