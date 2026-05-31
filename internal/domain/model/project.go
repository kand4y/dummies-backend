package model

import "time"

type Project struct {
	ID          int
	UserID      string
	Status      Status
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
