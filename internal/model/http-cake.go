package model

type ApiGetCakesQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size" binding:"required"`
}
