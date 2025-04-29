package routes

import (
	"AssessmentTest/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/auth/login", handlers.Login)

		v1.GET("/accounts/:accountId", handlers.GetAccount)
	}
}
