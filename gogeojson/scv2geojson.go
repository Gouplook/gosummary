package gogeojson

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type Properties struct {
	Name string `json:"name"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type CRSProperties struct {
	Name string `json:"name"`
}

type CRS struct {
	Properties CRSProperties `json:"properties"`
	Type       string        `json:"type"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
	CRS      CRS       `json:"crs"`
}

func Svs2Geojson() {
	// 读取CSV文件
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// 解析CSV数据
	csvData, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// 转换为GeoJSON格式
	feature := Feature{
		Type: "Feature",
		Geometry: Geometry{
			Type: "LineString",
			Coordinates: [][]float64{
				{
					parseCoordinate(csvData[0]),
					parseCoordinate(csvData[1]),
				},
			},
		},
		Properties: Properties{
			Name: csvData[2],
		},
	}

	features := []Feature{feature}

	crsProperties := CRSProperties{
		Name: "urn:ogc:def:crs:EPSG::32451",
	}

	crs := CRS{
		Properties: crsProperties,
		Type:       "name",
	}

	featureCollection := FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
		CRS:      crs,
	}

	// 转换为JSON字符串
	jsonData, err := json.MarshalIndent(featureCollection, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// 将GeoJSON写入文件
	err = writeToFile("demo.geojson", jsonData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GeoJSON写入文件成功")
}

func parseCoordinate(coordStr string) float64 {
	var coordinate float64
	// 这里可以根据实际的数据格式进行解析
	fmt.Sscanf(coordStr, "%f", &coordinate)
	return coordinate
}

func writeToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
