package handlers

import (
	"net/http"

	reqAPI "AssessmentTest/restapi/request"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	accountId := c.Param("accountId")

	// Dummy data
	if accountId == "12345" {
		account := reqAPI.AccountInfo{
			AccountID:   "12345",
			Balance:     1000000.50,
			Currency:    "IDR",
			AccountType: "Savings",
		}
		c.JSON(http.StatusOK, account)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
	}
}
