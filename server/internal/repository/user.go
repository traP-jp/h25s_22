package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// users table
	User struct {
		ID     uuid.UUID `db:"id"`
		RoomID uuid.UUID `db:"room_id"`
		Name   string    `db:"name"`
	}

	CreateUserParams struct {
		RoomID uuid.UUID `db:"room_id"`
		Name   string    `db:"name"`
	}
)

func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {
	users := []*User{}
	if err := r.db.SelectContext(ctx, &users, "SELECT * FROM users"); err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	return users, nil
}

func (r *Repository) GetUsersByRoomID(ctx context.Context, roomID uuid.UUID) ([]*User, error) {
	users := []*User{}
	if err := r.db.SelectContext(ctx, &users, "SELECT * FROM users WHERE room_id = ?", roomID); err != nil {
		return nil, fmt.Errorf("select users by roomID: %w", err)
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, params CreateUserParams) (uuid.UUID, error) {
	userID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO users (id, room_id, name) VALUES (?, ?, ?)", userID, params.RoomID, params.Name); err != nil {
		return uuid.Nil, fmt.Errorf("insert user: %w", err)
	}

	return userID, nil
}

func (r *Repository) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	user := &User{}
	if err := r.db.GetContext(ctx, user, "SELECT * FROM users WHERE id = ?", userID); err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	return user, nil
}
