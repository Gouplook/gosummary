package gojson

const (
	SPEED    = "speed"
	FLOW     = "flow"
	STOP     = "stop"
	DENSITY  = "density"
	LINEUP   = "lineup"
	LOCATION = "location"
)

func CarRegister(r Register, config Config) {
	// todo
	r.AddFeatures(SPEED, NewSpeedWithKalman(config))
}

func LanesRegister(r Register, config Config) {
	// todo
}
