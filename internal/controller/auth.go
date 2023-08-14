package controller

import "github.com/gin-gonic/gin"

type Service interface {
	TryAuth()
}

type Controller struct {
	service Service
}

func New(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) HandleLogin(ctx *gin.Context) {
	c.service.TryAuth()

	ctx.JSON(200, gin.H{
		"session_id": "123",
	})
}
