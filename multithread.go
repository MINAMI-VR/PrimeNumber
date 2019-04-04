package main

import (
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
)

const bufSize = 16000

const offset = 3

func main() {
	threads := getThreads()
	n := uint64(2*threads + offset)
	wg := &sync.WaitGroup{}
	for i := uint64(offset); i < n; i += 2 {
		if gcd(i, uint64(threads)) == 1 {
			wg.Add(1)
			go prime(i, threads, wg)
		} else if isPrime(i) {
			os.Stdout.Write(append([]byte(strconv.FormatUint(i, 10)), '\n'))
		}
	}
	wg.Wait()
}

func getThreads() int {
	numCPU := runtime.NumCPU()
	minThread := 2
	for {
		if minThread >= numCPU {
			break
		}
		minThread = minThread * 2
	}
	maxThreads := minThread * 2
	validThreads := minThread
	threads := minThread
	for i := minThread; i < maxThreads; i++ {
		count := 0
		n := i*2 + 3
		for j := 3; j < n; j += 2 {
			if gcd(uint64(j), uint64(i)) == 1 {
				count++
			}
		}
		if float64(count)/float64(i) < float64(validThreads)/float64(threads) && validThreads <= i {
			threads = i
			validThreads = count
		}
	}
	return threads
}

func prime(n uint64, thread int, wg *sync.WaitGroup) {
	increment := uint64(2 * thread)
	max := uint64(math.MaxUint64 - increment)
	buf := make([]byte, 0, bufSize)
	for i := n; i < max; i += increment {
		if isPrime(i) {
			buf = append(append(buf, strconv.FormatUint(i, 10)...), '\n')
			if len(buf) > bufSize-10 {
				os.Stdout.Write(buf)
				buf = make([]byte, 0, bufSize)
			}
		}
	}
	wg.Done()
}

func isPrime(i uint64) bool {
	n := uint64(math.Sqrt(float64(i)))
	for j := uint64(offset); j <= n; j += uint64(2) {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b uint64) uint64 {
	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	r := a % b
	for {
		if r == 0 {
			break
		}
		a = b
		b = r
		r = a % b
	}
	return b
}
