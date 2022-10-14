package goInterface

// 继承和接口区别
// 1： 继承主要解决代码的复用性
// 2： 接口价值设计好各种规范 让其他自定义类型去实现方法。
// 3： 继承是is-a 关系 接口是like - a 关系

type Person interface {
	GetName() string
	GetAge() int
}

type Student struct {
	Name string
	Age  int
}
type Notity struct {
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}

func (n *Notity) Get() string {
	return "xx"
}


