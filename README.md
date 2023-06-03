# goweather

Combine API [Open-Meteo](https://github.com/innotechdevops/openmeteo) and [Open Weather](https://github.com/innotechdevops/openweather)

## Install

```shell
go get github.com/innotechdevops/goweather
```

## How to use

- Open-Meteo

```go
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
```

- Open Weather

```go
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
```