package gogroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNumcpu(t *testing.T) {
	Numcpu()
}

var (
	mapDate = make(map[int]int, 10)
	lock    sync.Mutex
)

func TestGgogroutine(t *testing.T) {

	go Spinner(100 * time.Millisecond)
	const n = 25
	fibN := Fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}
