package handler

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"googlemaps.github.io/maps"
)
type AskSearch struct {
	Location  maps.LatLng `json:"location"`
	Name string `json:"name"`
	PlaceID string `json:"placeID"`
	Radius 	int `json:"radius"`
	Language string `json:"language"`
	Type string `json:"type"`
	MaxSize int `json:"maxSize"`
	Address string `json:"address"`
}

type GetSearch struct{
	Name string `json:"name"`
	PlaceID string `json:"placeID"`
	Location maps.LatLng `json:"location"`
	PhotoReference maps.Photo `json:"photoReference"`
	PriceLevel int `json:"priceLevel"`
	Address string `json:"address"`
}

func (h *Handler) GetSearch(c echo.Context) error {
	api_key, ok := os.LookupEnv("GOOGLE_API_KEY")
	if !ok  {
		return c.String(http.StatusInternalServerError,"Not Fund APIKey")
	}
	client, err := maps.NewClient(maps.WithAPIKey(api_key))

	nearRequest := &AskSearch{
		Name: c.QueryParam("name"),
		Type: c.QueryParam("type"),
		Language: c.QueryParam("language"),
	}

	if radiusStr := c.QueryParam("radius"); radiusStr != "" {
		if radius, err := strconv.Atoi(radiusStr); err == nil {
			nearRequest.Radius = radius
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid radius parameter")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "radius parameter is required")
	}

	if maxSizeStr := c.QueryParam("maxSize"); maxSizeStr != "" {
		if maxSize, err := strconv.Atoi(maxSizeStr); err == nil {
			nearRequest.MaxSize = maxSize
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid maxSize parameter")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "maxSize parameter is required")
	}

	if nearRequest.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name parameter is required")
	}

	nameSearch := maps.TextSearchRequest{
		Query: nearRequest.Name,
	}
	response,err := client.TextSearch(context.Background(),&nameSearch)
	if err != nil {
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
	if err != nil {
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
	api_key, ok := os.LookupEnv("GOOGLE_API_KEY")
	if !ok  {
		return c.String(http.StatusInternalServerError,"Not Fund APIKey")
	}
	client, err := maps.NewClient(maps.WithAPIKey(api_key))

	photoReq := &maps.PlacePhotoRequest{
		PhotoReference: c.QueryParam("photoReference"),
	}

	if maxWidthStr := c.QueryParam("maxWidth"); maxWidthStr != "" {
		if maxWidth, err := strconv.Atoi(maxWidthStr); err == nil {
			photoReq.MaxWidth = uint(maxWidth)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid maxWidth parameter")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "maxWidth parameter is required")
	}

	if maxHeightStr := c.QueryParam("maxHeight"); maxHeightStr != "" {
		if maxHeight, err := strconv.Atoi(maxHeightStr); err == nil {
			photoReq.MaxHeight = uint(maxHeight)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid maxHeight parameter")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "maxHeight parameter is required")
	}

	if photoReq.PhotoReference == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "photoReference parameter is required")
	}

	photo, err := client.PlacePhoto(context.Background(), photoReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,"fatal err: %s" , err)
	}
	return c.JSON(http.StatusOK,photo)
}
type GetDetail struct{
	Name string `json:"name"`
	Rate float32 `json:"rate"`
	PriceLevel int `json:"pricelevel"`
	Address string `json:"address"`


}
func (h *Handler) GetPlace(c echo.Context) error {
	api_key, ok := os.LookupEnv("GOOGLE_API_KEY")
	if !ok  {
		return c.String(http.StatusInternalServerError,"Not Fund APIKey")
	}
	client, err := maps.NewClient(maps.WithAPIKey(api_key))
	placeId := c.Param("placeId")
	detailReq := &maps.PlaceDetailsRequest{
		PlaceID:  placeId,
		Language: "ja",
	}
	detail, errors := client.PlaceDetails(context.Background(),detailReq)
	if err != nil || errors != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "fatal err: %s", err)
	}
	getDetail := GetDetail{
		Name: detail.Name,
		Rate: detail.Rating,
		PriceLevel: detail.PriceLevel,
		Address: detail.FormattedAddress,
	}
	return c.JSON(http.StatusOK, getDetail)
}
