package controller

type Service interface {
	TryAuth()
}

type Controller struct {
	service Service
}

func New(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) HandleAuth() {
	c.service.TryAuth()
}
