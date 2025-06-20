package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	// rooms table
	Room struct {
		ID          uuid.UUID `db:"id"`
		PlaceMax    int       `db:"place_max"`
		CenterPoint string    `db:"center_point"`
		Radius      float64   `db:"radius"`
		CreateAt    time.Time `db:"create_at"`
	}

	CreateRoomParams struct {
		PlaceMax    int     `db:"place_max"`
		CenterPoint string  `db:"center_point"`
		Radius      float64 `db:"radius"`
	}
)

func (r *Repository) GetRooms(ctx context.Context) ([]*Room, error) {
	rooms := []*Room{}
	if err := r.db.SelectContext(ctx, &rooms, "SELECT * FROM rooms"); err != nil {
		return nil, fmt.Errorf("select rooms: %w", err)
	}

	return rooms, nil
}

func (r *Repository) CreateRoom(ctx context.Context, params CreateRoomParams) (uuid.UUID, error) {
	roomID := uuid.New()
	now := time.Now().UTC()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO rooms (id, place_max, center_point, radius, create_at) VALUES (?, ?, ?, ?, ?)", roomID, params.PlaceMax, params.CenterPoint, params.Radius, now); err != nil {
		return uuid.Nil, fmt.Errorf("insert room: %w", err)
	}

	return roomID, nil
}

func (r *Repository) GetRoom(ctx context.Context, roomID uuid.UUID) (*Room, error) {
	room := &Room{}
	if err := r.db.GetContext(ctx, room, "SELECT * FROM rooms WHERE id = ?", roomID); err != nil {
		return nil, fmt.Errorf("select room: %w", err)
	}

	return room, nil
}
