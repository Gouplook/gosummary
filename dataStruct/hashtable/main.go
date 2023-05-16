package main

import (
	"fmt"
)

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}
type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   // 这是辅助指针
	var pre *Emp = nil // 这是一个辅助指针 pre 在cur前面
	// 如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp //完成
		return
	}
	//如果不是一个空链表,给emp找到对应的位置并插入
	//思路是 让 cur 和 emp 比较，然后让pre 保持在 cur 前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保证同步,置换
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur

}

//显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}

	//变量当前的链表，并显示数据
	cur := this.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}

//根据id查找对应的雇员，如果没有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//定义hashtable ,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

type HaseTable2 struct {
	// key 是
	LinkArr []EmpLink
}

//给HashTable 编写Insert 雇员的方法.
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp) //
}

//编写方法，显示hashtable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)

	}
}

//编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对于的链表的下标
}

//编写一个方法，完成查找
func (this *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定将该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}
