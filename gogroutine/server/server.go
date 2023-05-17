package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// 向c中写入数据
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	listener, err := net.Listen("tcp", "localhost:8009")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("localhost:8009....")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// 支持多个客户端同时链接。
		go handleConn(conn) // handle one connection at a time
	}
}
