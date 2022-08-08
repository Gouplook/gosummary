/**
 * @Author: yinjinlin
 * @File:  goslice_test
 * @Description:
 * @Date: 2021/12/7 下午1:45
 */

package goslice

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestIntSort(t *testing.T) {
	// IntSort()

	fmt.Println(7000 / 22.5 / 8 * 3)
	// fmt.Println(7000/22.5/)
}
func TestSliceSort(t *testing.T) {
	SliceSort()
}

func TestSliceContains(t *testing.T) {
	a := []int{0, 1}

	b := []int{1, 0}

	//a = tool.ArrayUniqueInt(a)
	SliceContains(a, b)
}

func Test_BB(t *testing.T) {

	m := []int{1, 33, 66, 12}
	ms := make([]string, 0)
	for _, v := range m {
		str := strconv.Itoa(v)
		ms = append(ms, str)
	}
	fmt.Println(ms)
	fmt.Println(strings.Join(ms, ","))
}

func TestStringMode(t *testing.T) {
	StringMode()
}
