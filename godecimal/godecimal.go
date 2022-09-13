package godecimal

import (
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
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

// 计算两个经纬度距离

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

// 路测数据。

func Location() {
	t := 66
	//t2 := 56
	//t3 := 50
	//t4 := 42
	//t5 := 32
	//t6 := 19

	// 模拟多车发送。
	// 500ml 发送一次
	var x0, y0 float64
	var carLocation = map[string][]float64{} // key :carId
	// 对于101车两次减速,三段速度，1.2, 0.8, 0.6,
	//carLocation["101"] = []float64{-197011.83, 3409274.43} // laneId= 0 15.89m  13s
	//carLocation["101"] = []float64{-197023.37, 3409263.93} // laneId= 0 80m  20.591 27s
	//carLocation["101"] = []float64{-197039.34, 3409249.39} // laneId= 0  0.6m/s   25.220 42s
	//carLocation["101"] = []float64{-197057.98, 3409232.43} // laneId= 0  0.4m/s   19.876 49s
	carLocation["101"] = []float64{-197072.48, 3409219.24} // laneId= 0  0.2m/s   2.383 12s

	// 对102车，由于101车进行了减速，102，速度直接降为0.8
	//carLocation["102"] = []float64{-197019.99, 3409267.03} // LaneId = 1 0.8m/s  25.240m  32s
	//carLocation["102"] = []float64{-197038.92, 3409249.80} // LaneId = 1  0.6m/s 24.472m  40s
	//carLocation["102"] = []float64{-197056.67, 3409233.65} // LaneId = 1  0.4m/s   19.683m  49s
	carLocation["102"] = []float64{-197071.17, 3409220.46} // LaneId = 1 0.2m/s  2.83m  56.6s
	// 对103车，由于102车进行了减速，103，速度直接降为0.8
	//carLocation["103"] = []float64{-197027.41, 3409260.28} // laneId = 1  0.8m/s  15.088M  19s
	//carLocation["103"] = []float64{-197038.65, 3409250.05} // laneId = 1    0.6m/s 24.14M  40s
	//carLocation["103"] = []float64{-197056.40, 3409233.90} // laneId = 1  0.4m/s 21.058   35s
	carLocation["103"] = []float64{-197066.75, 3409224.48} // laneId = 1  0.2m/s 2.387M  12s

	// 由于101车第二次减速，104，速度直接降为0.8
	//carLocation["104"] = []float64{-197035.43, 3409253.02} // laneId = 73 0.8m/s 5.092m  6s
	//carLocation["104"] = []float64{-197038.98, 3409249.79} // laneId = 73 0.6m/s 24.815m  41s
	//carLocation["104"] = []float64{-197057.17, 3409233.23} // laneId = 73 0.4m/s 20.406m 51s
	carLocation["104"] = []float64{-197072.26, 3409219.50} // laneId = 73 0.2m/s 2.891m 14s

	//carLocation["105"] = []float64{-197042.80, 3409246.30} // laneId = 73 0.6m/s 19.620m 32s
	//carLocation["105"] = []float64{-197057.00, 3409233.38} // laneId = 73  0.4m/s 19.175m  47s
	carLocation["105"] = []float64{-197070.90, 3409220.73} // laneId = 73  0.2m/s 2.891m 14s

	//carLocation["106"] = []float64{-197053.83, 3409236.17} // laneId = 73  0.6m/s 4.986m 8s
	//carLocation["106"] = []float64{-197057.38, 3409232.94} // laneId = 73  0.4m/s 19.85m  50s
	carLocation["106"] = []float64{-197071.28, 3409220.29} // laneId = 73  0.2m/s 2.321m 11s

	//for i := 0; i < 2*t; i++ {
	//	if value, ok := carLocation["102"]; ok {
	//		x0, y0 = value[0], value[1]
	//	}
	//	xl, yl, carId := Distance(x0, y0, d)
	//	x0, y0 = xl, yl
	//	fmt.Printf("carId=%s,xl=%.2f,yl=%.2f,time=%s\n", carId, xl, yl, time.Now().Local())
	//	time.Sleep(time.Second / 2)
	//}

	//d := 1.2 * 0.5
	//d1 := 0.8 * 0.5
	//d2 := 0.6 * 0.5
	//d3 := 0.4 * 0.5
	// 5+
	//
	// 思路：模拟数据发送。n台车，如何模拟间隔发送。
	// 101 车出发时间：9:25:31
	// 102 车出发时间：9:25:39    // 相隔8s
	// 103 车出发时间：9:26:3    //  32s
	// 104 车出发时间：9:26:10   //  39s
	// 105 车出发时间：9:26:21  //  50s

	// 101 先上传数据,

	d4 := 0.2 * 0.5
	if value, ok := carLocation["106"]; ok {
		x0, y0 = value[0], value[1]
		t = 11
		for i := 0; i < 2*t; i++ {
			xl, yl, _ := Distance(x0, y0, d4)
			x0, y0 = xl, yl
			fmt.Printf("carId=%s,xl=%.2f,yl=%.2f,timestmap=%d\n", "106", xl, yl, time.Now().Local().Unix())
			time.Sleep(time.Second / 2)
		}
	}
}

func TimeStamp() {
	for i := 0; i < 286; i++ {
		fmt.Printf("timesstmap =%d\n", (time.Now().Local().UnixNano())/1e6)
		time.Sleep(time.Second / 2)
	}

}

func MouNiShu() {
	// 从生成的数据里，给每条记录加上(当前时间）时间戳。
	fileName := "./101.yaml"
	file, err := os.OpenFile(fileName, 0, 0777)
	if err != nil {
		fmt.Println("file Open failed .....")
	}

	datas, err := ioutil.ReadAll(file)
	lanes := strings.Split(string(datas), "\n")
	fmt.Println(len(lanes)) //
	re := regexp.MustCompile(`carId=([^,]+),xl=([^,]+),yl=([^,]+),timesstmap\s*=([^,]+)`)
	for _, lane := range lanes {
		if strings.TrimSpace(lane) == "" {
			//fmt.Println("read space line from file.")
			continue
		}

		fields := re.FindStringSubmatch(lane)
		time.Sleep(time.Second / 2)
		//fmt.Println(fields)
		timeStamp := strings.TrimSpace(fields[4])
		timeStamp = "0094"

		// 定时发送。

		fmt.Println(timeStamp)

	}

}
