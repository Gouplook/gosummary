package main

import (
	"fmt"
	"time"
)

// 一次性定时器
func newTimer() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("Timer expired!")
}

// 周期行定时器
func newTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
func main() {
	//
	//newTimer()

	newTicker()
}
