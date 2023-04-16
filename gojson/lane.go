package gojson

import "sync"

type LaneBody struct {
	lock    sync.RWMutex
	laneId  int64
	config  Config
	feature map[string]Filter
	carIds  map[string]float64
}

func NewLaneBody(laneId, len int64, c Config) *LaneBody {
	newLaneBody := LaneBody{
		lock:   sync.RWMutex{},
		laneId: laneId,
		config: c,
		carIds: map[string]float64{},
	}
	//
	return &newLaneBody
}

func (l *LaneBody) Next(data CarStatus) {

}

func (l *LaneBody) AddFeatures(name string, f Filter) {
	if l.feature == nil {
		l.feature = map[string]Filter{}
	}
	l.feature[name] = f
}

// refresh 数据的更新，缓存一段时间数据
func (l *LaneBody) refresh() {

}

func (l *LaneBody) GetCars() []string {
	l.refresh()
	var ids []string
	l.lock.RLock()
	defer l.lock.RUnlock()
	for k, _ := range l.carIds {
		ids = append(ids, k)
	}
	return ids
}

type LaneMap struct {
	config Config
	lock   sync.RWMutex
	lanes  map[int64]*LaneBody
	cars   *CarMap
}

func NewLaneMap(c Config) *LaneMap {
	return &LaneMap{
		config: c,
		lock:   sync.RWMutex{},
		lanes:  map[int64]*LaneBody{},
		cars:   NewCarMap(c),
	}
}
