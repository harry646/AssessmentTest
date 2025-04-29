package handlers

import (
	"net/http"
	"time"

	reqAPI "AssessmentTest/restapi/request"
	resAPI "AssessmentTest/restapi/response"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req reqAPI.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Dummy auth check
	if req.Username == "user" && req.Password == "pass" {
		c.JSON(http.StatusOK, resAPI.LoginResponse{
			Token:     "dummy-jwt-token",
			ExpiresIn: int(time.Hour.Seconds()),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}
