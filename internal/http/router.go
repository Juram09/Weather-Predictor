package http

import (
	"database/sql"
	"github.com/Juram09/Weather-Predictor/internal/controllers"
	"github.com/Juram09/Weather-Predictor/internal/repository"
	"github.com/Juram09/Weather-Predictor/internal/service"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}
type router struct {
	eng *gin.Engine
	db  *sql.DB
}

func (r *router) MapRoutes() {
	r.addSystemPaths()
	r.buildRoutes()
}

func (r *router) buildRoutes() {
	weatherRepository := repository.NewWeatherRepository(r.db)
	weatherService := service.NewWeather(weatherRepository)
	weatherController := controllers.NewWeather(weatherService)
	r.eng.GET("/weather", weatherController.GetWeather())
	r.eng.GET("/weather/drought", weatherController.GetDrought())
	r.eng.GET("/weather/rainy", weatherController.GetRainy())
	r.eng.GET("/weather/optimal", weatherController.GetOptimal())
}

func (r *router) addSystemPaths() {
	r.eng.GET("/ping", controllers.Ping())
}
