package repository

import (
	"context"

	"github.com/forderation/ralali-test/internal/model"
)

//go:generate moq -out mock_interface.go . CakeDBInterface
type CakeDBInterface interface {
	// GetCakes: get all cake record with not soft delete, parameter with pagination
	GetCakes(ctx context.Context, param model.GetCakesQuery) ([]model.Cake, error)
	// CountCakes: get count all cake record with not soft delete, will return 0 and error exist if query error
	CountCakes(ctx context.Context) (int64, error)
	// GetCake: get single cake record, required id record, will return nil if record not found at *model.Cake
	GetCake(ctx context.Context, id int) (*model.Cake, error)
	// InsertCake: insert new cake record, required parameter refer to model.CakePayloadQuery
	InsertCake(ctx context.Context, param model.CakePayloadQuery) error
	// UpdateCake: update cake record data, required id record and parameter refer to model.CakePayloadQuery
	UpdateCake(ctx context.Context, id int, param model.CakePayloadQuery) error
	// SoftDeleteCake: updating cake record data with filled deleted_at, required id record
	SoftDeleteCake(ctx context.Context, id int) error
}
