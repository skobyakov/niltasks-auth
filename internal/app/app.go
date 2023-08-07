package app

import (
	"niltasks-auth/internal/controller"
	"niltasks-auth/internal/repository"
	"niltasks-auth/internal/service"
)

func Serve() {
	repo := repository.New()
	service := service.New(repo)
	controller := controller.New(service)

	controller.HandleAuth()
}
