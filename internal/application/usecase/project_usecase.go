package usecase

import (
	"context"
	"dummies-backend/internal/application/repository"
	"dummies-backend/internal/domain/model"
	"errors"
)

const ProjectsPerPage = 12

type ProjectUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewProjectUseCase(projectRepo repository.ProjectRepository) *ProjectUseCase {
	return &ProjectUseCase{projectRepo: projectRepo}
}

func (uc *ProjectUseCase) ListByUser(ctx context.Context, userID string, page int) ([]*model.Project, int, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * ProjectsPerPage
	return uc.projectRepo.FindByUserID(ctx, userID, ProjectsPerPage, offset)
}

func (uc *ProjectUseCase) GetByID(ctx context.Context, projectID int, userID string) (*model.Project, error) {
	project, err := uc.projectRepo.FindByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New("project not found")
	}
	if project.UserID != userID {
		return nil, errors.New("forbidden")
	}
	return project, nil
}

func (uc *ProjectUseCase) Create(ctx context.Context, userID, name string, description *string) (*model.Project, error) {
	project := &model.Project{
		UserID:      userID,
		Status:      model.StatusActive,
		Name:        name,
		Description: description,
	}
	return uc.projectRepo.Create(ctx, project)
}

func (uc *ProjectUseCase) Update(ctx context.Context, projectID int, userID, name string, description *string) error {
	project, err := uc.projectRepo.FindByID(ctx, projectID)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project not found")
	}
	if project.UserID != userID {
		return errors.New("forbidden")
	}
	project.Name = name
	project.Description = description
	return uc.projectRepo.Update(ctx, project)
}

func (uc *ProjectUseCase) Delete(ctx context.Context, projectID int, userID string) error {
	project, err := uc.projectRepo.FindByID(ctx, projectID)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project not found")
	}
	if project.UserID != userID {
		return errors.New("forbidden")
	}
	return uc.projectRepo.SoftDelete(ctx, projectID)
}
