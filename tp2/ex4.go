package main

import (
	"fmt"
	"time"
)

// remplit tab avec la valeur v
func Fill(tab []int, v int) {
	for i := range tab {
		tab[i] = v
	}
}

func FillGo(tab []int, v int) {
	c1 := make(chan bool)
	c2 := make(chan bool)
	c3 := make(chan bool)
	t := len(tab)
	slice1 := tab[0:(t / 3)]
	slice2 := tab[(t / 3) : (t/3)*2]
	slice3 := tab[(t/3)*2:]
	go fillSeq(slice1, v, c1)
	go fillSeq(slice2, v, c2)
	go fillSeq(slice3, v, c3)
	<-c1
	<-c2
	<-c3
}

func fillSeq(tab []int, val int, done chan bool) {
	for i := 0; i < len(tab); i++ {
		tab[i] = val
	}
	done <- true
}

func equalSeq(tab1 []int, tab2 []int, done chan bool) bool {
	if len(tab1) != len(tab2) {
		done <- false
		return false
	}
	for i := 0; i < len(tab1); i++ {
		if tab1[i] != tab2[i] {
			done <- false
			return false
		}
	}
	done <- true
	return true
}

func equalGo(tab1 []int, tab2 []int, done chan bool) bool {
	var c1Bool bool
	var c2Bool bool
	var c3Bool bool
	c1 := make(chan bool)
	c2 := make(chan bool)
	c3 := make(chan bool)
	t1 := len(tab1)
	t2 := len(tab2)
	slice1 := tab1[0:(t1 / 3)]
	slice2 := tab1[(t1 / 3) : (t1/3)*2]
	slice3 := tab1[(t1/3)*2:]
	slice11 := tab2[0:(t2 / 3)]
	slice22 := tab2[(t2 / 3) : (t2/3)*2]
	slice33 := tab2[(t2/3)*2:]
	go equalSeq(slice1, slice11, c1)
	go equalSeq(slice2, slice22, c2)
	go equalSeq(slice3, slice33, c3)
	c1Bool = <-c1
	c2Bool = <-c2
	c3Bool = <-c3
	return c1Bool && c2Bool && c3Bool
}

func ForEach(tab []int, f func(int) int) {
	for i := range tab {
		tab[i] = f(tab[i])
	}
}

// copy le tableau src dans dest
func Copy(src []int, dest []int) {
	for i := range src {
		dest[i] = src[i]
	}
}

// vérifie que tab1 et tab2 sont identiques
func Equal(tab1 []int, tab2 []int) {
	idem := true
	if len(tab1) != len(tab2) {
		idem = false
	} else {
		for i := range tab1 {
			if tab1[i] != tab2[i] {
				idem = false
				break
			}
		}
	}
	if idem {
		fmt.Println("idem")
	} else {
		fmt.Println("différent")
	}
}

func main() {
	//507.9µs Fill avec Go
	/*
		start := time.Now()
		var tab [9000000]int
		slice := tab[:]
		FillGo(slice, 5)
		elapsed := time.Since(start)
		fmt.Printf("page took %s", elapsed)
	*/
	/*18.9604ms Fill seq
	start := time.Now()
	var tab [9000000]int
	slice := tab[:]
	Fill(slice, 5)
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
	*/

	/*Equal 7.289ms
	var tab1 [9000000]int
	slice1 := tab1[:]
	var tab2 [9000000]int
	slice2 := tab2[:]
	FillGo(slice1, 5)
	FillGo(slice2, 5)
	start := time.Now()
	Equal(slice1, slice2)
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
	*/
	//EqualGo 4.2085ms
	var tab1 [9000000]int
	slice1 := tab1[:]
	var tab2 [9000000]int
	slice2 := tab2[:]
	FillGo(slice1, 5)
	FillGo(slice2, 5)
	start := time.Now()
	ch := make(chan bool)
	equalGo(slice1, slice2, ch)
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
}
