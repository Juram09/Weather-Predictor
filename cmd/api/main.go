package main

import (
	"database/sql"
	"fmt"
	"github.com/Juram09/Weather-Predictor/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	db, _ := initializeDB()
	/*
		if err != nil {
			db = nil
			panic(err)
		}
	*/
	eng := gin.Default()
	router := http.InitRouter(eng, db)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}

func initializeDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbName)
	return sql.Open("mysql", connectionString)
}
