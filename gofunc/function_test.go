/**
 * @Author: yinjinlin
 * @File:  function_test
 * @Description:
 * @Date: 2021/5/27 下午3:52
 */

package gofunc

import (
	"fmt"
	"testing"
)

func TestGenPalyer(t *testing.T) {
	// 创建一个玩家生成器
	generator := GenPalyer("mosou")

	// 返回新创建玩家的姓名, 血量
	name, hp := generator()

	fmt.Println(name, hp)

}

// 斐波那契数列
func TestFib(t *testing.T) {
	f00 := Fib()
	fmt.Println(f00(), f00(), f00(), f00())
}

// 判断图片文件后缀,若没有加上jpg后缀(使用闭包）
func TestMakeSuffix(t *testing.T) {
	ff := MakeSuffix(".jpg")

	fmt.Println(ff("win"))
	fmt.Println(ff("jack.jpg"))

}

// 匿名函数
func TestAnonymous(t *testing.T) {
	Anonymous(20, 60)
}

func TestRecvertest(t *testing.T) {
	Recvertest()
}

func TestLat(t *testing.T) {
	// 1点:[113.9275258976705, 22.31306742552032, 0.0]
	//2点: [113.92756334083255, 22.31296897641068, 0.0]
	//3点:[113.92992955993527, 22.313852420058637, 0.0]
	//4点:[113.92982479694706, 22.31368647761813, 0.0]

	//5点:[113.92930786718041, 22.313556962702336, 0.0]
	//Lng= 113.9293078671792, lat=22.31353265864581
	//lat= 22.313556962702247
	//Lng= 113.9293078671792

	//6点:[113.92928931487226, 22.313608083113454, 0.0]
	//7点:[113.92879828808876, 22.313483836816218, 0.0]
	//re:  113.9288688238274,  22.313483836820833
	// 地图
	//8点:[113.92850531058023, 22.313264729994497, 0.0]
	// 根据公式计算的点
	//Lng= 113.92850530912676, 22.313264729984787

	// 机器采集的节点
	//8. 5.56, -1.8

	//lat, lnt := Lat(5.56, -1.8)

	// 30b
	//lat, lnt := Lat(79.088, 8.17)
	//lat= 22.313776640402182
	//Lng= 113.92977214309296

	//g33 9.388, -3.436
	// 8.736, -0.720
	//lat, lnt := Lat(8.736, -0.720)
	//lat= 22.313259608397342
	//Lng= 113.92854601555067

	//lat= 22.313283444352475
	//Lng= 113.92853140435952

	// T16FG034.L 6.115, 8.237
	//T16FG034.R 20.303, 8.195
	// 3to2a 74.432, 11.279
	//3to2b 73.940, -0.747
	//2to3b 11.056, -0.445
	//2to3a 10.995, 11.774

	//--
	// TE_3TO2A   74.154,8.572
	//TE_3TO2B    74.024,1.461
	//TE_2TO3A    10.995,11.774
	//TE_2TO3B  11.056,-0.445

	// --- 33to36 ----
	//T16FXG036.1 -104.049, 2.822
	lat, lnt := Lat(-104.049, 2.822)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FXG036.2 -101.900,2.913
	lat, lnt = Lat(-101.900, 2.913)
	fmt.Printf("T16FXG036.2,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FXG036.3 -99.878, 2.97
	lat, lnt = Lat(-99.878, 2.97)
	fmt.Printf("T16FXG036.3,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG036L -78.255, 5.33
	lat, lnt = Lat(-78.255, 5.33)
	fmt.Printf("T16FG036L,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG035.L -78.280, -1.342
	lat, lnt = Lat(-78.280, -1.342)
	fmt.Printf("T16FG035.L,Lng =%f,Lat =%f\n", lnt, lat)
	//TE_T16FG036 -60.564, 9.288
	lat, lnt = Lat(-60.564, 9.288)
	fmt.Printf("TE_T16FG036,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG036.A -60.534, 7.368
	lat, lnt = Lat(-60.534, 7.368)
	fmt.Printf("T16FG036.A,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG036.R -53.504, 6.372
	lat, lnt = Lat(-53.504, 6.372)
	fmt.Printf("T16FG036.R,Lng =%f,Lat =%f\n", lnt, lat)
	//TE_T16FG035 -57.145, -3.622
	lat, lnt = Lat(-57.145, -3.622)
	fmt.Printf("TE_T16FG035,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG035.R -51.662, -0.528
	lat, lnt = Lat(-51.662, -0.528)
	fmt.Printf("T16FG035.R,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG035.A -57.148, -1.647
	lat, lnt = Lat(-57.148, -1.647)
	fmt.Printf("T16FG035.A,Lng =%f,Lat =%f\n", lnt, lat)
	//TE_T16FG033 9.388, -3.436
	lat, lnt = Lat(9.388, -3.436)
	fmt.Printf("TE_T16FG033,Lng =%f,Lat =%f\n", lnt, lat)

	//T16FG033.A 11.305, -0.381
	lat, lnt = Lat(11.305, -0.381)
	fmt.Printf("T16FG033.A,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG033.R 19.504, 1.047
	lat, lnt = Lat(19.504, 1.047)
	fmt.Printf("T16FG033.R,Lng =%f,Lat =%f\n", lnt, lat)

	//T16FG033.B 14.913, 4.802
	lat, lnt = Lat(14.913, 4.802)
	fmt.Printf("T16FG033.B,Lng =%f,Lat =%f\n", lnt, lat)

	//T16FG034.B 10.783, 4.634
	lat, lnt = Lat(10.783, 4.634)
	fmt.Printf("T16FG034.B,Lng =%f,Lat =%f\n", lnt, lat)
	//T16FG034.L 6.115, 8.237
	lat, lnt = Lat(6.115, 8.237)
	fmt.Printf("T16FG034.L,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG034.L Lat =", lat)
	fmt.Println("T16FG034.L Lng =", lnt)
	//T16FG034.A 11.936, 10.209
	lat, lnt = Lat(11.936, 10.209)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG034.A Lat =", lat)
	fmt.Println("T16FG034.A Lng =", lnt)
	//T16FG034.R 20.303, 8.195
	lat, lnt = Lat(20.303, 8.195)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG034.R Lat =", lat)
	fmt.Println("T16FG034.R Lng =", lnt)

	// TE_3TO2A   74.154,8.572
	lat, lnt = Lat(20.303, 8.195)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG034.R Lat =", lat)
	fmt.Println("T16FG034.R Lng =", lnt)
	//TE_3TO2B    74.024,1.461
	lat, lnt = Lat(20.303, 8.195)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG034.R Lat =", lat)
	fmt.Println("T16FG034.R Lng =", lnt)

	//T16FG032.L 11.068, 8.708
	lat, lnt = Lat2(11.068, 8.708)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG032.L Lat =", lat)
	fmt.Println("T16FG032.L Lng =", lnt)
	//T16FG032.R 24.008, 8.653
	lat, lnt = Lat2(24.008, 8.653)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG032.R Lat =", lat)
	fmt.Println("T16FG032.R Lng =", lnt)
	//T16FG031.R 24.072,1.508
	lat, lnt = Lat2(24.072, 1.508)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG031.R Lat =", lat)
	fmt.Println("T16FG031.R Lng =", lnt)
	//T16FG031.L 10.412,1.531
	lat, lnt = Lat2(10.412, 1.531)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG031.L Lat =", lat)
	fmt.Println("T16FG031.L Lng =", lnt)
	//T16FG032.A 16.896, 10.260
	lat, lnt = Lat2(16.896, 10.260)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG032.A Lat =", lat)
	fmt.Println("T16FG032.A Lng =", lnt)
	//T16FG032.B 15.276, 5.407
	lat, lnt = Lat2(15.276, 5.407)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG032.B Lat =", lat)
	fmt.Println("T16FG032.B Lng =", lnt)
	//T16FG031.B 20.552, 3.557
	lat, lnt = Lat2(20.552, 3.557)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG031.B Lat =", lat)
	fmt.Println("T16FG031.B Lng =", lnt)
	//T16FG031.A 16.469, -0.101
	lat, lnt = Lat2(16.469, -0.101)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG031.A Lat =", lat)
	fmt.Println("T16FG031.A Lng =", lnt)
	//T16FG029.L 84.193, 1.829
	lat, lnt = Lat2(84.193, 1.829)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG029.L Lat =", lat)
	fmt.Println("T16FG029.L Lng =", lnt)
	//T16FG030.L 82.344, 9.20
	lat, lnt = Lat2(82.344, 9.20)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("T16FG030.L Lat =", lat)
	fmt.Println("T16FG030.L Lng =", lnt)

	// TE_3TO2A   74.154,8.572
	lat, lnt = Lat(74.154, 8.572)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("TE_3TO2A Lat =", lat)
	fmt.Println("TE_3TO2A Lng =", lnt)
	//TE_3TO2B    74.024,1.461
	//3to2a 74.432, 11.279
	//3to2b 73.940, -0.747
	//2to3b 11.056, -0.445
	//2to3a 10.995, 11.774
	lat, lnt = Lat(74.024, 1.461)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("TE_3TO2B Lat =", lat)
	fmt.Println("TE_3TO2B Lng =", lnt)

	//TE_2TO3A    10.995,11.774
	lat, lnt = Lat2(10.995, 11.774)
	fmt.Printf("T16FXG036.1,Lng =%f,Lat =%f\n", lnt, lat)
	fmt.Println("TE_2TO3A Lat =", lat)
	fmt.Println("TE_2TO3A Lng =", lnt)
	//TE_2TO3B  11.056,-0.445
	lat, lnt = Lat2(11.056, -0.445)
	fmt.Printf("TE_2TO3B,Lng =%f,Lat =%f\n", lnt, lat)

	////3to2a 74.432, 11.279
	//lat, lnt = Lat(74.432, 11.279)
	//fmt.Println("3to2a Lat =", lat)
	//fmt.Println("3to2a Lng =", lnt)
	////3to2b 73.940, -0.747
	//lat, lnt = Lat(73.940, -0.747)
	//fmt.Println("3to2b Lat =", lat)
	//fmt.Println("3to2b Lng =", lnt)
	////2to3b 11.056, -0.445
	//lat, lnt = Lat2(11.056, -0.445)
	//fmt.Println("2to3b Lat =", lat)
	//fmt.Println("2to3b Lng =", lnt)
	////2to3a 10.995, 11.774
	//lat, lnt = Lat2(10.995, 11.774)
	//fmt.Println("2to3a Lat =", lat)
	//fmt.Println("2to3a Lng =", lnt)

}
