package goweather

import (
	"github.com/innotechdevops/openmeteo"
	"github.com/innotechdevops/openweather"
)

type GoWeather interface {
	GetFromOpenMeteo(parameter OpenMeteoParameter) (string, error)
	GetFromOpenWeather(parameter OpenWeatherParameter) (string, error)
}

type goWeather struct {
	OpenMeteo   openmeteo.OpenMeteo
	OpenWeather openweather.OpenWeather
}

func (g *goWeather) GetFromOpenMeteo(parameter OpenMeteoParameter) (string, error) {
	return g.OpenMeteo.Execute(parameter.Parameter)
}

func (g *goWeather) GetFromOpenWeather(parameter OpenWeatherParameter) (string, error) {
	return g.OpenWeather.Execute(parameter.Parameter)
}

func New() GoWeather {
	return &goWeather{
		OpenMeteo:   openmeteo.New(),
		OpenWeather: openweather.New(),
	}
}
