package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	thread := 15
	wg := &sync.WaitGroup{}
	for i := uint64(3); i < uint64(2*thread+3); i += 2 {
		go prime(i, thread, wg)
	}
	wg.Wait()
}

func prime(n uint64, thread int, wg *sync.WaitGroup) {
	wg.Add(1)
	increment := uint64(2 * thread)
	max := math.MaxUint64 - increment
	for i := n; i < max; i += increment {
		isPrime := true
		for j := uint64(3); j*j <= i; j += uint64(2) {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
	wg.Done()
}
