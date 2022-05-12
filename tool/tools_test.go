package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrExplode2IntArr(t *testing.T) {
	str := "112,34,56,78* 88  990"
	rs := StrExplode2IntArr(str, ",")
	fmt.Println(rs)
}

func TestStringsTrim(t *testing.T) {
	var s = "Hello,goodbye, etc!"
	str2 := strings.TrimRight(s, ",")
	fmt.Println("str2 = ", str2)
	fmt.Println("====:", strings.TrimRight("abba", "ba"))

	str := StringsTrim(s, ", etc")
	fmt.Println("str = ", str)
	//s = strings.TrimSuffix(s, "goodbye, etc!")
	s1 := strings.TrimSuffix(s, "planet") // 没有就返回原字符串
	fmt.Println(s)
	fmt.Println("s1= ", s1)
}

// 生产字符串编号
func TestGenRandom(t *testing.T) {
	s := GenRandom(11, 123)
	fmt.Println("ss", s)
}
