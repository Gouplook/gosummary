package gojson

import (
	"gonum.org/v1/gonum/mat"
	"gosummary/kalman"
	"math"
	"sync"
)

type SpeedCalculation struct {
	lock   sync.RWMutex
	size   int
	offset int
	cache  []CarStatus
	//matchData []CarStatus
}

func NewSpeedCalculation(size int) *SpeedCalculation {
	return &SpeedCalculation{
		offset: 0,
		cache:  make([]CarStatus, size, size),
	}
}

// 速度平滑
type SpeedWithKalman struct {
	ctx      *kalman.Context
	config   Config
	filter   kalman.Filter
	control  *mat.VecDense
	filtered mat.Vector
}

func NewSpeedWithKalman(c Config) *SpeedWithKalman {

	return &SpeedWithKalman{
		// todo
	}
}

// 数据入口
func (s *SpeedWithKalman) Apply(data CarStatus) {
	measurement := mat.NewVecDense(2, []float64{
		data.Speed * math.Cos(data.Theta),
		data.Speed * math.Sin(data.Theta),
	})
	s.filtered = s.filter.Apply(s.ctx, measurement, s.control)
}

func (s *SpeedWithKalman) Predict() interface{} {

	speed := math.Sqrt(
		math.Pow(s.filtered.AtVec(0), 2) + math.Pow(s.filtered.AtVec(1), 2),
	)
	if math.IsNaN(speed) {
		speed = 0.0
	}

	return speed
}
