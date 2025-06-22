package handler

import (
	"backend/internal/repository"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type GetVote struct {
	RoomID    uuid.UUID              `json:"roomID"`
	Name      string                 `json:"name"`
	TimeID    []uuid.UUID            `json:"timeId"`
	PlaceData []repository.PlaceData `json:"placeData"`
}
type ChangeVoteRequest struct {
	UserID     uuid.UUID `json:"user_id"`
	TimeOption uuid.UUID `json:"time_option"`
	PlaceID    uuid.UUID `json:"place_id"`
}

func (h *Handler) PostVote(c echo.Context) error {
	getVote := GetVote{}
	err := c.Bind(getVote)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "fatal err: %s", err)
	}

	userID, err := h.repo.CreateUser(context.Background(), repository.CreateUserParams{
		RoomID: getVote.RoomID,
		Name:   getVote.Name,
	})
	timeVotes := lo.Map(getVote.TimeID, func(x uuid.UUID, index int) repository.CreateTimeVoteParams {
		return repository.CreateTimeVoteParams{
			UserID: userID,
			TimeID: x,
		}
	})
	h.repo.CreateTimeVotes(context.Background(), timeVotes)

	placeVotes := lo.Map(getVote.PlaceData, func(x repository.PlaceData, index int) repository.CreatePlaceVoteParams {
		return repository.CreatePlaceVoteParams{
			UserID:  userID,
			PlaceID: x.ID,
			Rank:    x.Rank,
		}
	})

	h.repo.CreatePlaceVotes(context.Background(), placeVotes)

	return c.JSON(http.StatusOK, userID)
}

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
	roomIDStr := c.Param("roomID")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("fatal err: %s", err))
	}
	roomResult, err := h.repo.GetRoomResult(context.Background(), roomID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("fatal err: %s", err))
	}

	return c.JSON(http.StatusOK, roomResult)
}
