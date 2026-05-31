package repository

import (
	"context"
	"dummies-backend/internal/domain/model"
)

type ProjectRepository interface {
	FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Project, int, error)
	FindByID(ctx context.Context, id int) (*model.Project, error)
	Create(ctx context.Context, project *model.Project) (*model.Project, error)
	Update(ctx context.Context, project *model.Project) error
	SoftDelete(ctx context.Context, id int) error
}
