package weather

type RealWeather struct {
	Lat, Lon float64
}

func (r RealWeather) IsRainingAt() bool {
	var res Result
	GetCurrentWeather(&res, r.Lat, r.Lon)
	return IsRainingFromWeatherCode(res.CurrentWeather.Weathercode)
}

func (r RealWeather) CurrentTempAt() float64 {
	var res Result
	GetCurrentWeather(&res, r.Lat, r.Lon)
	return res.CurrentWeather.Temperature
}

func (r RealWeather) WindSpeedAt() float64 {
	var res Result
	GetCurrentWeather(&res, r.Lat, r.Lon)
	return res.CurrentWeather.Windspeed
}
