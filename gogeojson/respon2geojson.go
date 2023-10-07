package gogeojson

import (
	"bytes"
	"encoding/json"
	geojson "github.com/paulmach/go.geojson"
	"io/ioutil"
	"os"
	"strconv"
)

type Point struct {
	Lng   float64 `json:"lng"` //经度（x）
	Lat   float64 `json:"lat"`
	Theta float64 `json:"theta"`
}

type GetPathResponse struct {
	Points []Point `json:"points"`
}

type NTPoint struct {
	Lat       float64                `json:"lat"`
	Lng       float64                `json:"lng"`
	Z         float64                `json:"z,optional,omitempty,default=0"`
	Theta     *float64               `json:"theta,optional"`
	NcTheta   *float64               `json:"nc_theta,optional,omitempty"`
	Curvature *float64               `json:"curvature,optional,omitempty"`
	Extend    map[string]interface{} `json:"properties,optional,omitempty"`
}

type MultiplePoint struct {
	Origin      NTPoint   `json:"origin"`
	Destination NTPoint   `json:"destination"`
	Distance    float64   `json:"distance,optional"`
	Code        int       `json:"code,optional"`
	Route       []NTPoint `json:"route,optional"`
}

type MultipleResponse struct {
	Points []MultiplePoint `json:"points"`
}

// 多组点结构2Geojson
func MultipleResponse2Geojson() {
	by, err := ioutil.ReadFile("./multiple.json")
	if err != nil {
		panic(err)
	}
	response := MultipleResponse{}
	if err = json.Unmarshal(by, &response); err != nil {
		panic(err.Error())
	}

	for i, points := range response.Points {
		dense := make([][]float64, 0)
		for _, point := range points.Route {
			dense = append(dense, []float64{point.Lng, point.Lat})
		}
		// 写文件
		writeGeojson(i, dense)
	}

}

// nohup ./road -f etc/road.yaml &

func writeGeojson(wId int, dense [][]float64) {
	out := geojson.NewFeatureCollection()
	if out.CRS == nil {
		out.CRS = make(map[string]interface{})
	}
	out.CRS["type"] = "name"
	out.CRS["properties"] = map[string]string{"name": "urn:ogc:def:crs:EPSG::32451"}

	lane := geojson.NewLineStringFeature(dense)
	lane.SetProperty("name", "dense")
	out.AddFeature(lane)
	//序列化
	reJson, _ := out.MarshalJSON()
	var filewrite *os.File
	name := strconv.Itoa(wId) + ".geojson"
	filewrite, _ = os.Create(name)
	_, err := filewrite.WriteString(FormatJson(string(reJson)))
	if err != nil {
		panic(err)
	}
	filewrite.Close()
}

func Response2Geojson() {
	by, err := ioutil.ReadFile("./route.json")
	if err != nil {
		panic(err)
	}
	response := GetPathResponse{}
	if err = json.Unmarshal(by, &response); err != nil {
		panic(err.Error())
	}
	out := geojson.NewFeatureCollection()
	if out.CRS == nil {
		out.CRS = make(map[string]interface{})
	}
	out.CRS["type"] = "name"
	out.CRS["properties"] = map[string]string{"name": "urn:ogc:def:crs:EPSG::32451"}
	dense := make([][]float64, 0)
	for _, point := range response.Points {
		dense = append(dense, []float64{point.Lng, point.Lat})
	}
	// 生成线
	lane := geojson.NewLineStringFeature(dense)
	lane.SetProperty("name", "dense")
	out.AddFeature(lane)

	//序列化
	reJson, _ := out.MarshalJSON()
	var filewrite *os.File
	filewrite, _ = os.Create("./mutliipRoute-3.geojson")
	defer filewrite.Close()
	_, err = filewrite.WriteString(FormatJson(string(reJson)))
	if err != nil {
		panic(err)
	}

}
func FormatJson(data string) string {
	var out bytes.Buffer
	_ = json.Indent(&out, []byte(data), "", "    ")
	return out.String()
}
