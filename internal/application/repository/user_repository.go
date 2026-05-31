package repository

import (
	"context"
	"dummies-backend/internal/domain/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	SoftDelete(ctx context.Context, id string) error
}
