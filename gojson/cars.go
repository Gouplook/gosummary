package gojson

import "sync"

type CarBody struct {
	lock     sync.RWMutex
	config   Config
	carId    string
	features map[string]Filter
}

func NewCarBody(c Config, id string) *CarBody {
	newCar := CarBody{
		lock:   sync.RWMutex{},
		config: c,
		carId:  id,
	}
	//CarRegister(&newCar, c)

	return &newCar
}

func (c *CarBody) Next(data CarStatus) {
	for _, f := range c.features {
		f.Apply(data)
	}
}

func (c *CarBody) AddFeatures(name string, f Filter) {
	// todo
}

type CarMap struct {
	lock   sync.RWMutex
	config Config
	cars   map[string]*CarBody
}

func NewCarMap(c Config) *CarMap {
	return &CarMap{
		lock:   sync.RWMutex{},
		config: c,
		cars:   map[string]*CarBody{},
	}
}
func (c *CarMap) Next(data CarStatus) {
	// todo
	carId := data.CarId
	c.lock.Lock()
	defer c.lock.Unlock()
	if car, ok := c.cars[carId]; ok {
		car.Next(data)
	} else {
		newCar := NewCarBody(c.config, carId)
		c.cars[carId] = newCar
		c.cars[carId].Next(data)
	}
}
