/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/28 09:30
@Description:

*********************************************/
package gostruct

import (
	"fmt"
	"testing"
)

func TestStructBase(t *testing.T) {
	SStructBase()
}

//// 结构体与切片
func TestStructAndSlice(t *testing.T) {
	sStructAndSlice()
}

// 继承

func TestStructAndMap(t *testing.T) {
	//sStructAndSlice()
	StructAndMap()
}

func TestMaptoMap(t *testing.T) {
	MaptoMap()
}

func TestSet(t *testing.T) {
	//SetStruct()
	s := 472446402560
	fmt.Println(s >> 32)
	fmt.Println(450971566080 >> 32)
	fmt.Println(472446402560 >> 32)
}
