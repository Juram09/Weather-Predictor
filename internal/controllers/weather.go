package controllers

import (
	"github.com/Juram09/Weather-Predictor/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	IWeather interface {
		GetDrought() gin.HandlerFunc
		GetRainy() gin.HandlerFunc
		GetOptimal() gin.HandlerFunc
		GetWeather() gin.HandlerFunc
	}
	weather struct {
		weatherService service.IWeather
	}
)

func NewWeather(weatherService service.IWeather) IWeather {
	return &weather{
		weatherService: weatherService,
	}
}

func (w *weather) GetDrought() gin.HandlerFunc {
	return func(c *gin.Context) {
		years, err := strconv.Atoi(c.Query("years"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid years",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"years":    years,
			"droughts": w.weatherService.GetDrought(years),
		})
	}
}

func (w *weather) GetRainy() gin.HandlerFunc {
	return func(c *gin.Context) {
		years, err := strconv.Atoi(c.Query("years"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid years",
			})
			return
		}
		rains, maxRain := w.weatherService.GetRainy(years)
		c.JSON(http.StatusOK, gin.H{
			"years":     years,
			"rains":     rains,
			"rain_peak": maxRain,
		})
	}
}

func (w *weather) GetOptimal() gin.HandlerFunc {
	return func(c *gin.Context) {
		years, err := strconv.Atoi(c.Query("years"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid years",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"years":   years,
			"optimal": w.weatherService.GetOptimal(years),
		})
	}
}

func (w *weather) GetWeather() gin.HandlerFunc {
	return func(c *gin.Context) {
		day, err := strconv.Atoi(c.Query("day"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid day",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"day":     day,
			"weather": w.weatherService.GetWeather(day),
		})
	}
}
