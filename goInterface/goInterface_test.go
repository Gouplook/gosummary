package goInterface

import (
	"fmt"
	"testing"
)

func TestAAinterface(t *testing.T) {
	AAinterface()
}

// 抽象接口，对于不可导出的结构体，可以利用接口导出
func TestPerson_Greet(t *testing.T) {
	p := NewPerson("yin", 32)
	fmt.Println(p)
	// l
	p.Greet()

}

func TestWorking(t *testing.T) {
	phone := new(Phone)
	camera := Camera{}
	computer := new(Computer)

	fmt.Println("================")
	phone.Start()

	fmt.Println("================")
	computer.Working(phone)
	computer.Working(&camera)

	// 接口，传入什么，就适配什么。
	// 接口定义好了，任何自定义类型都可以实现接口的方法。
	var u Usb = phone
	var c Usb = &camera
	u.Stop()
	c.Start()

	// 多态数组
	var useArr [3]Usb
	useArr[0] = &Phone{"vivo"}
	useArr[1] = &Phone{"hawei"}
	useArr[2] = &Camera{"vivo"}

	for k, v := range useArr {
		fmt.Println(k)
		fmt.Println(v)
	}

}
