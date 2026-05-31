package model

import "time"

type User struct {
	ID         string
	Status     Status
	UserHandle string
	UserName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
