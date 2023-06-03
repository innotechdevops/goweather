package goweather

import (
	"github.com/innotechdevops/openmeteo"
	"github.com/innotechdevops/openweather"
)

type OpenMeteoParameter struct {
	openmeteo.Parameter
}

type OpenWeatherParameter struct {
	openweather.Parameter
}
