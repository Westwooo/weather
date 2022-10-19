package main

import (
	"fmt"

	"github.com/Westwooo/weather/weather"
)

func main() {
	lat, lon := 53.480759, -2.242631

	fmt.Printf("Real weather\n")
	real := weather.RealWeather{Lat: lat, Lon: lon}
	weather.GetWeatherBasics(real)

	fmt.Printf("Random weather\n")
	rand := weather.RandWeather{Lat: lat, Lon: lon}
	weather.GetWeatherBasics(rand)

	fmt.Printf("Random weather\n")
	weather.GetWeatherBasics(rand)
}
