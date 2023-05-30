package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func Mail163() {
	
}

func main() {
	// SXCTJZVLQAYOPJWM
	port := 465
	host := "smtp.163.com"
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("yinjinlin_uplook@163.com", "yinjinlin"))
	m.SetHeader("To", "jinlin.yin@uisee.com")
	m.SetHeader("Subject", "TH1")
	m.SetBody("text/html", "hello wold")
	d := gomail.NewDialer(host, port, "yinjinlin_uplook@163.com", "5Yue23ri")
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
}
