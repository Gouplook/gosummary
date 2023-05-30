package main

import (
	"container/list"
	"fmt"
)

//func New() *List
//func (l *List) Back() *Element   // 最后一个元素
//func (l *List) Front() *Element  // 第一个元素
//func (l *List) Init() *List  // 链表初始化
//func (l *List) InsertAfter(v interface{}, mark *Element) *Element // 在某个元素后插入
//func (l *List) InsertBefore(v interface{}, mark *Element) *Element  // 在某个元素前插入
//func (l *List) Len() int // 在链表长度

//func (l *List) MoveAfter(e, mark *Element)  // 把 e 元素移动到 mark 之后
//func (l *List) MoveBefore(e, mark *Element)  // 把 e 元素移动到 mark 之前
//func (l *List) MoveToBack(e *Element) // 把 e 元素移动到队列最后
//func (l *List) MoveToFront(e *Element) // 把 e 元素移动到队列最头部
//func (l *List) PushBack(v interface{}) *Element  // 在队列最后插入元素
//func (l *List) PushBackList(other *List)  // 在队列最后插入接上新队列
//func (l *List) PushFront(v interface{}) *Element  // 在队列头部插入元素
//func (l *List) PushFrontList(other *List) // 在队列头部插入接上新队列
//func (l *List) Remove(e *Element) interface{} // 删除某个元素

func main() {

	l := list.New()
	m := l.PushBack(1)
	s := l.PushBack(2)
	l.InsertBefore(22, s)
	l.InsertAfter(33, m)

	// 遍历链表
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)

	}

}
