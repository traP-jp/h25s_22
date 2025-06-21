package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// placeVotes table
	placeVote struct {
		ID      uuid.UUID `db:"id"`
		UserID  uuid.UUID `db:"user_id"`
		PlaceID uuid.UUID `db:"place_id"`
		Rank    int       `db:"rank"`
	}

	CreatePlaceVoteParams struct {
		UserID  uuid.UUID `db:"user_id"`
		PlaceID uuid.UUID `db:"place_id"`
		Rank    int       `db:"rank"`
	}
)

func (r *Repository) GetVotes(ctx context.Context) ([]*placeVote, error) {
	placeVotes := []*placeVote{}
	if err := r.db.SelectContext(ctx, &placeVotes, "SELECT * FROM placeVote"); err != nil {
		return nil, fmt.Errorf("select placeVote: %w", err)
	}

	return placeVotes, nil
}

func (r *Repository) CreateVote(ctx context.Context, params CreatePlaceVoteParams) (uuid.UUID, error) {
	voteID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO placeVote (id, user_id, place_id, rank) VALUES (?, ?, ?, ?)", voteID, params.UserID, params.PlaceID, params.Rank); err != nil {
		return uuid.Nil, fmt.Errorf("insert vote: %w", err)
	}

	return voteID, nil
}

func (r *Repository) GetVote(ctx context.Context, voteID uuid.UUID) (*placeVote, error) {
	vote := &placeVote{}
	if err := r.db.GetContext(ctx, vote, "SELECT * FROM placeVote WHERE id = ?", voteID); err != nil {
		return nil, fmt.Errorf("select vote: %w", err)
	}

	return vote, nil
}
