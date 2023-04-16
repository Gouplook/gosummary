/**
 * @Author: yinjinlin
 * @File:  function
 * @Description:
 * @Date: 2021/5/27 下午3:52
 */

package gofunc

import (
	"fmt"
	"strings"
)

// 函数 + 引用外部变量 = 闭包

// 函数闭包
func Square(x int) int {
	return x * x
}

func GenPalyer(name string) func() (string, int) {
	// 定义玩家血量
	hp := 1000
	return func() (string, int) {
		return name, hp
	}
}

// 斐波那契数列
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg ,
//   如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。
func MakeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 匿名函数
func Anonymous(a int, b int) {

	res := func(a int, b int) int {
		return a + b
	}
	rs := res(a, b)

	fmt.Println(rs)
}

//捕捉异常
func Recvertest() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	sum := num1 / num2

	fmt.Println(sum)
}

func Lat(x, y float64) (float64, float64) {
	// 1点:[113.9275258976705, 22.31306742552032, 0.0]
	//2点: [113.92756334083255, 22.31296897641068, 0.0]
	//3点:[113.92992955993527, 22.313852420058637, 0.0]
	//4点:[113.92982479694706, 22.31368647761813, 0.0]
	//5点:[113.92930786718041, 22.313556962702336, 0.0]
	//6点:[113.92928931487226, 22.313608083113454, 0.0]
	//7点:[113.92879828808876, 22.313483836816218, 0.0]
	//8点:[113.92850531058023, 22.313264729994497, 0.0]

	//1. -97.064, 6.544
	//2. -96.489, -4.422
	//3. 96.558, 11.924
	//4. 81.283, -1.011
	//5. 27.502, 2.432
	//6. 95.339, 8.062
	//7. 41.513, 11.196
	//8. 5.56, -1.8

	// Lat = ax+cy+e
	// lng = bx+dy+f
	// 33to39
	a := 0.000002688630
	b := 0.000009287578134
	c := 0.000009421554452
	d := -0.000003150106852
	e := 22.31326674
	f := 113.928448

	// 3to2a 74.432, 11.279
	//3to2b 73.940, -0.747

	// 29to32
	//a := 0.0000030024966
	//b := 0.0000094182350
	//c := 0.0000092832860
	//d := -0.0000030228490
	//e := 22.3134518110892
	//f := 113.9290561984490

	// --> gps
	lat := a*x + c*y + e
	lng := b*x + d*y + f

	// 地图部门给的点（景威）
	//8点:[113.92850531058023, 22.313264729994497, 0.0]
	// 根据公式计算的点
	//Lng= 113.92850530912676, 22.313264729984787

	// 机器采集的节点
	//8. 5.56, -1.8

	return lat, lng
}

func Lat2(x, y float64) (float64, float64) {

	// 29to32
	a := 0.0000030024966
	b := 0.0000094182350
	c := 0.0000092832860
	d := -0.0000030228490
	e := 22.3134518110892
	f := 113.9290561984490

	// --> gps
	lat := a*x + c*y + e
	lng := b*x + d*y + f

	return lat, lng
}
