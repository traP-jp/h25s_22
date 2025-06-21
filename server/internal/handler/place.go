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
type GetDetail struct{
	name string
	rate float32
	priceLevel int
	address string


}
func (h *Handler) GetPlace(c echo.Context) error {
	client, err := maps.NewClient(maps.WithAPIKey("api_key"))
	placeId := c.Param(":placeId")
	detailReq := &maps.PlaceDetailsRequest{
		PlaceID:  placeId,
		Language: "ja",
	}
	detail, errors := client.PlaceDetails(context.Background(),detailReq)
	if err != nil || errors != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "fatal err: %s", err)
	}
	getDetail := GetDetail{
		name: detail.Name,
		rate: detail.Rating,
		priceLevel: detail.PriceLevel,
		address: detail.FormattedAddress,
	}
	return c.JSON(http.StatusOK, getDetail)
}
