package main

import (
	"fmt"
)

type CarNode struct {
	CarId int
	Name  string
	Log   float64 // 经度
	Lat   float64 // 维度
	Next  *CarNode
}

// 添加
func InsertHeroNode(head *CarNode, newNode *CarNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = newNode

}

// 插入带排序的
func InsertSortHeroNode(head *CarNode, newNode *CarNode) {
	temp := head
	flag := true
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.Log >= newNode.Log {
			// 说明newNode应该插入temp后面，不应该是最后一个
			break
		} else if temp.Next.Log == newNode.Log {
			flag = false
			break
		}
		temp = temp.Next
	}
	if !flag {
		fmt.Println("No 已经存在了....")
		return
	} else {
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

// 删除一个节点
func DelCarNode(head *CarNode, carId int) {
	temp := head
	flag := false
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.CarId == carId {
			flag = true
			break
		}
		temp = temp.Next
	}
	if flag {
		temp.Next = temp.Next.Next
	} else {
		fmt.Println("sorry, 要删除的carId不存在")
	}
}
func ShowNode(head *CarNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("kong link..")
		return
	}
	// 遍历
	for {
		fmt.Printf("[%d, %s,%.2f] ==>", temp.Next.CarId, temp.Next.Name, temp.Next.Log)
		temp = temp.Next
		// 退出条件
		if temp.Next == nil {
			break
		}
	}
}

func main() {

	head := &CarNode{}
	car1 := &CarNode{
		CarId: 1,
		Name:  "AGV",
		Log:   55.3,
		Lat:   22.6,
	}
	car2 := &CarNode{
		CarId: 2,
		Name:  "AGV",
		Log:   59.3,
		Lat:   25.6,
	}
	car3 := &CarNode{
		CarId: 3,
		Name:  "AGV",
		Log:   52.2,
		Lat:   33.6,
	}
	//InsertHeroNode(head, car3)
	//InsertHeroNode(head, car2)
	//InsertHeroNode(head, car1)

	InsertSortHeroNode(head, car3)
	InsertSortHeroNode(head, car2)
	InsertSortHeroNode(head, car1)
	ShowNode(head)
}
