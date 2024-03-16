package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"hiyoko-echo/pkg/mypubliclib/ent"
	"hiyoko-echo/pkg/mypubliclib/ent/util"
	"hiyoko-echo/usecase"
	"net/http"
)

type UserHandler interface {
	ListUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) ListUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := h.UserUseCase.GetUsers(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.GetUser(ctx, util.ID(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &ent.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.UpdateUser(ctx, util.ID(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.UserUseCase.DeleteUser(ctx, util.ID(id)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	return c.NoContent(http.StatusNoContent)
}
