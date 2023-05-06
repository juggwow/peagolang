package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"message": err.Error(),
	})
}
