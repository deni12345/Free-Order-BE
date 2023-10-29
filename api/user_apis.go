package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, map[string]string{"name": "Liem"})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, map[string]string{"name": "Liem"})
}
