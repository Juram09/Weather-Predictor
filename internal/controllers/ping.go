package controllers

import (
	"github.com/Juram09/Weather-Predictor/internal/utils"
	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Success(c, 200, "Pong")
	}
}
