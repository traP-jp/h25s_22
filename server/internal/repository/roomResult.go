package repository

import (
	"context"
	"fmt"
	// "time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	//"github.com/pelletier/go-toml/query"
)

type PlaceResult struct{
	GooglePlaceID string    `db:"google_place_id"`
	PlaceName string 
	Point int 

}
type UserNameID struct{
	Name       string    `db:"name"`
	UserID     uuid.UUID `db:"id"`	
}
type TimeResult struct{
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
	user      []UserNameID
}
type RoomResult struct{
	place []PlaceResult
	time  []TimeResult
}


func (r *Repository) GetRoomResult(ctx context.Context, roomID uuid.UUID) (*RoomResult, error) {
	
	timeOptions := []*timeOption{}
	if err := r.db.SelectContext(ctx, &timeOptions, "SELECT * FROM timeOptions WHERE room_id = ?", roomID); err != nil {
		return nil, fmt.Errorf("select timeOptions by roomID: %w", err)
	}

	timeResults := make([]TimeResult, len(timeOptions))

	for i, timeOption := range timeOptions{
		userNameID := []UserNameID{}
		if err := r.db.SelectContext(ctx, &userNameID, "SELECT timeVotes.user_id, users.Name FROM timeOptions INNER JOIN timeVotes.user_id ON users.id  WHERE time_id = ?", timeOption.ID); err != nil {
			return nil, fmt.Errorf("select timeOptions by roomID: %w", err)
		}
		timeResults[i] = TimeResult{
			StartTime: timeOption.StartTime,
			EndTime: timeOption.EndTime,
			user : userNameID,
		}
	}
	places := []*Place{}

	if err := r.db.SelectContext(ctx, &places, "SELECT * FROM places WHERE room_id = ?", roomID); err != nil {
		return nil, fmt.Errorf("select place by roomID: %w", err)
	}

	placeResults := make([]PlaceResult, len(places))

	for i, place := range places{
		ranks := []int{}
		if err := r.db.SelectContext(ctx, &ranks, "SELECT rank FROM placeVotes WHERE place_id = ?", place.ID); err != nil {
			return nil, fmt.Errorf("select timeOptions by roomID: %w", err)
		}
		placeResults[i] = PlaceResult{
			GooglePlaceID: place.GooglePlaceID,
			Point: lo.Reduce(ranks, func(agg int, item int, _ int) int{
				 return agg + len(places) - item + 1
			}, 0),
		}
	}

	roomResults := &RoomResult{
		place: placeResults,
		time: timeResults,
	}

	return roomResults, nil
}