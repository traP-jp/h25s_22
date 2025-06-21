package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"googlemaps.github.io/maps"
)
type AskSearch struct {
	Location  maps.LatLng `json:"location"`
	Name string `json:"name"`
	PlaceID string `json:"placeID"`
	Radius 	float64 `json:"radius"`
	Language string `json:"language"`
	Type string `json:"type"`
	MaxSize int `json:"maxSize"`
	Address string `json:"address"`
	
}

type GetSearch struct{
	Name string `json:"name"`
	PlaceID string `json:"placeID`
	Location maps.LatLng `json:"location"`
	PhotoReference maps.Photo `json:"photoRefarence"`
	PriceLevel int `json:"priceLevel"`
	Address string `json:"address"`

}

func (h *Handler) GetSearch(c echo.Context) error {
	client, err := maps.NewClient(maps.WithAPIKey("api_key"))
	nearRequest := &AskSearch{}
	errors := c.Bind(nearRequest)
	nameSearch := maps.TextSearchRequest{
		Query: nearRequest.Name,
	}
	response,erro := client.TextSearch(context.Background(),&nameSearch)
	if erro != nil {		
		return echo.NewHTTPError(http.StatusBadRequest,"fatal err: %s" , err)
	}
	nearRequest.Location = maps.LatLng{
		Lat: response.Results[0].Geometry.Location.Lat,
		Lng: response.Results[0].Geometry.Location.Lng,
	}
	nearRequestMap := maps.NearbySearchRequest{
		Name: nearRequest.Name,
		Type: maps.PlaceType(nearRequest.Type),
		Location: &nearRequest.Location,
		Radius: uint(nearRequest.Radius),
		Language: nearRequest.Language,
	}



	route, err := client.NearbySearch(context.Background(),&nearRequestMap)
	if err != nil || errors != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "fatal err: %s", err)
	}
	search := route.Results[:nearRequest.MaxSize]
	getSearch := lo.Map(search, func(x maps.PlacesSearchResult, index int) GetSearch {
		return GetSearch{
			Name: x.Name,
			PlaceID: x.PlaceID,
			Location:x.Geometry.Location,
			PhotoReference: x.Photos[0],
			PriceLevel: x.PriceLevel,
			Address: x.FormattedAddress,

		}
	})


	return c.JSON(http.StatusOK, getSearch)
}



func (h *Handler) GetPhoto(c echo.Context) error {
	client, err := maps.NewClient(maps.WithAPIKey("api_key"))
	photoReq := &maps.PlacePhotoRequest{}
	errors := c.Bind(photoReq)
	photo, err := client.PlacePhoto(context.	Background(), photoReq)
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
