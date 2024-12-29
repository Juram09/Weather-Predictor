package service

import (
	"github.com/Juram09/Weather-Predictor/internal/defines"
	"github.com/Juram09/Weather-Predictor/internal/domain/entities"
	"github.com/Juram09/Weather-Predictor/internal/repository"
	"math"
)

type (
	IWeather interface {
		GetDrought(int) int
		GetRainy(int) (int, int)
		GetOptimal(int) int
		GetWeather(int) string
	}
	weather struct {
		weatherRepository repository.IWeather
	}
)

func NewWeather(weatherRepository repository.IWeather) IWeather {
	return &weather{
		weatherRepository: weatherRepository,
	}
}

func (w *weather) GetDrought(years int) int {
	droughtPeriods := 0
	for i := 0; i <= defines.YearDuration*years; i++ {
		if w.GetWeather(i) == defines.Drought {
			droughtPeriods++
		}
	}
	return droughtPeriods
}

func (w *weather) GetRainy(years int) (int, int) {
	rainyPeriods := 0
	maxRainDay := 0
	maxRainPerimeter := 0.0
	for i := 0; i <= defines.YearDuration*years; i++ {
		if w.GetWeather(i) == defines.Rain {
			rainyPeriods++
			perimeter := trianglePerimeter(i)
			if perimeter > maxRainPerimeter {
				maxRainPerimeter = perimeter
				maxRainDay = i
			}
		}
	}
	return rainyPeriods, maxRainDay
}

func (w *weather) GetOptimal(years int) int {
	optimalPeriods := 0
	for i := 0; i <= defines.YearDuration*years; i++ {
		if w.GetWeather(i) == defines.Optimal {
			optimalPeriods++
		}
	}
	return optimalPeriods
}

func (w *weather) GetWeather(day int) string {
	conditions, err := w.weatherRepository.GetWeather(day)
	if err != nil || conditions == "" {
		weatherConditions := calculateWeather(day)
		w.weatherRepository.SaveWeather(day, weatherConditions)
		return weatherConditions
	}
	return conditions
}

func calculateWeather(day int) string {
	x1, y1 := calculatePosition(defines.Ferengi, day)
	x2, y2 := calculatePosition(defines.Vulcano, day)
	x3, y3 := calculatePosition(defines.Betazoide, day)

	if areAligned(x1, y1, x2, y2, x3, y3) {
		if areAligned(x1, y1, x2, y2, 0, 0) {
			return defines.Drought
		} else {
			return defines.Optimal
		}
	} else {
		if containsSun(x1, y1, x2, y2, x3, y3) {
			return defines.Rain
		}
	}
	return defines.Normal
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func toOneDecimal(value float64) float64 {
	return math.Round(value*10) / 10
}

func calculatePosition(planet entities.Planet, day int) (float64, float64) {
	angle := float64(day) * planet.Velocity
	if !planet.Clockwise {
		angle = -angle
	}
	angle = math.Mod(angle, 360)
	radians := degreesToRadians(angle)
	return toOneDecimal(planet.Distance * math.Cos(radians)), toOneDecimal(planet.Distance * math.Sin(radians))
}

func areAligned(x1, y1, x2, y2, x3, y3 float64) bool {
	return math.Abs((y2-y1)*(x3-x2)-(x2-x1)*(y3-y2)) < 50000
}

func containsSun(x1, y1, x2, y2, x3, y3 float64) bool {
	area := math.Abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)) / 2

	area1 := math.Abs(0*(y2-y3)+x2*(y3-0)+x3*(0-y2)) / 2
	area2 := math.Abs(x1*(0-y3)+0*(y3-y1)+x3*(y1-0)) / 2
	area3 := math.Abs(x1*(y2-0)+x2*(0-y1)+0*(y1-y2)) / 2

	return math.Abs(area-(area1+area2+area3)) < 1
}

func trianglePerimeter(day int) float64 {
	x1, y1 := calculatePosition(defines.Ferengi, int(day))
	x2, y2 := calculatePosition(defines.Vulcano, int(day))
	x3, y3 := calculatePosition(defines.Betazoide, int(day))
	return distance(x1, y1, x2, y2) + distance(x2, y2, x3, y3) + distance(x1, y1, x3, y3)
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
