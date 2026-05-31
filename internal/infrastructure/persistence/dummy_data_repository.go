package persistence

import (
	"context"
	"dummies-backend/internal/domain/model"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DummyDataRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewDummyDataRepository(pool *pgxpool.Pool) *DummyDataRepositoryImpl {
	return &DummyDataRepositoryImpl{pool: pool}
}

func (r *DummyDataRepositoryImpl) FindByProjectID(ctx context.Context, projectID int) ([]*model.DummyData, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, project_id, status, table_name, column_name, column_type, column_validate, created_at, updated_at
		 FROM public.dummy_data WHERE project_id = $1 AND status = 'active'
		 ORDER BY created_at DESC`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*model.DummyData
	for rows.Next() {
		var d model.DummyData
		if err := rows.Scan(&d.ID, &d.ProjectID, &d.Status, &d.TableName, &d.ColumnName, &d.ColumnType, &d.ColumnValidate, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &d)
	}
	return list, rows.Err()
}

func (r *DummyDataRepositoryImpl) FindByID(ctx context.Context, id string) (*model.DummyData, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, project_id, status, table_name, column_name, column_type, column_validate, created_at, updated_at
		 FROM public.dummy_data WHERE id = $1 AND status = 'active'`, id)

	var d model.DummyData
	err := row.Scan(&d.ID, &d.ProjectID, &d.Status, &d.TableName, &d.ColumnName, &d.ColumnType, &d.ColumnValidate, &d.CreatedAt, &d.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DummyDataRepositoryImpl) Create(ctx context.Context, data *model.DummyData) (*model.DummyData, error) {
	row := r.pool.QueryRow(ctx,
		`INSERT INTO public.dummy_data (project_id, status, table_name, column_name, column_type, column_validate)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, project_id, status, table_name, column_name, column_type, column_validate, created_at, updated_at`,
		data.ProjectID, data.Status, data.TableName, data.ColumnName, data.ColumnType, data.ColumnValidate)

	var d model.DummyData
	err := row.Scan(&d.ID, &d.ProjectID, &d.Status, &d.TableName, &d.ColumnName, &d.ColumnType, &d.ColumnValidate, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DummyDataRepositoryImpl) Update(ctx context.Context, data *model.DummyData) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE public.dummy_data SET table_name = $1, column_name = $2, column_type = $3, column_validate = $4
		 WHERE id = $5 AND status = 'active'`,
		data.TableName, data.ColumnName, data.ColumnType, data.ColumnValidate, data.ID)
	return err
}

func (r *DummyDataRepositoryImpl) SoftDelete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE public.dummy_data SET status = 'deleted' WHERE id = $1`, id)
	return err
}
