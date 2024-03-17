package shared

import (
	"net/http"

	"hiyoko-echo/pkg/logging/file"

	"github.com/labstack/echo/v4"
)

const (
	indent = "\t"
)

func ResponseOK(c echo.Context, data interface{}) error {
	if err := c.JSONPretty(http.StatusOK, data, indent); err != nil {
		logger.Error("OK response JSON conversion failed", "error", err, "data", data)
		return err
	}
	return nil
}

func ResponseCreate(c echo.Context, data interface{}) error {
	if err := c.JSONPretty(http.StatusCreated, data, indent); err != nil {
		logger.Error("Created response JSON conversion failed", "error", err, "data", data)
		return err
	}
	return nil
}

func ResponseNoContent(c echo.Context) error {
	if err := c.NoContent(http.StatusNoContent); err != nil {
		logger.Error("NoContent response JSON conversion failed", "error", err)
		return err
	}
	return nil
}

func ResponseBadRequest(c echo.Context, code string) error {
	if err := c.JSONPretty(http.StatusBadRequest, ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusBadRequest),
	}, indent); err != nil {
		logger.Error("BadRequest response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}

func ResponseNotFound(c echo.Context, code string) error {
	if err := c.JSONPretty(http.StatusNotFound, ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusNotFound),
	}, indent); err != nil {
		logger.Error("NotFound response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}
