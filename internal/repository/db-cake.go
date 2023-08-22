package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/forderation/ralali-test/internal/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

const (
	COUNT_CAKES_STMT int = iota
	GET_CAKES_STMT
	GET_CAKE_STMT
	INSERT_CAKE_STMT
	UPDATE_CAKE_STMT
	SOFT_DELETE_CAKE_STMT
)

type CakeDBRepository struct {
	db            *sql.DB
	queryPrepared map[int]*sql.Stmt
}

func NewCakeDBRepository(db *sql.DB, tableName string) *CakeDBRepository {
	if db == nil {
		logrus.Panic("db param for NewCakeDBRepository is nil")
	}
	queryPrepared := make(map[int]*sql.Stmt, 0)
	sqlStmtGetCakes, err := db.Prepare(fmt.Sprintf("SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM %s WHERE deleted_at IS NULL ORDER BY rating DESC, title ASC LIMIT ? OFFSET ?", tableName))
	if err != nil {
		logrus.Panic("error execute prepared statement sqlStmtGetCakes : ", err)
	}
	sqlStmtInsertCake, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", tableName))
	if err != nil {
		logrus.Panic("error execute prepared statement sqlStmtInsertCake : ", err)
	}
	sqlStmtUpdateCake, err := db.Prepare(fmt.Sprintf("UPDATE %s SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?", tableName))
	if err != nil {
		logrus.Panic("error execute prepared statement sqlStmtUpdateCake : ", err)
	}
	sqlStmtSoftDeleteCake, err := db.Prepare(fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", tableName))
	if err != nil {
		logrus.Panic("error execute prepared statement sqlStmtSoftDeleteCake : ", err)
	}
	sqlStmtCountCakes, err := db.Prepare(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE deleted_at IS NULL", tableName))
	if err != nil {
		logrus.Panic("error execute prepared statement sqlStmtCountCakes : ", err)
	}
	sqlStmtGetCake, err := db.Prepare(fmt.Sprintf("SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM %s WHERE deleted_at IS NULL AND id = ? LIMIT 1", tableName))
	if err != nil {
		logrus.Panic("error execute prepared statement sqlStmtGetCakes : ", err)
	}
	queryPrepared[GET_CAKE_STMT] = sqlStmtGetCake
	queryPrepared[GET_CAKES_STMT] = sqlStmtGetCakes
	queryPrepared[INSERT_CAKE_STMT] = sqlStmtInsertCake
	queryPrepared[UPDATE_CAKE_STMT] = sqlStmtUpdateCake
	queryPrepared[COUNT_CAKES_STMT] = sqlStmtCountCakes
	queryPrepared[SOFT_DELETE_CAKE_STMT] = sqlStmtSoftDeleteCake
	return &CakeDBRepository{
		db:            db,
		queryPrepared: queryPrepared,
	}
}

func (repo *CakeDBRepository) GetCakes(ctx context.Context, param model.GetCakesQuery) ([]model.Cake, error) {
	stmt := repo.queryPrepared[GET_CAKES_STMT]
	rows, err := stmt.QueryContext(ctx, param.Limit, param.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []model.Cake{}
	for rows.Next() {
		var cake model.Cake
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, cake)
	}
	return result, nil
}

func (repo *CakeDBRepository) CountCakes(ctx context.Context) (int64, error) {
	stmt := repo.queryPrepared[COUNT_CAKES_STMT]
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var result int64
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return 0, err
		}
	}
	return result, nil
}

func (repo *CakeDBRepository) GetCake(ctx context.Context, id int) (*model.Cake, error) {
	stmt := repo.queryPrepared[GET_CAKE_STMT]
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []model.Cake{}
	for rows.Next() {
		var cake model.Cake
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, cake)
	}
	if len(result) > 0 {
		return &result[0], nil
	}
	return nil, nil
}

func (repo *CakeDBRepository) InsertCake(ctx context.Context, param model.CakePayloadQuery) error {
	stmt := repo.queryPrepared[INSERT_CAKE_STMT]
	timeCreated := time.Now().UTC()
	_, err := stmt.ExecContext(ctx, param.Title, param.Description, param.Rating, param.Image, timeCreated, timeCreated)
	return err
}

func (repo *CakeDBRepository) UpdateCake(ctx context.Context, id int, param model.CakePayloadQuery) error {
	stmt := repo.queryPrepared[UPDATE_CAKE_STMT]
	timeUpdated := time.Now().UTC()
	_, err := stmt.ExecContext(ctx, param.Title, param.Description, param.Rating, param.Image, timeUpdated, id)
	return err
}

func (repo *CakeDBRepository) SoftDeleteCake(ctx context.Context, id int) error {
	stmt := repo.queryPrepared[SOFT_DELETE_CAKE_STMT]
	timeDeleted := time.Now().UTC()
	_, err := stmt.ExecContext(ctx, timeDeleted, id)
	return err
}
