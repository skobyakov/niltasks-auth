package controller

import (
	"net/http"
	"niltasks-auth/internal/models"

	"github.com/gin-gonic/gin"
)

type Service interface {
	TryAuth(req *models.AuthRequest) (string, error)
}

type Controller struct {
	service Service
}

func New(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) HandleLogin(ctx *gin.Context) {
	var req models.AuthRequest

	if ctx.ShouldBindJSON(&req) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session, err := c.service.TryAuth(&req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res := models.AuthResponse{Session: session}

	ctx.JSON(200, res)
}
