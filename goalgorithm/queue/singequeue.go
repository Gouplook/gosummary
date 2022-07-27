package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	MaxSize int
	Array   [5]int
	Front   int // 表示指向队列首
	Rear    int // 表示指向队列的尾部
}

func (q *Queue) AddQueue(val int) (err error) {

	if q.Rear == q.MaxSize-1 {
		fmt.Println("Queue full....")
		return errors.New("queue full...")
	}
	q.Rear++
	q.Array[q.Rear] = val

	return
}

func (q *Queue) GetQueue() (val int, err error) {

	if q.Front == q.Rear {
		return -1, errors.New("queue empty...")
	}
	q.Front++
	val = q.Array[q.Front]
	return val, err
}
func (q *Queue) ShowQueue() {
	fmt.Println("q.Front=", q.Front)
	fmt.Println("q.Rear=", q.Rear)
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, q.Array[i])
	}
	fmt.Println()
}

func main() {
	//先创建一个队列
	queue := &Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
