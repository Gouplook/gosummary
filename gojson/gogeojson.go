package gojson

import (
	"fmt"
	geojson "github.com/paulmach/go.geojson"
)

var carsF map[string]map[string]interface{}

func GoGeoJson() {
	geo := geojson.NewFeatureCollection()
	if geo.CRS == nil {
		geo.CRS = make(map[string]interface{})
	}
	//  每个因子都需要有
	// CarBody {carId, features map[string]Filter } 车辆ID，车辆特征（入口，数据,出口：指标值预测值）
	// CarMap { cars map[string]*CarBody } {car_id: carBody}
	// Lane LaneBody { laneId, feature map[string]Filter, carIds map[string]float64
	//      LaneMap  { cars *CarMap, lanes map[int64]*LaneBody } {car_id:laneBody }
	//

	for _, fs := range carsF {
		if cars, ok := fs["name"]; !ok {
			fmt.Println(cars)
		}
	}
	fmt.Println(12 + 22 + 7.7 + 6 + 38 + 56 + 3 + 2 + 4 + 4 + 3)

}
