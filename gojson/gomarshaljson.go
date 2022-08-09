package gojson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type JsonT struct {
	Name string `json:"name" default:"bbc"`
	Age  int    `json:"age" default:"-1"`
}

func MarshalJSON(i interface{}) ([]byte, error) {
	typeof := reflect.TypeOf(i)
	valueof := reflect.ValueOf(i)
	for i := 0; i < typeof.Elem().NumField(); i++ {
		if valueof.Elem().Field(i).IsZero() {
			def := typeof.Elem().Field(i).Tag.Get("default")
			if def != "" {
				switch typeof.Elem().Field(i).Type.String() {
				case "int":
					result, _ := strconv.Atoi(def)
					valueof.Elem().Field(i).SetInt(int64(result))
				case "string":
					valueof.Elem().Field(i).SetString(def)
				}
			}
		}
	}
	return json.Marshal(i)
}

func www() {
	t := &JsonT{
		Name: "YS",
		//Age:  1,
	}
	data, err := MarshalJSON(t)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
