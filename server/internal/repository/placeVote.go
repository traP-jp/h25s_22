package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
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

func (r *Repository) GetPlaceVotes(ctx context.Context) ([]*placeVote, error) {
	placeVotes := []*placeVote{}
	if err := r.db.SelectContext(ctx, &placeVotes, "SELECT * FROM placeVotes"); err != nil {
		return nil, fmt.Errorf("select placeVotes: %w", err)
	}

	return placeVotes, nil
}

func (r *Repository) GetPlaceVotesByUserID(ctx context.Context, userID uuid.UUID) ([]*placeVote, error) {
	placeVotes := []*placeVote{}
	if err := r.db.SelectContext(ctx, &placeVotes, "SELECT * FROM placeVotes WHERE user_id = ?", userID); err != nil {
		return nil, fmt.Errorf("select placeVotes by userID: %w", err)
	}
	return placeVotes, nil
}

func (r *Repository) CreatePlaceVote(ctx context.Context, params CreatePlaceVoteParams) (uuid.UUID, error) {
	voteID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO placeVotes (id, user_id, place_id, rank) VALUES (?, ?, ?, ?)", voteID, params.UserID, params.PlaceID, params.Rank); err != nil {
		return uuid.Nil, fmt.Errorf("insert vote: %w", err)
	}

	return voteID, nil
}

func (r *Repository) CreatePlaceVotes(ctx context.Context, params []CreatePlaceVoteParams) ([]uuid.UUID, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("faild start transaction: %w", err)
	}
	placeVoteIDs := lo.Times(len(params), func(_ int) uuid.UUID {
		return uuid.New()
	})
	for i, param := range params {
		if _, err := tx.ExecContext(ctx, "INSERT INTO placeVotes (id, user_id, place_id, rank) VALUES (?, ?, ?, ?)", placeVoteIDs[i], param.UserID, param.PlaceID, param.Rank); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("faild insert timeVote: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("faild commit transaction: %w", err)
	}

	return placeVoteIDs, nil
}

func (r *Repository) GetPlaceVote(ctx context.Context, voteID uuid.UUID) (*placeVote, error) {
	vote := &placeVote{}
	if err := r.db.GetContext(ctx, vote, "SELECT * FROM placeVotes WHERE id = ?", voteID); err != nil {
		return nil, fmt.Errorf("select vote: %w", err)
	}

	return vote, nil
}
