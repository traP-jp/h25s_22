package handler

import (
	"backend/internal/repository"
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)
type GetVote struct{
	RoomID uuid.UUID `json:"roomID"`
	Name string	`json:"name"`
	TimeID []uuid.UUID	`json:"timeId"`
	PlaceData []repository.PlaceData `json:"placeData"`

}
func (h *Handler) PostVote(c echo.Context) error {
	getVote := GetVote{}
	err := c.Bind(getVote)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,"fatal err: %s" , err)
	}

	userID,err := h.repo.CreateUser(context.Background(), repository.CreateUserParams{
		RoomID: getVote.RoomID,
		Name: getVote.Name,

	} )
	timeVotes := lo.Map(getVote.TimeID, func(x uuid.UUID, index int)repository.CreateTimeVoteParams{
		return repository.CreateTimeVoteParams{
			UserID: userID,
			TimeID: x,
		}	
	})
	h.repo.CreateTimeVotes(context.Background(), timeVotes)

	placeVotes := lo.Map(getVote.placeData, func(x repository.PlaceData, index int)repository.CreatePlaceVoteParams{
		return repository.CreatePlaceVoteParams{
			UserID: userID,
			PlaceID:x.ID,
			Rank:x.Rank,

		}
	})


	h.repo.CreatePlaceVotes(context.Background(),placeVotes)
	
	

	return c.JSON(http.StatusOK, userID)
}