package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	x, y    int     // 节点坐标
	g, h, f float64 // g: 起点到当前节点的距离, h: 当前节点到终点的估算距离, f: g + h
	parent  *Node   // 父节点
}

func NewNode(x, y int, parent *Node, g, h float64) *Node {
	return &Node{
		x:      x,
		y:      y,
		parent: parent,
		g:      g,
		h:      h,
	}
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	node := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return node
}

func reverse(nodes []*Node) []*Node {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
	return nodes
}
func AStar(startX, startY, endX, endY int, obstacles [][]bool) []*Node {
	// 地图边界
	maxY, maxX := len(obstacles), len(obstacles[0])
	fmt.Println(maxY, maxX)
	// openList中的节点
	start := NewNode(startX, startY, nil, 0, manhattanDistance(startX, startY, endX, endY))
	openList := PriorityQueue{start}
	// 在closedList中的节点
	closedList := make(map[Node]struct{})

	for len(openList) > 0 {
		// 从openlist中获取f值最小的节点
		currentNode := heap.Pop(&openList).(*Node)

		// 如果当前节点就是终点，返回路径
		if currentNode.x == endX && currentNode.y == endY {
			path := make([]*Node, 0)
			for currentNode != nil {
				path = append(path, currentNode)
				currentNode = currentNode.parent
			}
			reverse(path)

			return path
		}

		// 将当前节点加入closedList
		closedList[*currentNode] = struct{}{}

		// 寻找当前节点的邻居
		neighbors := []*Node{
			NewNode(currentNode.x, currentNode.y-1, currentNode, currentNode.g+1, manhattanDistance(currentNode.x, currentNode.y-1, endX, endY)),
			NewNode(currentNode.x, currentNode.y+1, currentNode, currentNode.g+1, manhattanDistance(currentNode.x, currentNode.y+1, endX, endY)),
			NewNode(currentNode.x-1, currentNode.y, currentNode, currentNode.g+1, manhattanDistance(currentNode.x-1, currentNode.y, endX, endY)),
			NewNode(currentNode.x+1, currentNode.y, currentNode, currentNode.g+1, manhattanDistance(currentNode.x+1, currentNode.y, endX, endY)),
		}

		for _, neighbor := range neighbors {
			// 如果是障碍，则跳过
			if obstacles[neighbor.y][neighbor.x] {
				continue
			}
			// 如果在closedList中，则跳过
			if _, ok := closedList[*neighbor]; ok {
				continue
			}
			// 如果不在openList中，则加入
			if index := pqIndex(openList, neighbor); index == -1 {
				heap.Push(&openList, neighbor)
			} else {
				// 如果在openList中，则比较新g值和旧g值的大小
				oldNode := openList[index]
				if neighbor.g < oldNode.g {
					openList[index] = neighbor
					heap.Fix(&openList, index)
				}
			}
		}
	}
	return nil
}

func pqIndex(pq PriorityQueue, node *Node) int {
	for i, n := range pq {
		if n.x == node.x && n.y == node.y {
			return i
		}
	}
	return -1
}

func manhattanDistance(x1, y1, x2, y2 int) float64 {
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))
	return dx + dy
}

func main() {

}
