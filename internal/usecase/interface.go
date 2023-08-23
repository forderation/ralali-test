package usecase

import (
	"context"

	"github.com/forderation/ralali-test/internal/model"
)

//go:generate moq -out mock_interface.go . CakeUsecaseInterface
type CakeUsecaseInterface interface {
	DeleteCake(ctx context.Context, id int) (*model.CakeDeleteResponse, *model.ErrorResponse)
	UpdateCake(ctx context.Context, id int, payload model.CakePayloadQuery) (*model.CakeMutationResponse, *model.ErrorResponse)
	CreateCake(ctx context.Context, payload model.CakePayloadQuery) (*model.CakeMutationResponse, *model.ErrorResponse)
	GetDetailCake(ctx context.Context, id int) (*model.CakeResponse, *model.ErrorResponse)
	GetCakes(ctx context.Context, param model.GetCakesUsecaseParam) (*model.GetCakesResponse, *model.ErrorResponse)
}
