package usecase

import (
	"context"
	"dummies-backend/internal/application/repository"
	"dummies-backend/internal/domain/model"
	"errors"
)

type DummyDataUseCase struct {
	dummyDataRepo repository.DummyDataRepository
	projectRepo   repository.ProjectRepository
}

func NewDummyDataUseCase(dummyDataRepo repository.DummyDataRepository, projectRepo repository.ProjectRepository) *DummyDataUseCase {
	return &DummyDataUseCase{
		dummyDataRepo: dummyDataRepo,
		projectRepo:   projectRepo,
	}
}

func (uc *DummyDataUseCase) verifyProjectOwnership(ctx context.Context, projectID int, userID string) error {
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
	return nil
}

func (uc *DummyDataUseCase) ListByProject(ctx context.Context, projectID int, userID string) ([]*model.DummyData, error) {
	if err := uc.verifyProjectOwnership(ctx, projectID, userID); err != nil {
		return nil, err
	}
	return uc.dummyDataRepo.FindByProjectID(ctx, projectID)
}

func (uc *DummyDataUseCase) GetByID(ctx context.Context, id string, userID string) (*model.DummyData, error) {
	data, err := uc.dummyDataRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("dummy data not found")
	}
	if err := uc.verifyProjectOwnership(ctx, data.ProjectID, userID); err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *DummyDataUseCase) Create(ctx context.Context, projectID int, userID string, tableName, columnName, columnType string, columnValidate *string) (*model.DummyData, error) {
	if err := uc.verifyProjectOwnership(ctx, projectID, userID); err != nil {
		return nil, err
	}
	data := &model.DummyData{
		ProjectID:      projectID,
		Status:         model.StatusActive,
		TableName:      tableName,
		ColumnName:     columnName,
		ColumnType:     columnType,
		ColumnValidate: columnValidate,
	}
	return uc.dummyDataRepo.Create(ctx, data)
}

func (uc *DummyDataUseCase) Update(ctx context.Context, id string, userID string, tableName, columnName, columnType string, columnValidate *string) error {
	data, err := uc.dummyDataRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("dummy data not found")
	}
	if err := uc.verifyProjectOwnership(ctx, data.ProjectID, userID); err != nil {
		return err
	}
	data.TableName = tableName
	data.ColumnName = columnName
	data.ColumnType = columnType
	data.ColumnValidate = columnValidate
	return uc.dummyDataRepo.Update(ctx, data)
}

func (uc *DummyDataUseCase) Delete(ctx context.Context, id string, userID string) error {
	data, err := uc.dummyDataRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("dummy data not found")
	}
	if err := uc.verifyProjectOwnership(ctx, data.ProjectID, userID); err != nil {
		return err
	}
	return uc.dummyDataRepo.SoftDelete(ctx, id)
}
