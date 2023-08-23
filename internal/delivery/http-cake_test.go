package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/forderation/ralali-test/internal/model"
	"github.com/forderation/ralali-test/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestNewCakeDelivery(t *testing.T) {
	type args struct {
		cakeUsecase usecase.CakeUsecaseInterface
	}
	tests := []struct {
		name string
		args args
		want *CakeDelivery
	}{
		{
			name: "basic test",
			args: args{
				cakeUsecase: &usecase.CakeUsecaseInterfaceMock{},
			},
			want: &CakeDelivery{
				cakeUsecase: &usecase.CakeUsecaseInterfaceMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCakeDelivery(tt.args.cakeUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCakeDelivery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeDelivery_GetCakes(t *testing.T) {
	type fields struct {
		cakeUsecase usecase.CakeUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	mockCakeUsecase := &usecase.CakeUsecaseInterfaceMock{}
	mockResponse := model.GetCakesResponse{
		Meta: model.MetaPagination{
			PageCount: 1,
			TotalData: 1,
		},
		Data: []model.CakeResponse{},
	}
	mockCakeUsecase.GetCakesFunc = func(ctx context.Context, param model.GetCakesUsecaseParam) (*model.GetCakesResponse, *model.ErrorResponse) {
		return &mockResponse, nil
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
	}
	queryRequest := url.Values{}
	queryRequest.Add("page", "1")
	queryRequest.Add("page_size", "1")
	ctx.Request.Form = queryRequest
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic test, make sure not error / panic",
			args: args{
				c: ctx,
			},
			fields: fields{
				cakeUsecase: mockCakeUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CakeDelivery{
				cakeUsecase: tt.fields.cakeUsecase,
			}
			d.GetCakes(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestCakeDelivery_GetCake(t *testing.T) {
	type fields struct {
		cakeUsecase usecase.CakeUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	mockCakeUsecase := &usecase.CakeUsecaseInterfaceMock{}
	mockCakeUsecase.GetDetailCakeFunc = func(ctx context.Context, id int) (*model.CakeResponse, *model.ErrorResponse) {
		return &model.CakeResponse{}, nil
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
	}
	ctx.AddParam("id", "1")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic test, make sure not error / panic",
			fields: fields{
				cakeUsecase: mockCakeUsecase,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CakeDelivery{
				cakeUsecase: tt.fields.cakeUsecase,
			}
			d.GetCake(tt.args.c)
		})
		assert.EqualValues(t, http.StatusOK, w.Code)
	}
}

func TestCakeDelivery_CreateCake(t *testing.T) {
	type fields struct {
		cakeUsecase usecase.CakeUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	mockCakeUsecase := &usecase.CakeUsecaseInterfaceMock{}
	mockCakeUsecase.CreateCakeFunc = func(ctx context.Context, payload model.CakePayloadQuery) (*model.CakeMutationResponse, *model.ErrorResponse) {
		return &model.CakeMutationResponse{}, nil
	}
	payload := model.ApiMutationCakePayload{
		Title:       "title",
		Description: null.StringFrom("description").Ptr(),
		Rating:      4.30,
		Image:       null.StringFrom("image").Ptr(),
	}
	byteJson, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Method: http.MethodPost,
		Body:   ioutil.NopCloser(bytes.NewBuffer(byteJson)),
	}
	ctx.Request.Header.Set("Content-Type", "application/json")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic test, make sure not error / panic",
			fields: fields{
				cakeUsecase: mockCakeUsecase,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CakeDelivery{
				cakeUsecase: tt.fields.cakeUsecase,
			}
			d.CreateCake(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestCakeDelivery_DeleteCake(t *testing.T) {
	type fields struct {
		cakeUsecase usecase.CakeUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	mockCakeUsecase := &usecase.CakeUsecaseInterfaceMock{}
	mockCakeUsecase.DeleteCakeFunc = func(ctx context.Context, id int) (*model.CakeDeleteResponse, *model.ErrorResponse) {
		return &model.CakeDeleteResponse{}, nil
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Method: http.MethodPost,
	}
	ctx.AddParam("id", "1")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic test, make sure not error / panic",
			fields: fields{
				cakeUsecase: mockCakeUsecase,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CakeDelivery{
				cakeUsecase: tt.fields.cakeUsecase,
			}
			d.DeleteCake(tt.args.c)
		})
	}
}

func TestCakeDelivery_UpdateCake(t *testing.T) {
	type fields struct {
		cakeUsecase usecase.CakeUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	mockCakeUsecase := &usecase.CakeUsecaseInterfaceMock{}
	mockCakeUsecase.UpdateCakeFunc = func(ctx context.Context, id int, payload model.CakePayloadQuery) (*model.CakeMutationResponse, *model.ErrorResponse) {
		return &model.CakeMutationResponse{}, nil
	}
	payload := model.ApiMutationCakePayload{
		Title:       "title",
		Description: null.StringFrom("description").Ptr(),
		Rating:      4.30,
		Image:       null.StringFrom("image").Ptr(),
	}
	byteJson, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Method: http.MethodPut,
		Body:   ioutil.NopCloser(bytes.NewBuffer(byteJson)),
	}
	ctx.AddParam("id", "1")
	ctx.Request.Header.Set("Content-Type", "application/json")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic test, make sure not error / panic",
			fields: fields{
				cakeUsecase: mockCakeUsecase,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CakeDelivery{
				cakeUsecase: tt.fields.cakeUsecase,
			}
			d.UpdateCake(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}
