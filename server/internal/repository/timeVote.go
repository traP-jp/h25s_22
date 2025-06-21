package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// timeVote table
	TimeVote struct {
		ID     uuid.UUID `db:"id"`
		UserID uuid.UUID `db:"user_id"`
		TimeID uuid.UUID `db:"time_id"`
	}

	CreateTimeVoteParams struct {
		UserID uuid.UUID `db:"user_id"`
		TimeID uuid.UUID `db:"time_id"`
	}
)

func (r *Repository) GetTimeVotes(ctx context.Context) ([]*TimeVote, error) {
	timeVotes := []*TimeVote{}
	if err := r.db.SelectContext(ctx, &timeVotes, "SELECT * FROM timeVotes"); err != nil {
		return nil, fmt.Errorf("select timeVotes: %w", err)
	}

	return timeVotes, nil
}

func (r *Repository) CreateTimeVote(ctx context.Context, params CreateTimeVoteParams) (uuid.UUID, error) {
	timeVoteID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO timeVotes (id, user_id, time_id) VALUES (?, ?, ?)", timeVoteID, params.UserID, params.TimeID); err != nil {
		return uuid.Nil, fmt.Errorf("insert timeVote: %w", err)
	}

	return timeVoteID, nil
}

func (r *Repository) GetTimeVote(ctx context.Context, timeVoteID uuid.UUID) (*TimeVote, error) {
	timeVote := &TimeVote{}
	if err := r.db.GetContext(ctx, timeVote, "SELECT * FROM timeVotes WHERE id = ?", timeVoteID); err != nil {
		return nil, fmt.Errorf("select timeVote: %w", err)
	}
	return timeVote, nil
}
