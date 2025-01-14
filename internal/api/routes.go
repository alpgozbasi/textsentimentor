package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", HealthCheckHandler)
	r.POST("/analyze", AnalyzeHandler)
}
