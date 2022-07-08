package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Pong Pong Pong"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", Ping)
	_ = r.Run("0.0.0.0:8888")
}
