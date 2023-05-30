package main

import (
	"fmt"
	"os"
	"os/signal"
)

func notify() {
	ch := make(chan os.Signal)
	//notify用于监听信号
	//参数1表示接收信号的channel
	//参数2及后面的表示要监听的信号
	//os.Interrupt 表示中断
	//os.Kill 杀死退出进程
	signal.Notify(ch, os.Interrupt, os.Kill)

	//获取信号，如果没有会一直阻塞在这里。
	s := <-ch
	//我们通过Ctrl+C或用taskkill /pid -t -f来杀死进程，查看效果。
	fmt.Println("信号：", s)
}

func main() {
	ch := make(chan os.Signal)
	//如果不指定要监听的信号，那么默认是所有信号
	signal.Notify(ch)

	//停止向ch转发信号，ch将不再收到任何信号
	signal.Stop(ch)
	//ch将一直阻塞在这里，因为它将收不到任何信号
	//所以下面的exit输出也无法执行
	<-ch
	fmt.Println("exit")
}
