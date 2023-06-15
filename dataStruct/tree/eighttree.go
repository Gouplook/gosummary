package main

type Point struct {
	X, Y float64
}

type Region struct {
	X     float64
	Y     float64
	Widht float64
	Hight float64
}

type NodeTree struct {
	Value     Point
	childrens []*NodeTree
}
