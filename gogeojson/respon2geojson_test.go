//go:build !windows

package gogeojson

import (
	"testing"
)

func TestResponse2Geojson(t *testing.T) {
	Response2Geojson()
}

// 多站点距离请求
func TestMultipleResponse2Geojson(t *testing.T) {
	MultipleResponse2Geojson()
	// 定义弧度值
}
