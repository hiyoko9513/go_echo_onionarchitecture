package router

import (
	"hiyoko-echo/internal/presentation/http/app/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	api := e.Group("")
	v1 := api.Group("/v1")
	v1.POST("/users", h.CreateUser)
	v1.GET("/users", h.ListUsers)
	v1.GET("/users/:id", h.GetUser)

	//v1Guard := v1.Use(middleware.auth())
	//v1Guard.GET("/users/me", h.GetMe)
	//v1Guard.PUT("/users/:id", h.UpdateUser)
	//v1Guard.DELETE("/users/:id", h.DeleteUser)
}
