//WaitGroup
package main

import (
	"sync"
)

func Fill(array []int, val int) {
	for i := range array {
		array[i] = val
	}
}

func FillConc(array []int, val int, workerCount int) {
	arraySize := len(array)
	step := arraySize / workerCount

	var wg sync.WaitGroup

	wg.Add(workerCount)
	for i := 0; i < workerCount-1; i++ {
		go func(i int) { // attention, il faut capturer le i du for qui peut changer !!!
			Fill(array[i*step:(i+1)*step], val)
			wg.Done()
		}(i)
	}
	go func() {
		Fill(array[(workerCount-1)*step:], val)
		wg.Done()
	}()
	wg.Wait()
}
