package main

import (
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

const bufSize = 16000

const offset uint64 = 3

const threads uint64 = 45

const two uint64 = 2

const max uint64 = math.MaxUint64 - 90

const maxUint64 uint64 = math.MaxUint64

func main() {
	n := uint64(93)
	wg := &sync.WaitGroup{}
	buf := make([]byte, 1, 1000)
	buf[0] = '2'
	for i := offset; i < n; i += 2 {
		if gcd(i, threads) == 1 {
			wg.Add(1)
			go prime(i, wg)
		} else if isPrime(i) {
			buf = append(append(buf, strconv.FormatUint(i, 10)...), '\n')
		}
	}
	os.Stdout.Write(buf)
	go func() {
		buf := make([]byte, 0, 1000)
		for i := max; i < maxUint64; i += 2 {
			if isPrimeSleep(i) {
				buf = append(append(buf, strconv.FormatUint(i, 10)...), '\n')
			}
		}
		if isPrimeSleep(maxUint64) {
			buf = append(append(buf, strconv.FormatUint(maxUint64, 10)...), '\n')
		}
		os.Stdout.Write(buf)
	}()
	wg.Wait()
}

func prime(n uint64, wg *sync.WaitGroup) {
	buf := make([]byte, 0, bufSize)
	for i := n; i < max; i += 90 {
		if isPrime(i) {
			buf = append(append(buf, strconv.FormatUint(i, 10)...), '\n')
			if len(buf) > 1590 {
				os.Stdout.Write(buf)
				buf = make([]byte, 0, bufSize)
			}
		}
	}
	os.Stdout.Write(buf)
	wg.Done()
}

func isPrimeSleep(i uint64) bool {
	for j := offset; j*j <= i; j += two {
		if i%j == 0 {
			return false
		}
		time.Sleep(time.Nanosecond)
	}
	return true
}

func isPrime(i uint64) bool {
	for j := offset; j*j <= i; j += two {
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
