package gogroutine

import (
	"fmt"
	"runtime"
	"time"
)

//gogroutine 特点
//1：具有独立栈空间
//2：共享程序堆内存空间
//3：轻量级线程
//4：调度由用户控制

// go 锁的分类
//1: 普通锁：sync.Mutex
//2: 读写互斥锁：sync.RWMutex
//3: 一次性锁 sync.Once
//4: 原子锁 sync/atomic

//MPG

//
func Numcpu() {
	numCpu := runtime.NumCPU
	fmt.Println("CPU num := ", numCpu())
}



func Spinner(delay time.Duration) {

    for {
        for _, r := range `-\|/` {
            fmt.Printf("\r%c", r)
            time.Sleep(delay)
		
        }
    }
}

func Fib(x int) int {
    if x < 2 {
        return x
    }
    return Fib(x-1) + Fib(x-2)
}