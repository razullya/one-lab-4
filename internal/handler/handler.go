package handler

import (
	"lab4/internal/service"

	"github.com/labstack/echo"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) InitRoutes() {
	e := echo.New()

	e.POST("/users/create", h.createUser)
	e.GET("/users/:id", h.getUserById)
	e.PUT("/users/update", h.updateUser)
	e.DELETE("/users/:id", h.deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
