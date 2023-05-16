package godecimal

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

func BigDataProcess() {
	add := decimal.NewFromFloat(123.66).Add(decimal.NewFromFloat(22))
	dec := decimal.NewFromFloat(123.66).Sub(decimal.NewFromFloat(22))
	mul := decimal.NewFromFloat(10).Mul(decimal.NewFromFloat(22))
	div := decimal.NewFromFloat(float64(12)).Div(decimal.NewFromFloat(float64(7)))
	div2, b := decimal.NewFromFloat(float64(12)).Div(decimal.NewFromFloat(float64(7))).Truncate(4).Float64()
	// divStr, b := decimal.NewFromFloat(float64(12)).Div(decimal.NewFromFloat(float64(7))).Truncate(4).String()

	fmt.Println(add)
	fmt.Println(dec)
	fmt.Println(mul)
	fmt.Println(div)
	fmt.Println("div=", div2, b)
	fmt.Println(add)

	var RealPrice float64 = 10
	var Price float64 = 50

	// 保留2位有效数字
	// 充值卡： 充10 送50元，  10/（10+50） 面值 60，实际支付 10 元
	discount, _ := decimal.NewFromFloat(RealPrice).Div(decimal.NewFromFloat(Price + RealPrice)).Truncate(1).Float64()
	fmt.Println("dicount= ", discount)
}

// DistanceBetween 计算两个经纬度距离
func DistanceBetween(lat1, lat2, lng1, lng2 float64) float64 {
	radius := 6378.137 // 地球半径

	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lat2 = lat2 * rad
	lng1 = lng1 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

// 计算地图位置信息
func Distance(sL float64, eL float64, d float64) (float64, float64, string) {
	//x0 := -197011.83
	//y0 := 3409274.43
	c := 3.88
	//
	x1 := sL + math.Cos(c)*d
	y1 := eL + math.Sin(c)*d
	//
	//fmt.Println("x1=", x1)
	//fmt.Printf("y1=%.2f", y1)
	//fmt.Println(time.Second / 2)

	return x1, y1, "101"

}
