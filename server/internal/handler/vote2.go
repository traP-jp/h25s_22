package handler

import (
	"backend/internal/repository"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	ChangeVoteRequest struct {
		UserID     uuid.UUID `json:"user_id"`
		TimeOption uuid.UUID `json:"time_option"`
		PlaceID    uuid.UUID `json:"place_id"`
	}
)

func (h *Handler) ChangeVote(c echo.Context) error {
	req := &repository.UpdateVoteReq{}
	err := c.Bind(req)
	if err != nil ||
		req.UserID == uuid.Nil ||
		req.TimeVotes == nil ||
		req.PlaceVotes == nil ||
		len(req.PlaceVotes) == 0 ||
		len(req.TimeVotes) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	if err := h.repo.UpdateVote(c.Request().Context(), req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Failed to update vote: %w", err))
	}
	return c.String(http.StatusOK, "success update vote")
}

func (h *Handler) GetResult(c echo.Context) error {
	// roomID := c.Param("roomId")

	return c.String(http.StatusOK, "pong")
}
