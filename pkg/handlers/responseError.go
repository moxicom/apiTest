package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type responseError struct {
	Message string `json:"message"`
}

func newResponseError(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, responseError{message})
}
