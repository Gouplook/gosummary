/**
 * @Author: yinjinlin
 * @File:  avdMap_test
 * @Description:
 * @Date: 2021/6/17 上午10:16
 */

package gomap

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"testing"
)

// map和map切片测试
func TestAdvanceMap(t *testing.T) {
	AdvanceMap()
}

func TestSileIn(t *testing.T) {
	SileIn()
}

// 双map测试
func TestAdvMapMap(t *testing.T) {
	AdvMapMap()
}

func TestSlieIn2(t *testing.T) {
	SlieIn2()
}

func TestMapMap(t *testing.T) {
	//MapMap2()
	MapMap()
}

func TestMapSplitToStruct2(t *testing.T) {
	// MapSplitToStruct2(CardIcad{})
	fmt.Println()
}

type ABase struct {
	RR string
	B  []BBase
}
type BBase struct {
	Name string
	Age  int
	C    []CBase
}

type CBase struct {
	CC string
}

func TestName1(t *testing.T) {

	bb := make([]map[string]interface{}, 0)
	bb = append(bb, map[string]interface{}{
		"Name": "xxx",
		"Age":  10,
		"C":    []map[string]interface{}{{"CC": "xxx"}, {"CC": "yyy"}},
	})
	a := ABase{}
	mapstructure.WeakDecode(bb, &a.B)
	bytes, _ := json.Marshal(a)
	t.Log(string(bytes))

}

func TestMapToStruct(t *testing.T) {
	MapToStruct()
}

// 正向转换
func TestMapToStruct2(t *testing.T) {
	MapToStruct1()
}

// 反向转换
func TestStructToMap(t *testing.T) {
	StructToMap()
}

func TestDoubleMap(t *testing.T) {
	DoubleMap()
}

func TestMapSlice(t *testing.T) {
	MapSlice()
}

func TestMapSliceStruct(t *testing.T) {
	MapSliceStruct()
}
