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
	PlaceID []uuid.UUID	`json:"placeId"`
	Rank int	`json:"rank"`

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
	lo.Map(getVote.TimeID, func(x uuid.UUID, index int)error{
			h.repo.CreateTimeVote(context.Background(), repository.CreateTimeVoteParams{
			UserID: userID,
			TimeID: x,
		})
		return echo.NewHTTPError(http.StatusBadRequest,"fatal err: %s" , err)
	})

	lo.Map(getVote.PlaceID, func(x uuid.UUID, index int)error{
		h.repo.CreatePlaceVote(context.Background(),repository.CreatePlaceVoteParams{
		UserID : userID,
    	PlaceID : x,
   		Rank    : getVote.Rank,  
	})
		return echo.NewHTTPError(http.StatusBadRequest,"fatal err: %s" , err)
	})
	
	

	return c.JSON(http.StatusOK, userID)
}