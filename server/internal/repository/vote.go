package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	PlaceData struct {
		ID   uuid.UUID `db:"id"`
		Rank int       `db:"rank"`
	}

	CreateVoteReq struct {
		RoomID     uuid.UUID `db:"room_id"`
		Name       string    `db:"name"`
		TimeVotes  []uuid.UUID
		PlaceVotes []PlaceData
	}

	UpdateVoteReq struct {
		UserID     uuid.UUID `db:"user_id"`
		TimeVotes  []uuid.UUID
		PlaceVotes []PlaceData
	}
)

func (r *Repository) CreateVote(ctx context.Context, param CreateVoteReq) (uuid.UUID, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return uuid.Nil, fmt.Errorf("faild start transaction: %w", err)
	}

	userID := uuid.New()
	if _, err := tx.ExecContext(ctx, "INSERT INTO users (id, room_id, name) VALUES (?, ?, ?)", userID, param.RoomID, param.Name); err != nil {
		tx.Rollback()
		return uuid.Nil, fmt.Errorf("faild insert user: %w", err)
	}

	for _, timeVote := range param.TimeVotes {
		if _, err := tx.ExecContext(ctx, "INSERT INTO timeVotes (id, user_id, time_id) VALUES (?, ?, ?)", uuid.New(), userID, timeVote); err != nil {
			tx.Rollback()
			return uuid.Nil, fmt.Errorf("faild insert timeVote: %w", err)
		}
	}

	for _, placeVote := range param.PlaceVotes {
		if _, err := tx.ExecContext(ctx, "INSERT INTO placeVotes (id, user_id, place_id, rank) VALUES (?, ?, ?, ?)", uuid.New(), userID, placeVote.ID, placeVote.Rank); err != nil {
			tx.Rollback()
			return uuid.Nil, fmt.Errorf("faild insert timeVote: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, fmt.Errorf("faild commit transaction: %w", err)
	}

	return userID, nil
}

func (r *Repository) UpdateVote(ctx context.Context, param UpdateVoteReq) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("faild start transaction: %w", err)
	}

	if _, err := tx.ExecContext(ctx, "DELETE FROM timeVotes WHERE user_id = ?", uuid.New(), param.UserID); err != nil {
		tx.Rollback()
		return fmt.Errorf("faild delete timeVote: %w", err)
	}

	if _, err := tx.ExecContext(ctx, "DELETE FROM placeVotes WHERE user_id = ?", uuid.New(), param.UserID); err != nil {
		tx.Rollback()
		return fmt.Errorf("faild delete timeVote: %w", err)
	}

	for _, timeVote := range param.TimeVotes {
		if _, err := tx.ExecContext(ctx, "INSERT INTO timeVotes (id, user_id, time_id) VALUES (?, ?, ?)", uuid.New(), param.UserID, timeVote); err != nil {
			tx.Rollback()
			return fmt.Errorf("faild insert timeVote: %w", err)
		}
	}

	for _, placeVote := range param.PlaceVotes {
		if _, err := tx.ExecContext(ctx, "INSERT INTO placeVotes (id, user_id, place_id, rank) VALUES (?, ?, ?, ?)", uuid.New(), param.UserID, placeVote.ID, placeVote.Rank); err != nil {
			tx.Rollback()
			return fmt.Errorf("faild insert timeVote: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("faild commit transaction: %w", err)
	}

	return nil
}
