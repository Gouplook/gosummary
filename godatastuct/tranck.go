package main

import "fmt"

// 动态交通计算状态思路:

//速度：
//  初始条件：车辆500ms上报一次信息，地图匹配也500ms匹配一次，交通状态服务计算周期5s，路段动态交通状态10min更新一次。
//  500ms车辆上报一次信息，进行车辆与道路匹配。根据车辆相邻的位置信息和时间周期，滚动计算（计算周期内如5s计算一次）车辆的速度，对于Tn某个时刻，
//  或者几个时刻，位置信息丢失，或者存在噪音数据，利用卡尔曼滤波算法，推算出位置信息和速度行，过滤掉噪音数据。

// 流量：计算周期5s
// 车道laneId = 1 在5s内通过车辆数N
// V = N/T

// 密度：计算周期5s
// 车道laneId = 1 在5s内通过车辆数N
// K= N/L   L=道路长度。

// 5s   carId = 1 :v= 30km/h    laneId = 5
// 5s   carId = 2 :v= 40km/h    laneId = 7
// 5s   carId = 3 :v= 20km/h    laneId = 5

// 10min，
//10*60/5 = 120 个时间周期的数据，计算出道路如LaneId = 1  根据周期（10min）多辆车在该道路上速度，和车密度，

// 根据500ms周期向后计算，统计出车辆的位置，车辆在道路上的数量，

// 车辆节点信息
type CarsNode struct {
	CarId     int     // 车辆编号
	CarType   int     // 车辆类型  0=AGV...
	Lng       float64 // 位置经度
	Lat       float64 // 位置维度
	Theta     float64 // 位置角度
	Speed     float64 // 车速
	TimeStamp int64   // 时间戳
}

type LaneBody struct {
	Number     int      // 车辆数量len（head）
	FlowMax    int      // 最大通行量（车辆数）
	Lanelength float64  //路段长度
	Head       *CarBody //
}

type CarBody struct {
	Speed float64 // 当前速度
	Stops int     // 停车次数
	// ...... 基于车辆的其它指标
	History []CarsNode // 缓存历史上报点位置；维护定长区域(优化速度计算)
	Next    *CarBody
}

func main() {
	var LaneArray []LaneBody
	var Cycle []LaneBody

	fmt.Println(LaneArray)
	fmt.Println(Cycle)
}
