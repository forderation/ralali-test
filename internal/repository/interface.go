package repository

import (
	"context"

	"github.com/forderation/ralali-test/internal/model"
)

//go:generate moq -out mock_interface.go . CakeDBInterface
type CakeDBInterface interface {
	GetCakes(ctx context.Context, param model.GetCakesQuery) ([]model.Cake, error)
	CountCakes(ctx context.Context) (int64, error)
	GetCake(ctx context.Context, id int) (*model.Cake, error)
	InsertCake(ctx context.Context, param model.CakePayloadQuery) error
	UpdateCake(ctx context.Context, id int, param model.CakePayloadQuery) error
	SoftDeleteCake(ctx context.Context, id int) error
}
