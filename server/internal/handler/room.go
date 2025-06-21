package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateRoom(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (h *Handler) GetRoom(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (h *Handler) PostVote(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (h *Handler) ChangeVote(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (h *Handler) GetResult(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

