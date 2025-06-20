package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	// timeOption table
	timeOption struct {
		ID        uuid.UUID `db:"id"`
		StartTime time.Time `db:"start_time"`
		EndTime   time.Time `db:"end_time"`
	}

	CreateTimeOptionParams struct {
		StartTime time.Time `db:"start_time"`
		EndTime   time.Time `db:"end_time"`
	}
)

func (r *Repository) GetTimeOptions(ctx context.Context) ([]*timeOption, error) {
	timeOptions := []*timeOption{}
	if err := r.db.SelectContext(ctx, &timeOptions, "SELECT * FROM timeOptions"); err != nil {
		return nil, fmt.Errorf("select timeOptions: %w", err)
	}

	return timeOptions, nil
}

func (r *Repository) CreateTimeOption(ctx context.Context, params CreateTimeOptionParams) (uuid.UUID, error) {
	timeOptionID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO timeOptions (id, start_time, end_time) VALUES (?, ?, ?)", timeOptionID, params.StartTime, params.EndTime); err != nil {
		return uuid.Nil, fmt.Errorf("insert timeOption: %w", err)
	}

	return timeOptionID, nil
}

func (r *Repository) GetTimeOption(ctx context.Context, timeOptionID uuid.UUID) (*timeOption, error) {
	timeOption := &timeOption{}
	if err := r.db.GetContext(ctx, timeOption, "SELECT * FROM timeOptions WHERE id = ?", timeOptionID); err != nil {
		return nil, fmt.Errorf("select timeOption: %w", err)
	}
	return timeOption, nil
}
