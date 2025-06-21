package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetSearch(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (h *Handler) GetPhoto(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (h *Handler) GetPlace(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}