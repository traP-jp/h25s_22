package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// places table
	Place struct {
		ID            uuid.UUID `db:"id"`
		RoomID        uuid.UUID `db:"room_id"`
		GooglePlaceID uuid.UUID `db:"google_place_id"`
	}

	CreatePlaceParams struct {
		RoomID        uuid.UUID `db:"room_id"`
		GooglePlaceID uuid.UUID `db:"google_place_id"`
	}
)

func (r *Repository) GetPlaces(ctx context.Context) ([]*Place, error) {
	places := []*Place{}
	if err := r.db.SelectContext(ctx, &places, "SELECT * FROM places"); err != nil {
		return nil, fmt.Errorf("select places: %w", err)
	}

	return places, nil
}

func (r *Repository) GetPlacesByRoomID(ctx context.Context, roomID uuid.UUID) ([]*Place, error) {
	places := []*Place{}
	if err := r.db.SelectContext(ctx, &places, "SELECT * FROM places WHERE room_id = ?", roomID); err != nil {
		return nil, fmt.Errorf("select places by roomID: %w", err)
	}
	return places, nil
}

func (r *Repository) CreatePlace(ctx context.Context, params CreatePlaceParams) (uuid.UUID, error) {
	PlaceID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO places (id, room_id, google_place_id) VALUES (?, ?, ?)", PlaceID, params.RoomID, params.GooglePlaceID); err != nil {
		return uuid.Nil, fmt.Errorf("insert place: %w", err)
	}

	return PlaceID, nil
}

func (r *Repository) GetPlace(ctx context.Context, PlaceID uuid.UUID) (*Place, error) {
	place := &Place{}
	if err := r.db.GetContext(ctx, place, "SELECT * FROM places WHERE id = ?", PlaceID); err != nil {
		return nil, fmt.Errorf("select place: %w", err)
	}

	return place, nil
}
