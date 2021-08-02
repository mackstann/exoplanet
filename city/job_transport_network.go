package city

import (
	"log"

	"github.com/mackstann/exopolis/heatsim"
)

type JobTransportNetwork struct {
	city *City
	*heatsim.HeatGrid
}

const (
	dirtConductivity    float64 = 0.1
	roadConductivity            = 0.9
	defaultConductivity         = 0.0
)

func NewJobTransportNetwork(city *City) *JobTransportNetwork {
	temperature := func(x int, y int) (float64, bool) {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			return 0, false
		}
		return (*city)[y][x].Resources.Jobs, true
	}
	setTemperature := func(x int, y int, val float64) {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			log.Panicf("setTemperature: out of bounds: (%d,%d)", x, y)
		}
		(*city)[y][x].Resources.Jobs = val
	}
	getConductivity := func(x int, y int) (float64, bool) {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			return 0, false
		}

		switch (*city)[y][x].Typ {
		case Dirt:
			return dirtConductivity, true
		case Road:
			return roadConductivity, true
		default:
			return defaultConductivity, true
		}
	}
	// TODO: pointer makes reads vs. writes mysterious
	// Use getter/setter. Conductivity only needs getter.
	// return secondary bool value in place of nil
	const efficiency = 0.9
	heat := heatsim.NewHeatGrid(efficiency, temperature, setTemperature, getConductivity)
	return &JobTransportNetwork{
		city:     city,
		HeatGrid: heat,
	}
}
