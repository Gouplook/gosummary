package main

import (
	"fmt"
	"regexp"
	"time"
)

func Mail163() {

}

func emailVerification() {
	email := "#yinjinlinuplook@163.com" // 要验证的邮箱地址
	// 正则表达式模式用于验证邮箱地址的格式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// 使用正则表达式进行匹配
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		fmt.Println("验证出错:", err)
		return
	}

	if match {
		fmt.Println("邮箱地址格式正确")
	} else {
		fmt.Println("邮箱地址格式不正确")
	}
}

func main() {
	// SXCTJZVLQAYOPJWM
	// port := 465
	// host := "smtp.163.com"
	// m := gomail.NewMessage()
	// m.SetHeader("From", m.FormatAddress("yinjinlin_uplook@163.com", "yinjinlin"))
	// m.SetHeader("To", "jinlin.yin@uisee.com")
	// m.SetHeader("Subject", "TH1")
	// m.SetBody("text/html", "hello wold")
	// d := gomail.NewDialer(host, port, "yinjinlin_uplook@163.com", "5Yue23ri")
	// err := d.DialAndSend(m)
	// if err != nil {
	//	fmt.Println(err)
	//}

	//emailVerification()
	//str := "大象二期 | 内外饰-排气 | rc29(总报警: 8处,本次报警共: 1处)"
	//
	//// 查找"总报警: "的位置 13
	//index := strings.Index(str, "总报警: ")
	//if index >= 0 {
	//	// 截取数字部分
	//	numberStr := str[index+len("总报警: "):]
	//	// 查找数字的结束位置
	//	endIndex := strings.IndexFunc(numberStr, func(r rune) bool {
	//		return r < '0' || r > '9'
	//	})
	//	if endIndex > 0 {
	//		number, err := strconv.Atoi(numberStr[:endIndex])
	//		if err == nil {
	//			fmt.Println(number) // 输出结果为: 8
	//		}
	//	}
	//}

	// 获取当前时间
	//now := time.Now()	now2 :=  time.Now().Unix()
	//
	//// 计算最近30天的时间戳
	//latest30Days := now.AddDate(0, 0, -30).Unix()
	//latest30Days2 := now2 - (30 * 24 * 60 * 60)
	//// 计算最近7天的时间戳
	//latest7Days := now.AddDate(0, 0, -7).Unix()
	//fmt.Printf("%v\n", now)
	//fmt.Printf("最近30天的时间戳：%v\n", latest30Days)
	//fmt.Printf("最近7天的时间戳：%v\n", latest7Days)
	////// 获取当前时间戳
	//now := time.Now().Unix()
	//
	//// 计算最近30天的时间戳
	//latest30Days := now - (30 * 24 * 60 * 60)
	//
	//// 计算最近7天的时间戳
	//latest7Days := now - (7 * 24 * 60 * 60)
	//
	//fmt.Println()
	//fmt.Printf("最近30天的时间戳：%v\n", latest30Days)
	//fmt.Printf("最近7天的时间戳：%v\n", latest7Days)

	ticker := time.NewTicker(10 * time.Second) // 定义一个每10秒执行一次的计时器
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				// 在这里实现你的定时任务逻辑
				fmt.Println("Tick at ", t)
			}
		}
	}()

	time.Sleep(60 * time.Second) // 等待60秒钟

	done <- true // 停止计时器
	ticker.Stop()
	fmt.Println("Ticker stopped.")
}
