package gostrings

import (
	"fmt"
	"testing"
)

func TestGoStrings2(t *testing.T) {
	//GoStrings2()
	//s := "e200.situ.car9"
	//index := strings.Index(s, ".") // 4
	//
	//str := s[0:index]
	//switch str {
	//case "e100":
	//	fmt.Println("A")
	//case "e200":
	//	fmt.Println("b")
	//default:
	//	fmt.Println("....")
	//}

	//fmt.Println(str)
	//laneId := 1000727379968 //
	//fmt.Printf("seg_id=%d,lane_id=%d\n", laneId>>32, laneId<<32>>32)

	//b := strings.Contains("e200.situ.car10", "situ")
	//fmt.Printf("b=%v\n", b)

	//s := fmt.Sprintf("waypoints(%v) can", 1)
	//fmt.Println(s)
	a := []string{"abc", "bcd", "bv"}
	s := Join(a, ":")
	fmt.Println(s)

}

func TestGoReader(t *testing.T) {
	GoReader()
}
