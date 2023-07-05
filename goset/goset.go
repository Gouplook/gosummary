package goset

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func GoSet() {
	b := mapset.NewSet()
	b.Add(1)
	b.Add(2)
	b.Add(66)

	fmt.Println(b.Clone())
	fmt.Println(b.Contains(66))
	for nl := range b.Iter() {
		fmt.Println(nl)
	}
}
