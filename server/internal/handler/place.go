package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

func (h *Handler) GetSearch(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}



func (h *Handler) GetPhoto(c echo.Context) error {
	client, err := maps.NewClient(maps.WithAPIKey("api_key"))
	photoReq := &maps.PlacePhotoRequest{}
	errors := c.Bind(photoReq)
	photo, err := client.PlacePhoto(context.Background(), photoReq)
	if err != nil || errors != nil {
		return echo.NewHTTPError(http.StatusBadRequest,"fatal err: %s" , err)
	} 
	return c.JSON(http.StatusOK,photo)
}
func (h *Handler) GetPlace(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
