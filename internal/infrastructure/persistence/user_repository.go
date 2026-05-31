package persistence

import (
	"context"
	"dummies-backend/internal/domain/model"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{pool: pool}
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (*model.User, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, status, user_handle, user_name, created_at, updated_at
		 FROM public.users WHERE id = $1 AND status = 'active'`, id)

	var u model.User
	err := row.Scan(&u.ID, &u.Status, &u.UserHandle, &u.UserName, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE public.users SET user_handle = $1, user_name = $2 WHERE id = $3 AND status = 'active'`,
		user.UserHandle, user.UserName, user.ID)
	return err
}

func (r *UserRepositoryImpl) SoftDelete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE public.users SET status = 'deleted' WHERE id = $1`, id)
	return err
}
