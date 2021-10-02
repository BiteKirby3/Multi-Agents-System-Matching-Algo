package main

import (
	"sync"
)

func ForEach(array []int, f func(int) int) {
	for i, v := range array {
		array[i] = f(v)
	}
}

func ForEachConc(array []int, f func(int) int, workerCount int) {
	arraySize := len(array)
	step := arraySize / workerCount

	var wg sync.WaitGroup

	wg.Add(workerCount)
	for i := 0; i < workerCount-1; i++ {
		go func(i int) {
			ForEach(array[i*step:(i+1)*step], f)
			wg.Done()
		}(i)
	}
	go func() {
		ForEach(array[(workerCount-1)*step:], f)
		wg.Done()
	}()
	wg.Wait()
}
