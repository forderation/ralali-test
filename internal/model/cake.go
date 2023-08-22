package model

import (
	"database/sql"
	"time"
)

// Cake: represent model of cakes table
type Cake struct {
	ID          int
	Title       string
	Description sql.NullString
	Rating      float64
	Image       sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
