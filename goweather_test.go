package goweather_test

import (
	"fmt"
	"github.com/innotechdevops/goweather"
	"github.com/innotechdevops/openmeteo"
	"github.com/innotechdevops/openweather"
	"testing"
)

func TestGoWeather_GetFromOpenMeteo(t *testing.T) {
	param := goweather.OpenMeteoParameter{
		Parameter: openmeteo.Parameter{
			Latitude:  openmeteo.Float32(13.7248785),
			Longitude: openmeteo.Float32(100.4683022),
			Hourly: &[]string{
				openmeteo.HourlyTemperature2m,
				openmeteo.HourlyRelativeHumidity2m,
				openmeteo.HourlyWindSpeed10m,
			},
			CurrentWeather: openmeteo.Bool(true),
		},
	}

	o := goweather.New()
	resp, err := o.GetFromOpenMeteo(param)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(resp)
}

func TestGoWeather_GetFromOpenWeather(t *testing.T) {
	param := goweather.OpenWeatherParameter{
		Parameter: openweather.Parameter{
			Latitude:  openweather.Float32(33.44),
			Longitude: openweather.Float32(-94.04),
			Exclude: &[]string{
				openweather.ExcludeHourly,
				openweather.ExcludeDaily,
			},
			AppID: openweather.String("abc"),
		},
	}

	o := goweather.New()
	resp, err := o.GetFromOpenWeather(param)

	if err.Error() != "Unauthorized" {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(resp)
}
