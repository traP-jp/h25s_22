package handler

import (
	"backend/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) SetupRoutes(api *echo.Group) {
	// api is "/api/v1"
	// ping API
	pingAPI := api.Group("/ping")
	{
		pingAPI.GET("", h.Ping)
	}

	roomAPI := api.Group("/room")
	{
		roomAPI.POST("/create", h.CreateRoom)
		roomAPI.GET("/:roomID", h.GetRoom)
		roomAPI.POST("/vote", h.PostVote)
		roomAPI.POST("/vote/change", h.ChangeVote)
		roomAPI.GET("/vote/result/:roomID", h.GetResult)

	}
	placeAPI := api.Group("/place")
	{
		placeAPI.GET("/nearSearch", h.GetSearch)
		placeAPI.GET("/photo", h.GetPhoto)
		placeAPI.GET("/:placeId", h.GetPlace)
	}
}
