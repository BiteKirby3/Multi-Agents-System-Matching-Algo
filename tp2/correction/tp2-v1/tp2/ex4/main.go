package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func Equal(array1 []int, array2 []int) {
	if len(array1) != len(array2) {
		panic("diff arrays")
	}

	for i := range array1 {
		if array1[i] != array2[i] {
			panic("diff arrays")
		}
	}
}

func main() {
	const size = 128 * (1 << 20)

	var array1 [size]int
	var array2 [size]int

	fmt.Println("#proc:", runtime.NumCPU())
	workerCount := runtime.NumCPU() / 2

	t := time.Now()
	Equal(array1[:], array2[:])
	fmt.Println(time.Since(t))

	fmt.Println("*** FILL ***")
	t = time.Now()
	Fill(array1[:], 12)
	fmt.Println("seq :", time.Since(t))

	t = time.Now()
	FillConc(array2[:], 12, workerCount)
	fmt.Println("conc:", time.Since(t))

	Equal(array1[:], array2[:])

	fmt.Println("*** FOREACH +1 ***")

	square := func(i int) int {
		return i * i
	}

	t = time.Now()
	ForEach(array1[:], square)
	fmt.Println("seq :", time.Since(t))

	t = time.Now()
	ForEachConc(array2[:], square, workerCount)
	fmt.Println("conc:", time.Since(t))

	Equal(array1[:], array2[:])

	fmt.Println("*** FOREACH RAND ***")

	// petite coquetterie pour créer une fonction qui renvoie un nombre aléatoire entre 0 et 10...
	frand := func(n int) func(int) int {
		return func(int) int {
			return rand.Intn(n)
		}
	}(10)

	rand.Seed(12)
	t = time.Now()
	ForEach(array1[:], frand)
	fmt.Println("seq :", time.Since(t))

	rand.Seed(12)
	t = time.Now()
	ForEachConc(array2[:], frand, workerCount)
	fmt.Println("conc:", time.Since(t))
	// rq: rand fait un mutex, pas de gain de temps !
	// source: https://cs.opensource.google/go/go/+/refs/tags/go1.17.1:src/math/rand/rand.go;l=125;drc=refs%2Ftags%2Fgo1.17.1;bpv=1;bpt=1

	// Equal(array1[:], array2[:]) // n'a pas de sens avec les rand

	fmt.Println("*** FOREACH 10^n ***")

	f3 := func(i int) int {
		return int(math.Pow10(i))
	}

	t = time.Now()
	ForEach(array1[:], f3)
	fmt.Println("seq :", time.Since(t))

	t = time.Now()
	ForEachConc(array2[:], f3, workerCount)
	fmt.Println("conc:", time.Since(t))
}
