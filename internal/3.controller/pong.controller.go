package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "thaitran")
	uid := c.Query("uid")
	// c.ShouldBindJSON()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong..pong.." + name,
		"uid":     uid,
	})
}
