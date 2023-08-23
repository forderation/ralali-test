package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/forderation/ralali-test/internal/model"
	"github.com/forderation/ralali-test/internal/repository"
	"github.com/forderation/ralali-test/util"
)

type CakeUsecase struct {
	dbCakeRepository repository.CakeDBInterface
}

func NewCakeUsecase(dbCakeRepository repository.CakeDBInterface) CakeUsecaseInterface {
	return &CakeUsecase{
		dbCakeRepository: dbCakeRepository,
	}
}

func (uc *CakeUsecase) DeleteCake(ctx context.Context, id int) (*model.CakeDeleteResponse, *model.ErrorResponse) {
	_, errResponse := uc.GetDetailCake(ctx, id)
	if errResponse != nil {
		return nil, errResponse
	}
	err := uc.dbCakeRepository.SoftDeleteCake(ctx, id)
	if err != nil {
		return nil, &model.ErrorResponse{
			HttpStatusCode: http.StatusInternalServerError,
			Err:            errors.New("error delete cake data"),
			ErrData: model.ErrorDetailResponse{
				Detail: err.Error(),
			},
		}
	}
	return &model.CakeDeleteResponse{
		ID: id,
	}, nil
}

func (uc *CakeUsecase) UpdateCake(ctx context.Context, id int, payload model.CakePayloadQuery) (*model.CakeMutationResponse, *model.ErrorResponse) {
	_, errResponse := uc.GetDetailCake(ctx, id)
	if errResponse != nil {
		return nil, errResponse
	}
	err := uc.dbCakeRepository.UpdateCake(ctx, id, payload)
	if err != nil {
		return nil, &model.ErrorResponse{
			HttpStatusCode: http.StatusInternalServerError,
			Err:            errors.New("error update cake data"),
			ErrData: model.ErrorDetailResponse{
				Detail: err.Error(),
			},
		}
	}
	return &model.CakeMutationResponse{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
	}, nil
}

func (uc *CakeUsecase) CreateCake(ctx context.Context, payload model.CakePayloadQuery) (*model.CakeMutationResponse, *model.ErrorResponse) {
	err := uc.dbCakeRepository.InsertCake(ctx, payload)
	if err != nil {
		return nil, &model.ErrorResponse{
			HttpStatusCode: http.StatusInternalServerError,
			Err:            errors.New("error add cake data"),
			ErrData: model.ErrorDetailResponse{
				Detail: err.Error(),
			},
		}
	}
	return &model.CakeMutationResponse{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
	}, nil
}

func (uc *CakeUsecase) GetDetailCake(ctx context.Context, id int) (*model.CakeResponse, *model.ErrorResponse) {
	cake, err := uc.dbCakeRepository.GetCake(ctx, id)
	if err != nil {
		return nil, &model.ErrorResponse{
			HttpStatusCode: http.StatusInternalServerError,
			Err:            errors.New("error get cake data"),
			ErrData: model.ErrorDetailResponse{
				Detail: err.Error(),
			},
		}
	}
	if cake == nil {
		return nil, &model.ErrorResponse{
			HttpStatusCode: http.StatusNotFound,
			Err:            fmt.Errorf("cake data with id %d not found", id),
		}
	}
	response := mapCakeDataResponse(*cake)
	return &response, nil
}

func (uc *CakeUsecase) GetCakes(ctx context.Context, param model.GetCakesUsecaseParam) (*model.GetCakesResponse, *model.ErrorResponse) {
	offset := int(0)
	if param.Page > 0 {
		offset = (param.Page - 1) * param.PageSize
	}
	limit := param.PageSize
	var totalData int64
	var errTotal, errCakes error
	var cakes []model.Cake
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		totalData, errTotal = uc.dbCakeRepository.CountCakes(ctx)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		cakes, errCakes = uc.dbCakeRepository.GetCakes(ctx, model.GetCakesQuery{
			Limit:  limit,
			Offset: offset,
		})
		wg.Done()
	}()
	wg.Wait()
	if errCakes != nil {
		return nil, &model.ErrorResponse{
			Err:            errors.New("error on get data cakes"),
			HttpStatusCode: http.StatusInternalServerError,
			ErrData: model.ErrorDetailResponse{
				Detail: errCakes.Error(),
			},
		}
	}
	if errTotal != nil {
		return nil, &model.ErrorResponse{
			Err:            errors.New("error on count total cakes"),
			HttpStatusCode: http.StatusInternalServerError,
			ErrData: model.ErrorDetailResponse{
				Detail: errTotal.Error(),
			},
		}
	}
	response := model.GetCakesResponse{
		Meta: model.MetaPagination{
			PageCount: util.GetPageCount(param.PageSize, int(totalData)),
			TotalData: totalData,
		},
		Data: make([]model.CakeResponse, 0),
	}
	for _, v := range cakes {
		data := mapCakeDataResponse(v)
		response.Data = append(response.Data, data)
	}
	return &response, nil
}

func mapCakeDataResponse(cake model.Cake) model.CakeResponse {
	return model.CakeResponse{
		ID:          cake.ID,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
		CreatedAt:   cake.CreatedAt.Local().Format(time.DateTime),
		UpdatedAt:   cake.UpdatedAt.Local().Format(time.DateTime),
	}
}
