package model

import "time"

type DummyData struct {
	ID             string
	ProjectID      int
	Status         Status
	TableName      string
	ColumnName     string
	ColumnType     string
	ColumnValidate *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
