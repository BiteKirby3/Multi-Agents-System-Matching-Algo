package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Fill(sl []int) {
	for i := range sl {
		sl[i] = rand.Intn(100)
	}
	//fmt.Println(sl)
}
func Moyenne(sl []int) (moy float64) {
	sum := 0.0
	if len(sl) == 0 {
		moy = 0
		return
	} else {
		for i := range sl {
			sum += float64(sl[i])
		}
		moy = sum / float64(len(sl))
		return
	}
}

func ValeurCentrales(sl []int) (moy float64, med int) {
	moy = Moyenne(sl)
	//médiane
	sort.Ints(sl)
	if len(sl)%2 == 0 {
		med = (sl[len(sl)/2] + sl[len(sl)/2-1]) / 2
	} else {
		med = sl[((len(sl)+1)/2)-1]
	}
	return
}

func Plus1(sl []int) {
	for i := range sl {
		sl[i] = sl[i] + 1
	}
}

func Compte(n int, tab []int) {
	cpt := 0
	for i := range tab {
		if tab[i] == n {
			cpt++
		}
	}
	fmt.Println(cpt)
}

func main() {
	slice := make([]int, 3)
	rand.Seed(time.Now().Unix())
	Fill(slice)
	fmt.Println(slice)
	fmt.Println(ValeurCentrales(slice))
	Compte(55, slice)
	Plus1(slice)
	fmt.Println(slice)
	fmt.Println("Moyenne : ", Moyenne(slice))
}
