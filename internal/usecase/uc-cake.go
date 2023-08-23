package usecase

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/forderation/ralali-test/internal/model"
	"github.com/forderation/ralali-test/internal/repository"
	"github.com/forderation/ralali-test/util"
)

type CakeUsecase struct {
	dbCakeRepository *repository.CakeDBRepository
}

func NewCakeUsecase(dbCakeRepository *repository.CakeDBRepository) *CakeUsecase {
	return &CakeUsecase{
		dbCakeRepository: dbCakeRepository,
	}
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
			Err:            errors.New("cake data not found"),
		}
	}
	response := uc.MapCakeDataResponse(*cake)
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
		data := uc.MapCakeDataResponse(v)
		response.Data = append(response.Data, data)
	}
	return &response, nil
}

func (uc *CakeUsecase) MapCakeDataResponse(cake model.Cake) model.CakeResponse {
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
