package goInterface

import "fmt"

// 继承和接口区别
// 1： 继承主要解决代码的复用性
// 2： 接口价值设计好各种规范 让其他自定义类型去实现方法。
// 3： 继承是is-a 关系 接口是like - a 关系

type Ainterface interface {
	Flying()
}
type Binterface interface {
	Swimming()
}

type Mokeny struct {
	Name string
}

type LitterMokeny struct {
	Mokeny
}

func (m *Mokeny) climbing() {
	fmt.Println("climbing.....")
}

func (l *LitterMokeny) Swimming() {
	fmt.Println("swimming...")
}
func (l *LitterMokeny) Flying() {
	fmt.Println("flying...")
	fmt.Println(l.Name)
}

func AAinterface() {
	mokeny := LitterMokeny{
		Mokeny{
			Name: "swk",
		},
	}

	fmt.Println(mokeny.Name)
	mokeny.climbing()
	mokeny.Flying()
	mokeny.Swimming()
}

// 接口
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p *Phone) Start() {
	fmt.Println("Phone 实现Start 接口......")
}

func (p *Phone) Stop() {
	fmt.Println("Phone 实现Stop 接口......")
}

func (p *Phone) Call() {
	fmt.Println("Phone 自定义call 方法")
}

type Camera struct {
	name string
}

func (c *Camera) Start() {
	fmt.Println("Camera 实现Start 接口......")
}

func (c *Camera) Stop() {
	fmt.Println("Camera 实现Stop 接口......")
}

type Computer struct {
}

func (c *Computer) Working(usb Usb) {
	usb.Stop()
	usb.Start()

	if phone, ok := usb.(*Phone); ok {
		phone.Call()
	}
	fmt.Println("working.......")
}

// 类型断言
