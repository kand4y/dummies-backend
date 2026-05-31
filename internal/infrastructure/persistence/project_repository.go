package persistence

import (
	"context"
	"dummies-backend/internal/domain/model"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewProjectRepository(pool *pgxpool.Pool) *ProjectRepositoryImpl {
	return &ProjectRepositoryImpl{pool: pool}
}

func (r *ProjectRepositoryImpl) FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Project, int, error) {
	var totalCount int
	err := r.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM public.projects WHERE user_id = $1 AND status = 'active'`, userID).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, status, name, description, created_at, updated_at
		 FROM public.projects WHERE user_id = $1 AND status = 'active'
		 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		var p model.Project
		if err := rows.Scan(&p.ID, &p.UserID, &p.Status, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, 0, err
		}
		projects = append(projects, &p)
	}
	return projects, totalCount, rows.Err()
}

func (r *ProjectRepositoryImpl) FindByID(ctx context.Context, id int) (*model.Project, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, user_id, status, name, description, created_at, updated_at
		 FROM public.projects WHERE id = $1 AND status = 'active'`, id)

	var p model.Project
	err := row.Scan(&p.ID, &p.UserID, &p.Status, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProjectRepositoryImpl) Create(ctx context.Context, project *model.Project) (*model.Project, error) {
	row := r.pool.QueryRow(ctx,
		`INSERT INTO public.projects (user_id, status, name, description)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, user_id, status, name, description, created_at, updated_at`,
		project.UserID, project.Status, project.Name, project.Description)

	var p model.Project
	err := row.Scan(&p.ID, &p.UserID, &p.Status, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProjectRepositoryImpl) Update(ctx context.Context, project *model.Project) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE public.projects SET name = $1, description = $2 WHERE id = $3 AND status = 'active'`,
		project.Name, project.Description, project.ID)
	return err
}

func (r *ProjectRepositoryImpl) SoftDelete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE public.projects SET status = 'deleted' WHERE id = $1`, id)
	return err
}
