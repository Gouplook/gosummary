package main

import (
	"fmt"
	"math"
	"math/rand"
)

//模拟退火算法
const (
	MaxTemp     = 1000.0 // 最大温度
	MinTemp     = 1e-4   // 最小温度
	DeltaFactor = 0.99   // 降温系数
	N           = 100    // 搜索次数
)

type Point struct { // 点的结构体
	X, Y float64
}

func Dist(p1, p2 Point) float64 { // 计算两点之间距离
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

// SimulatedAnnealing 模拟退火算法解决路径搜索问题
func SimulatedAnnealing(points []Point) []Point {
	current, best := make([]int, len(points)), make([]int, len(points))
	for i := range current { // 初始化路径
		current[i] = i
	}
	best = current
	temp := MaxTemp
	for temp > MinTemp { // 降温
		for i := 0; i < N; i++ { // 搜索次数
			j, k := rand.Intn(len(points)), rand.Intn(len(points))
			current[j], current[k] = current[k], current[j] // 交换两个随机位置的点
			dE := 0.0
			for l := 1; l < len(points); l++ { // 计算路径长度变化
				dE += Dist(points[current[l]], points[current[l-1]]) - Dist(points[best[l]], points[best[l-1]])
			}
			if dE < 0 || math.Exp(-dE/temp) > rand.Float64() { // 更新路径
				copy(best, current)
			} else {
				current[j], current[k] = current[k], current[j]
			}
		}
		temp *= DeltaFactor
	}
	result := make([]Point, len(points))
	for i := range best { // 根据路径计算最终顺序
		result[i] = points[best[i]]
	}
	return result
}

func main() {
	points := []Point{{0, 0}, {4, 4}, {1, 1}, {2, 2}, {3, 3}} // 五个点
	result := SimulatedAnnealing(points)
	for _, p := range result { // 输出最优顺序
		fmt.Printf("(%v, %v) ", p.X, p.Y)
	}
}
