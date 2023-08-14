package app

import (
	"fmt"
	"niltasks-auth/config"
	"niltasks-auth/internal/controller"
	"niltasks-auth/internal/repository"
	"niltasks-auth/internal/service"

	"github.com/gin-gonic/gin"
)

func Serve() {
	cfg := config.MustLoad()

	repo := repository.New()
	service := service.New(repo)
	controller := controller.New(service)

	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.POST("/login", controller.HandleLogin)

	host := cfg.Server.Host + ":" + cfg.Server.Port
	fmt.Printf("Starting server on %s\n", host)
	err := r.Run(host)
	if err != nil {
		panic(err)
	}
}
