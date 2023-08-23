package usecase

import (
	"context"
	"errors"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/forderation/ralali-test/internal/model"
	"github.com/forderation/ralali-test/internal/repository"
	"gopkg.in/guregu/null.v4"
)

func TestNewCakeUsecase(t *testing.T) {
	type args struct {
		dbCakeRepository repository.CakeDBInterface
	}
	mockCakeRepo := &repository.CakeDBInterfaceMock{}
	tests := []struct {
		name string
		args args
		want CakeUsecaseInterface
	}{
		{
			name: "basic test",
			args: args{
				dbCakeRepository: mockCakeRepo,
			},
			want: &CakeUsecase{
				dbCakeRepository: mockCakeRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCakeUsecase(tt.args.dbCakeRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCakeUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeUsecase_DeleteCake(t *testing.T) {
	type fields struct {
		dbCakeRepository repository.CakeDBInterface
	}
	type args struct {
		ctx context.Context
		id  int
	}
	mockCakeRepo := &repository.CakeDBInterfaceMock{}
	mockCakeRepo.GetCakeFunc = func(ctx context.Context, id int) (*model.Cake, error) {
		return &model.Cake{}, nil
	}
	mockCakeRepo.SoftDeleteCakeFunc = func(ctx context.Context, id int) error {
		if id == 1 {
			return nil
		}
		return errors.New("error mock")
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CakeDeleteResponse
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			want: &model.CakeDeleteResponse{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name: "error test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				id:  2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CakeUsecase{
				dbCakeRepository: tt.fields.dbCakeRepository,
			}
			got, err := uc.DeleteCake(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeUsecase.DeleteCake() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeUsecase.DeleteCake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeUsecase_UpdateCake(t *testing.T) {
	type fields struct {
		dbCakeRepository repository.CakeDBInterface
	}
	type args struct {
		ctx     context.Context
		id      int
		payload model.CakePayloadQuery
	}
	mockCakeRepo := &repository.CakeDBInterfaceMock{}
	mockCakeRepo.GetCakeFunc = func(ctx context.Context, id int) (*model.Cake, error) {
		return &model.Cake{}, nil
	}
	mockCakeRepo.UpdateCakeFunc = func(ctx context.Context, id int, param model.CakePayloadQuery) error {
		if id == 1 {
			return nil
		}
		return errors.New("error mock")
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CakeMutationResponse
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				id:  1,
				payload: model.CakePayloadQuery{
					Title:       "title",
					Description: null.StringFrom("description").Ptr(),
					Rating:      float32(4.32),
					Image:       null.StringFrom("image").Ptr(),
				},
			},
			want: &model.CakeMutationResponse{
				Title:       "title",
				Description: null.StringFrom("description").Ptr(),
				Rating:      float32(4.32),
				Image:       null.StringFrom("image").Ptr(),
			},
		},
		{
			name: "error test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx:     context.TODO(),
				id:      2,
				payload: model.CakePayloadQuery{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CakeUsecase{
				dbCakeRepository: tt.fields.dbCakeRepository,
			}
			got, err := uc.UpdateCake(tt.args.ctx, tt.args.id, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeUsecase.UpdateCake() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeUsecase.UpdateCake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeUsecase_CreateCake(t *testing.T) {
	type fields struct {
		dbCakeRepository repository.CakeDBInterface
	}
	type args struct {
		ctx     context.Context
		payload model.CakePayloadQuery
	}
	mockCakeRepo := &repository.CakeDBInterfaceMock{}
	mockCakeRepo.InsertCakeFunc = func(ctx context.Context, param model.CakePayloadQuery) error {
		if param.Title == "title" {
			return nil
		}
		return errors.New("error mock")
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CakeMutationResponse
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				payload: model.CakePayloadQuery{
					Title:       "title",
					Description: null.StringFrom("description").Ptr(),
					Rating:      float32(4.32),
					Image:       null.StringFrom("image").Ptr(),
				},
			},
			want: &model.CakeMutationResponse{
				Title:       "title",
				Description: null.StringFrom("description").Ptr(),
				Rating:      float32(4.32),
				Image:       null.StringFrom("image").Ptr(),
			},
		},
		{
			name: "error test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx:     context.TODO(),
				payload: model.CakePayloadQuery{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CakeUsecase{
				dbCakeRepository: tt.fields.dbCakeRepository,
			}
			got, err := uc.CreateCake(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeUsecase.CreateCake() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeUsecase.CreateCake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeUsecase_GetDetailCake(t *testing.T) {
	type fields struct {
		dbCakeRepository repository.CakeDBInterface
	}
	type args struct {
		ctx context.Context
		id  int
	}
	timeMock, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	if err != nil {
		log.Fatalf("an error '%s' was not expected", err)
	}
	timeMock = timeMock.UTC()
	mockCakeRepo := &repository.CakeDBInterfaceMock{}
	mockCakeRepo.GetCakeFunc = func(ctx context.Context, id int) (*model.Cake, error) {
		if id == 1 {
			return &model.Cake{
				ID:          1,
				Title:       "title",
				Description: null.StringFrom("description").Ptr(),
				Rating:      float32(4.32),
				Image:       null.StringFrom("image").Ptr(),
				CreatedAt:   timeMock,
				UpdatedAt:   timeMock,
			}, nil
		}
		return nil, nil
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CakeResponse
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			want: &model.CakeResponse{
				ID:          1,
				Title:       "title",
				Description: null.StringFrom("description").Ptr(),
				Rating:      float32(4.32),
				Image:       null.StringFrom("image").Ptr(),
				CreatedAt:   "2006-01-02 22:04:05",
				UpdatedAt:   "2006-01-02 22:04:05",
			},
		},
		{
			name: "error test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				id:  2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CakeUsecase{
				dbCakeRepository: tt.fields.dbCakeRepository,
			}
			got, err := uc.GetDetailCake(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeUsecase.GetDetailCake() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeUsecase.GetDetailCake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeUsecase_GetCakes(t *testing.T) {
	type fields struct {
		dbCakeRepository repository.CakeDBInterface
	}
	type args struct {
		ctx   context.Context
		param model.GetCakesUsecaseParam
	}
	timeMock, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	if err != nil {
		log.Fatalf("an error '%s' was not expected", err)
	}
	timeMock = timeMock.UTC()
	mockCakeRepo := &repository.CakeDBInterfaceMock{}
	mockCakeRepo.GetCakesFunc = func(ctx context.Context, param model.GetCakesQuery) ([]model.Cake, error) {
		if param.Limit == 1 {
			return []model.Cake{
				{
					ID:          1,
					Title:       "title",
					Description: null.StringFrom("description").Ptr(),
					Rating:      float32(4.32),
					Image:       null.StringFrom("image").Ptr(),
					CreatedAt:   timeMock,
					UpdatedAt:   timeMock,
				},
			}, nil
		}
		return nil, errors.New("error mock")
	}
	mockCakeRepo.CountCakesFunc = func(ctx context.Context) (int64, error) {
		return 1, nil
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.GetCakesResponse
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				param: model.GetCakesUsecaseParam{
					Page:     1,
					PageSize: 1,
				},
			},
			want: &model.GetCakesResponse{
				Meta: model.MetaPagination{
					TotalData: 1,
					PageCount: 1,
				},
				Data: []model.CakeResponse{
					{
						ID:          1,
						Title:       "title",
						Description: null.StringFrom("description").Ptr(),
						Rating:      float32(4.32),
						Image:       null.StringFrom("image").Ptr(),
						CreatedAt:   "2006-01-02 22:04:05",
						UpdatedAt:   "2006-01-02 22:04:05",
					},
				},
			},
		},
		{
			name: "error test",
			fields: fields{
				dbCakeRepository: mockCakeRepo,
			},
			args: args{
				ctx: context.TODO(),
				param: model.GetCakesUsecaseParam{
					Page:     1,
					PageSize: 2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CakeUsecase{
				dbCakeRepository: tt.fields.dbCakeRepository,
			}
			got, err := uc.GetCakes(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeUsecase.GetCakes() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeUsecase.GetCakes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapCakeDataResponse(t *testing.T) {
	type args struct {
		cake model.Cake
	}
	timeMock, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	if err != nil {
		log.Fatalf("an error '%s' was not expected", err)
	}
	timeMock = timeMock.UTC()
	tests := []struct {
		name string
		args args
		want model.CakeResponse
	}{
		{
			name: "basic test",
			args: args{
				cake: model.Cake{
					ID:          1,
					Title:       "title",
					Description: null.StringFrom("description").Ptr(),
					Rating:      float32(4.32),
					Image:       null.StringFrom("image").Ptr(),
					CreatedAt:   timeMock,
					UpdatedAt:   timeMock,
				},
			},
			want: model.CakeResponse{
				ID:          1,
				Title:       "title",
				Description: null.StringFrom("description").Ptr(),
				Rating:      float32(4.32),
				Image:       null.StringFrom("image").Ptr(),
				CreatedAt:   "2006-01-02 22:04:05",
				UpdatedAt:   "2006-01-02 22:04:05",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapCakeDataResponse(tt.args.cake); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapCakeDataResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
