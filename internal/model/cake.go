package model

import (
	"time"
)

// Cake: represent model of cakes table
type Cake struct {
	ID          int
	Title       string
	Description *string
	Rating      float32
	Image       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
