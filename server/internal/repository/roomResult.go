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
	GooglePlaceID string `db:"google_place_id" json:"googlePlaceID"`
	PlaceName     string `json:"placeName"`
	Point         int    `json:"point"`
}

type UserNameID struct{
	Name   string    `db:"name" json:"name"`
	UserID uuid.UUID `db:"id" json:"userID"`	
}

type TimeResult struct{
	StartTime string       `db:"start_time" json:"startTime"`
	EndTime   string       `db:"end_time" json:"endTime"`
	User      []UserNameID `json:"user"`
}

type RoomResult struct{
	Place []PlaceResult `json:"place"`
	Time  []TimeResult  `json:"time"`
}


func (r *Repository) GetRoomResult(ctx context.Context, roomID uuid.UUID) (*RoomResult, error) {
	
	timeOptions := []*timeOption{}
	if err := r.db.SelectContext(ctx, &timeOptions, "SELECT * FROM timeOptions WHERE room_id = ?", roomID); err != nil {
		return nil, fmt.Errorf("select timeOptions by roomID: %w", err)
	}

	timeResults := make([]TimeResult, len(timeOptions))

	for i, timeOption := range timeOptions{
		userNameID := []UserNameID{}
		if err := r.db.SelectContext(ctx, &userNameID, "SELECT users.id, users.name FROM timeVotes INNER JOIN users ON timeVotes.user_id = users.id WHERE timeVotes.time_id = ?", timeOption.ID); err != nil {
			return nil, fmt.Errorf("select users by timeID: %w", err)
		}
		timeResults[i] = TimeResult{
			StartTime: timeOption.StartTime,
			EndTime: timeOption.EndTime,
			User: userNameID,
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
			return nil, fmt.Errorf("select ranks by placeID: %w", err)
		}
		placeResults[i] = PlaceResult{
			GooglePlaceID: place.GooglePlaceID,
			PlaceName:     fmt.Sprintf("Place %s", place.GooglePlaceID), // TODO: 実際の場所名を取得
			Point: lo.Reduce(ranks, func(agg int, item int, _ int) int{
				 return agg + len(places) - item + 1
			}, 0),
		}
	}

	roomResults := &RoomResult{
		Place: placeResults,
		Time:  timeResults,
	}

	return roomResults, nil
}