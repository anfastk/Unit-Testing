package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogoutHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}