package goset

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func GoSet() {
	b := mapset.NewSet()
	for va := range b.Iter() {
		b.Add(va)
	}
	fmt.Println(b)
}
