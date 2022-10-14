package gojson

type MapMessage struct {
	Id         int64   `json:"id"`
	LaneId     int64   `json:"lane_id"`
	LaneLength float64 `json:"lane_length"`
	LaneType   int     `json:"lane_type"` // 车辆类型，2= 路口
	SegId      int64   `json:"seg_id"`
}

type CarStatus struct {
	CarId     string       `json:"car_id"`
	CarType   int          `json:"car_type"`
	Lat       float64      `json:"lat"`
	Lng       float64      `json:"lng"`
	Theta     float64      `json:"theta"`
	Speed     float64      `json:"speed"`
	PreMapMsg []MapMessage `json:"pre_map_msg"`
	MapMsg    []MapMessage `json:"map_msg"`
	TimeStamp float64      `json:"time_stamp"`
}

type Config struct {
}

type Register interface {
	AddFeatures(name string, filter Filter)
}

type Filter interface {
	Apply(status CarStatus)
	Predict() interface{}
}

type StreamHandle interface {
	Next(status CarStatus)
}