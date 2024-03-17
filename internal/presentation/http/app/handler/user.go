package handler

import (
	"net/http"

	"hiyoko-echo/internal/application/usecase"
	"hiyoko-echo/internal/pkg/ent/util"
	"hiyoko-echo/internal/presentation/http/app/oapi"
	"hiyoko-echo/internal/shared"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUser(ctx echo.Context, id string) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) GetUser(c echo.Context, id string) error {
	ctx := c.Request().Context()

	user, err := h.UserUseCase.GetUser(ctx, util.ID(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, oapi.DefaultErrorSchema{
			Code:     shared.NoneCode,
			Messages: shared.GetErrorMessage(http.StatusNotFound),
		})
	}

	return c.JSON(http.StatusOK, user)
}
