package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/forderation/ralali-test/internal/model"
	"gopkg.in/guregu/null.v4"
)

func TestNewCakeDBRepository(t *testing.T) {
	tableName := "cakes"
	db, _ := InitTestDB(tableName)
	defer db.Close()
	type args struct {
		db        *sql.DB
		tableName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "make sure prepared query already meet expectation",
			args: args{
				db:        db,
				tableName: tableName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewCakeDBRepository(tt.args.db, tt.args.tableName)
		})
	}
}

func InitTestDB(tableName string) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectPrepare(regexp.QuoteMeta(fmt.Sprintf("SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM %s WHERE deleted_at IS NULL ORDER BY rating DESC, title ASC LIMIT ? OFFSET ?", tableName)))
	mock.ExpectPrepare(regexp.QuoteMeta(fmt.Sprintf("INSERT INTO %s (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", tableName)))
	mock.ExpectPrepare(regexp.QuoteMeta(fmt.Sprintf("UPDATE %s SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?", tableName)))
	mock.ExpectPrepare(regexp.QuoteMeta(fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", tableName)))
	mock.ExpectPrepare(regexp.QuoteMeta(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE deleted_at IS NULL", tableName)))
	mock.ExpectPrepare(regexp.QuoteMeta(fmt.Sprintf("SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM %s WHERE deleted_at IS NULL AND id = ? LIMIT 1", tableName)))
	return db, mock
}

func TestCakeDBRepository_GetCakes(t *testing.T) {
	tableName := "cakes"
	db, mock := InitTestDB(tableName)
	timeMock, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	if err != nil {
		log.Fatalf("an error '%s' was not expected", err)
	}
	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"})
	rows.AddRow(1, "title", "desc", float32(4.5), "image", timeMock, timeMock, nil)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM cakes WHERE deleted_at IS NULL ORDER BY rating DESC, title ASC LIMIT ? OFFSET ?")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)
	repo := NewCakeDBRepository(db, tableName)
	type args struct {
		ctx   context.Context
		param model.GetCakesQuery
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Cake
		wantErr bool
	}{
		{
			name: "basic test get cakes",
			args: args{
				ctx: context.TODO(),
				param: model.GetCakesQuery{
					Limit:  10,
					Offset: 1,
				},
			},
			want: []model.Cake{
				{
					ID:          1,
					Title:       "title",
					Description: null.StringFrom("desc").Ptr(),
					Rating:      4.5,
					Image:       null.StringFrom("image").Ptr(),
					CreatedAt:   timeMock,
					UpdatedAt:   timeMock,
					DeletedAt:   nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetCakes(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeDBRepository.GetCakes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeDBRepository.GetCakes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeDBRepository_CountCakes(t *testing.T) {
	tableName := "cakes"
	db, mock := InitTestDB(tableName)
	rows := sqlmock.NewRows([]string{"count"})
	rows.AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM cakes WHERE deleted_at IS NULL")).WillReturnRows(rows)
	repo := NewCakeDBRepository(db, tableName)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				ctx: context.TODO(),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CountCakes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeDBRepository.CountCakes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CakeDBRepository.CountCakes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeDBRepository_GetCake(t *testing.T) {
	tableName := "cakes"
	db, mock := InitTestDB(tableName)
	timeMock, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	if err != nil {
		log.Fatalf("an error '%s' was not expected", err)
	}
	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"})
	rows.AddRow(1, "title", "desc", float32(4.5), "image", timeMock, timeMock, nil)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM cakes WHERE deleted_at IS NULL AND id = ? LIMIT 1")).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)
	repo := NewCakeDBRepository(db, tableName)
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Cake
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			want: &model.Cake{
				ID:          1,
				Title:       "title",
				Description: null.StringFrom("desc").Ptr(),
				Rating:      4.5,
				Image:       null.StringFrom("image").Ptr(),
				CreatedAt:   timeMock,
				UpdatedAt:   timeMock,
				DeletedAt:   nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetCake(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CakeDBRepository.GetCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CakeDBRepository.GetCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCakeDBRepository_InsertCake(t *testing.T) {
	tableName := "cakes"
	db, mock := InitTestDB(tableName)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")).WithArgs(
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
	repo := NewCakeDBRepository(db, tableName)
	type args struct {
		ctx   context.Context
		param model.CakePayloadQuery
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				ctx:   context.TODO(),
				param: model.CakePayloadQuery{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.InsertCake(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("CakeDBRepository.InsertCake() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCakeDBRepository_UpdateCake(t *testing.T) {
	tableName := "cakes"
	db, mock := InitTestDB(tableName)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?")).WithArgs(
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
	repo := NewCakeDBRepository(db, tableName)
	type args struct {
		ctx   context.Context
		id    int
		param model.CakePayloadQuery
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				ctx:   context.TODO(),
				id:    1,
				param: model.CakePayloadQuery{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.UpdateCake(tt.args.ctx, tt.args.id, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("CakeDBRepository.UpdateCake() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCakeDBRepository_SoftDeleteCake(t *testing.T) {
	tableName := "cakes"
	db, mock := InitTestDB(tableName)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE cakes SET deleted_at = ? WHERE id = ?")).WithArgs(
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
	repo := NewCakeDBRepository(db, tableName)
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.SoftDeleteCake(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CakeDBRepository.SoftDeleteCake() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
