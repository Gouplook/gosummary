package main

import "fmt"

type MotData struct {
	frameList []struct{}
}

func main() {
	m := new(MotData)
	a := []bool{true, false, true, true, false}
	for _, ok := range a {
		if ok {
			m.frameList = append(m.frameList, struct{}{})
		}
	}
	fmt.Println(m.frameList)

}
