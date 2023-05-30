package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// once
func onceFunc() {
	var once sync.Once
	oneBody := func() {
		fmt.Println("oneBody...")
	}
	twoBody := func() {
		fmt.Println("twoBody...")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(oneBody)
			once.Do(twoBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

}

// pool
var createNum int32

func createBuffer() interface{} {
	atomic.AddInt32(&createNum, 1)
	buffer := make([]byte, 1024)
	return buffer
}

func poolFunc() {
	bufferPool := &sync.Pool{New: createBuffer}

	workerPool := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerPool)

	for i := 0; i < workerPool; i++ {
		go func() {
			defer wg.Done()
			buffer := bufferPool.Get()
			_ = buffer.([]byte)
			defer bufferPool.Put(buffer)
			//buffer := createBuffer()
			//_ = buffer.([]byte)
		}()
	}
	wg.Wait()
	fmt.Printf(" %d buffer objects were create.\n", createNum)
	time.Sleep(5 * time.Second)

}
func main() {

}
