package dto

import (
	"dummies-backend/internal/domain/model"
	"encoding/json"
	"time"
)

type UserResponse struct {
	ID         string `json:"id"`
	UserHandle string `json:"user_handle"`
	UserName   string `json:"user_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func ToUserResponse(u *model.User) *UserResponse {
	return &UserResponse{
		ID:         u.ID,
		UserHandle: u.UserHandle,
		UserName:   u.UserName,
		CreatedAt:  u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  u.UpdatedAt.Format(time.RFC3339),
	}
}

type ProjectResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ToProjectResponse(p *model.Project) *ProjectResponse {
	return &ProjectResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
	}
}

type ProjectListResponse struct {
	Projects   []*ProjectResponse `json:"projects"`
	TotalCount int                `json:"total_count"`
	Page       int                `json:"page"`
	PerPage    int                `json:"per_page"`
}

type DummyDataResponse struct {
	ID             string   `json:"id"`
	ProjectID      int      `json:"project_id"`
	TableName      string   `json:"table_name"`
	ColumnName     []string `json:"column_name"`
	ColumnType     []string `json:"column_type"`
	ColumnValidate []string `json:"column_validate"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

func ToDummyDataResponse(d *model.DummyData) *DummyDataResponse {
	var colNames, colTypes, colValidates []string
	_ = json.Unmarshal([]byte(d.ColumnName), &colNames)
	_ = json.Unmarshal([]byte(d.ColumnType), &colTypes)
	if d.ColumnValidate != nil {
		_ = json.Unmarshal([]byte(*d.ColumnValidate), &colValidates)
	}
	if colValidates == nil {
		colValidates = []string{}
	}

	return &DummyDataResponse{
		ID:             d.ID,
		ProjectID:      d.ProjectID,
		TableName:      d.TableName,
		ColumnName:     colNames,
		ColumnType:     colTypes,
		ColumnValidate: colValidates,
		CreatedAt:      d.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      d.UpdatedAt.Format(time.RFC3339),
	}
}
