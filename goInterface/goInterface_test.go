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
