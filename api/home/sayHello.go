package home

import (
	"batchLog/api/response"
	"time"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	requestedTime := time.Now().UTC()

	data := "hello"

	response.Success[string](c,requestedTime,data)
}