package repository

import (
	"context"
	"fmt"
	// "time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type (
	// timeVote table
	TimeVote struct {
		ID     uuid.UUID `db:"id"`
		UserID uuid.UUID `db:"user_id"`
		TimeID uuid.UUID `db:"time_id"`
	}

	ReturnTimeVote struct {
		ID        uuid.UUID `db:"id"`
		UserID    uuid.UUID `db:"user_id"`
		StartTime string `db:"start_time"`
		EndTime   string `db:"end_time"`
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

func (r *Repository) GetTimeVotesByUserID(ctx context.Context, userID uuid.UUID) ([]*ReturnTimeVote, error) {
	timeVotes := []*ReturnTimeVote{}
	if err := r.db.SelectContext(ctx, &timeVotes, "SELECT timeVotes.id, timeVotes.user_id, timeOptions.start_time, timeOptions.end_time FROM timeVotes JOIN timeOptions ON timeVotes.time_id = timeOptions.id WHERE timeVotes.user_id = ?", userID); err != nil {
		return nil, fmt.Errorf("select timeVotes by userID: %w", err)
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

func (r *Repository) CreateTimeVotes(ctx context.Context, params []CreateTimeVoteParams) (error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("faild start transaction: %w", err)
	}
	timeVoteIDs := lo.Times(len(params), func(_ int) uuid.UUID {
		return uuid.New()
	})
	for i, param := range params {
		if _, err := tx.ExecContext(ctx, "INSERT INTO timeVotes (id, user_id, time_id) VALUES (?, ?, ?)", timeVoteIDs[i], param.UserID, param.TimeID); err != nil {
			tx.Rollback()
			return fmt.Errorf("faild insert timeVote: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("faild commit transaction: %w", err)	
	}

	return nil
}

func (r *Repository) GetTimeVote(ctx context.Context, timeVoteID uuid.UUID) (*TimeVote, error) {
	timeVote := &TimeVote{}
	if err := r.db.GetContext(ctx, timeVote, "SELECT * FROM timeVotes WHERE id = ?", timeVoteID); err != nil {
		return nil, fmt.Errorf("select timeVote: %w", err)
	}
	return timeVote, nil
}
