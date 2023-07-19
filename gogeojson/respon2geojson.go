package gogeojson

import (
	"bytes"
	"github.com/goccy/go-json"
	geojson "github.com/paulmach/go.geojson"
	"io/ioutil"
	"os"
)

type Point struct {
	Lng   float64 `json:"long"` //经度（x）
	Lat   float64 `json:"lat"`
	Theta float64 `json:"theta"`
}

type GetPathResponse struct {
	Points []Point `json:"points"`
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
	filewrite, _ = os.Create("./path-1.geojson")
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
