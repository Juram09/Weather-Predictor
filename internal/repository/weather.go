package repository

import (
	"database/sql"
	"github.com/Juram09/Weather-Predictor/internal/defines"
	"log"
)

type (
	IWeather interface {
		GetWeather(int) (string, error)
		SaveWeather(int, string)
	}
	weatherRepository struct {
		db *sql.DB
	}
)

func NewWeatherRepository(db *sql.DB) IWeather {
	return &weatherRepository{db: db}
}

func (wr *weatherRepository) GetWeather(day int) (string, error) {
	if wr.db != nil {
		row := wr.db.QueryRow(defines.GetWeather, day)
		weather := new(string)
		err := row.Scan(weather)
		if err != nil {
			log.Println(err)
			return *weather, err
		}
		return *weather, nil
	}
	return "", nil
}
func (wr *weatherRepository) SaveWeather(day int, weather string) {
	if wr.db != nil {
		stmt, err := wr.db.Prepare(defines.SaveWeather)
		if err != nil {
			log.Println(err)
		}
		_, err = stmt.Exec(&day, &weather)
		if err != nil {
			log.Println(err)
		}
	}
}
