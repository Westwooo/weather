package weather

import "math/rand"

type RandWeather struct {
	Lat, Lon float64
}

func (r RandWeather) IsRainingAt() bool {
	return IsRainingFromWeatherCode(rand.Intn(100))
}

func (r RandWeather) CurrentTempAt() float64 {
	return rand.Float64() * 60
}

func (r RandWeather) WindSpeedAt() float64 {
	return rand.Float64() * 20
}
