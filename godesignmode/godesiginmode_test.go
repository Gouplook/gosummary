package godesignmode

import (
	"fmt"
	"testing"
)

func Test_mode(t *testing.T) {

	operator := Operator{}
	operator.setStrategy(&add{})
	rs := operator.calculate(1, 2)

	fmt.Println(rs)

	operator.setStrategy(&reduce{})
	result := operator.calculate(2, 1)
	fmt.Println("reduce:", result)

}
