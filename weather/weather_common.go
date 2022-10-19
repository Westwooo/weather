package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Timezone       string         `json:"timezone"`
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type CurrentWeather struct {
	Weathercode int     `json:"weathercode"`
	Temperature float64 `json:"temperature"`
	Windspeed   float64 `json:"windspeed"`
}

type WeatherChecker interface {
	IsRainingAt() bool
	CurrentTempAt() float64
	WindSpeedAt() float64
}

func GetWeatherBasics(w WeatherChecker) {
	print("    Location ")
	fmt.Println(w)
	fmt.Println("    Is Raining:", w.IsRainingAt())
	fmt.Println("    Current temp:", w.CurrentTempAt())
	fmt.Println("    Wind speed:", w.WindSpeedAt())
}

func IsRainingFromWeatherCode(i int) bool {
	if 51 <= i && i <= 67 || 80 <= i && i <= 82 {
		return true
	}
	return false
}

func GetCurrentWeather(res *Result, lat float64, lon float64) {
	latString := fmt.Sprintf("%.2f", lat)
	lonString := fmt.Sprintf("%.2f", lon)

	getString := "https://api.open-meteo.com/v1/forecast?latitude=" + latString + "&longitude=" +
		lonString + "&current_weather=true"

	resp, err := http.Get(getString)

	if err != nil {
		print("Could not get the current weather")
	}

	body, err := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
}
