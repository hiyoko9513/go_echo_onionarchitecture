package handler

import (
	"hiyoko-echo/internal/application/usecase"
	"hiyoko-echo/internal/pkg/ent"
	"hiyoko-echo/internal/pkg/ent/util"
	"hiyoko-echo/internal/shared"

	"github.com/labstack/echo/v4"
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

	users, err := h.UserUseCase.GetUsers(ctx)
	if err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	return shared.ResponseOK(c, users)
}

func (h *userHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.UserUseCase.GetUser(ctx, util.ID(id))
	if err != nil {
		return shared.ResponseNotFound(c, shared.NoneCode)
	}

	return shared.ResponseOK(c, user)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := &ent.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	user, err := h.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	return shared.ResponseCreate(c, user)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.UserUseCase.UpdateUser(ctx, util.ID(id))
	if err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	return shared.ResponseOK(c, user)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if err := h.UserUseCase.DeleteUser(ctx, util.ID(id)); err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	return shared.ResponseNoContent(c)
}
