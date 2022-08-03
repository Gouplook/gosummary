package main

// 车辆基本结构
type Car struct {
	CarId     int     // 车辆ID
	CarType   int     // 车辆类型  0=AGV...
	Lng       float64 // 位置经度
	Lat       float64 // 位置维度
	Theta     float64 // 位置角度
	Speed     float64 // 车速
	TimeStamp int64   // 时间戳
}

// 计算文件过程文件
type CarBody struct {
	CarId int     // 车辆编号
	Speed float32 // 车速
	Stops int     // 停车次数
}

// 路段
type Lane struct {
	LaneId     int     // 路段Id
	LaneLength float32 // 道路长度
	FlowMax    int     // 最大通行量（车辆数）

}

type LaneBody struct {
	LaneId  int     // 路段Id
	Number  int     // 车辆数量len（head） 实在变化的
	FlowMax int     // 最大通行量（车辆数）
	Density float32 // 车辆密度
}

type LaneLink struct {
	Car // 车节点
}

func main() {
	var LaneArray [100]Lane
	// 遍历所有路段Id
	for i := 0; i <= len(LaneArray); i++ {

	}
	// 车道与车辆的关系

	// 道路与车的关系，一个车道中
	// 车道链表
	var Lande [int]Car
}
