package goInterface

import (
	"fmt"
	"math"
	"testing"
)

func Test_Person(t *testing.T) {

	//var person Person
	stu := Student{
		Name: "DZ",
		Age:  12,
	}
	fmt.Println(stu.GetName())
	//stu.GetName()
	var per Person
	per = &stu

	fmt.Println(per.GetName(), per.GetAge())

	length := int64(math.MaxInt64)
	fmt.Println(length)

}
