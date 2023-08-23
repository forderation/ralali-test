package delivery

import (
	"net/http"
	"strconv"

	"github.com/forderation/ralali-test/internal/model"
	"github.com/forderation/ralali-test/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CakeDelivery struct {
	cakeUsecase usecase.CakeUsecaseInterface
}

func NewCakeDelivery(cakeUsecase usecase.CakeUsecaseInterface) *CakeDelivery {
	return &CakeDelivery{
		cakeUsecase: cakeUsecase,
	}
}

func (d *CakeDelivery) GetCakes(c *gin.Context) {
	ctx := c.Request.Context()
	var query model.ApiGetCakesQuery
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "invalid query parameter:" + err.Error()})
		return
	}
	if query.PageSize > 100 {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "maximum page_size is 100"})
		return
	}
	response, errResponse := d.cakeUsecase.GetCakes(ctx, model.GetCakesUsecaseParam{
		Page:     query.Page,
		PageSize: query.PageSize,
	})
	if errResponse != nil {
		c.JSON(errResponse.HttpStatusCode, model.JsonErrorResp{ErrorMessage: errResponse.Err.Error(), ErrData: errResponse.ErrData})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func (d *CakeDelivery) GetCake(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "invalid parameter id"})
		return
	}
	response, errResponse := d.cakeUsecase.GetDetailCake(ctx, id)
	if errResponse != nil {
		c.JSON(errResponse.HttpStatusCode, model.JsonErrorResp{ErrorMessage: errResponse.Err.Error(), ErrData: errResponse.ErrData})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func (d *CakeDelivery) CreateCake(c *gin.Context) {
	ctx := c.Request.Context()
	var payload model.ApiMutationCakePayload
	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: err.Error()})
		return
	}
	err = payload.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "invalid payload: " + err.Error()})
		return
	}
	response, errResponse := d.cakeUsecase.CreateCake(ctx, model.CakePayloadQuery{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
	})
	if errResponse != nil {
		c.JSON(errResponse.HttpStatusCode, model.JsonErrorResp{ErrorMessage: errResponse.Err.Error(), ErrData: errResponse.ErrData})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func (d *CakeDelivery) DeleteCake(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "invalid parameter id"})
		return
	}
	response, errResponse := d.cakeUsecase.DeleteCake(ctx, id)
	if errResponse != nil {
		c.JSON(errResponse.HttpStatusCode, model.JsonErrorResp{ErrorMessage: errResponse.Err.Error(), ErrData: errResponse.ErrData})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func (d *CakeDelivery) UpdateCake(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "invalid parameter id"})
		return
	}
	var payload model.ApiMutationCakePayload
	err = c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: err.Error()})
		return
	}
	err = payload.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.JsonErrorResp{ErrorMessage: "invalid payload: " + err.Error()})
		return
	}
	response, errResponse := d.cakeUsecase.UpdateCake(ctx, id, model.CakePayloadQuery{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
	})
	if errResponse != nil {
		c.JSON(errResponse.HttpStatusCode, model.JsonErrorResp{ErrorMessage: errResponse.Err.Error(), ErrData: errResponse.ErrData})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
