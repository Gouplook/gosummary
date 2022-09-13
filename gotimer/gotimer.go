package main

import (
	"fmt"
	"time"
)

func main() {
	tiker := time.NewTimer(3 * time.Second)

	fmt.Printf("timerType ", tiker)
	fmt.Println(time.Now())
	c := <-tiker.C
	fmt.Println(c)
}
