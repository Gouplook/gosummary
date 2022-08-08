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

	computer.Working(phone)
	computer.Working(&camera)

	var u Usb = phone
	var c Usb = &camera
	u.Stop()
	c.Start()

	// 多态数组
	var useArr [3]Usb
	useArr[0] = &Phone{"vivo"}
	useArr[1] = &Phone{"hawei"}
	useArr[2] = &Camera{"vivo"}

	fmt.Println(useArr[0])

}
