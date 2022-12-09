package gostrings

import (
	"fmt"
	"strings"
	"testing"
)

func TestGoStrings2(t *testing.T) {
	GoStrings2()
	s := "e200.situ.car9"
	index := strings.Index(s, ".") // 4

	str := s[0:index]
	switch str {
	case "e100":
		fmt.Println("A")
	case "e200":
		fmt.Println("b")
	default:
		fmt.Println("....")
	}

	fmt.Println(str)

}
