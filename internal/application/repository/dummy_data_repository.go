package repository

import (
	"context"
	"dummies-backend/internal/domain/model"
)

type DummyDataRepository interface {
	FindByProjectID(ctx context.Context, projectID int) ([]*model.DummyData, error)
	FindByID(ctx context.Context, id string) (*model.DummyData, error)
	Create(ctx context.Context, data *model.DummyData) (*model.DummyData, error)
	Update(ctx context.Context, data *model.DummyData) error
	SoftDelete(ctx context.Context, id string) error
}
