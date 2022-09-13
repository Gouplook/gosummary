package godecimal

import (
	"fmt"
	"testing"
)

// Big Data Processing
func TestBigDataProcess(t *testing.T) {
	BigDataProcess()
}

func TestDistanceBetween(t *testing.T) {
	lat1, lng1 := 32.060255, 118.796877
	lat2, lng2 := 39.904211, 116.407395

	dis := DistanceBetween(lat1, lat2, lng1, lng2)
	fmt.Printf("%fkm", dis)

}

func TestDistance(t *testing.T) {
	//Distance()
	//Distance(0, 0, 0)
	Location()
}

func TestTimeStamp(t *testing.T) {
	//TimeStamp()
	MouNiShu()
}
