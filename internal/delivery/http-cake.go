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

// GetCakes godoc
//
//	@Summary	GetCakes
//	@Tags		cakes
//	@Param		page		query	integer	false	"default page is at page 1"
//	@Param		page_size	query	integer	true	"maximum value is 100"
//	@Produce	json
//	@Success	200	{object}	model.GetCakesResponse
//	@Router		/cakes [get]
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
	if query.Page <= 0 {
		query.Page = 1
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

// GetCake godoc
//
//	@Summary	GetCake
//	@Tags		cakes
//	@Param		id	path	string	true	"param id (cake record)"
//	@Produce	json
//	@Success	200	{object}	model.CakeResponse
//	@Router		/cakes/{id} [get]
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

// CreateCake godoc
//
//	@Summary	CreateCake
//	@Tags		cakes
//	@Param		data	body	model.ApiMutationCakePayload	true	"body data".
//	@Produce	json
//	@Success	200	{object}	model.CakeMutationResponse
//	@Router		/cakes [post]
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

// DeleteCake godoc
//
//	@Summary	DeleteCake
//	@Tags		cakes
//	@Param		id	path	string	true	"param id (cake record)"
//	@Produce	json
//	@Success	200	{object}	model.CakeDeleteResponse
//	@Router		/cakes/{id} [delete]
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

// UpdateCake godoc
//
//	@Summary	UpdateCake
//	@Tags		cakes
//	@Param		id		path	string							true	"param id (cake record)"
//	@Param		data	body	model.ApiMutationCakePayload	true	"body data".
//	@Produce	json
//	@Success	200	{object}	model.CakeMutationResponse
//	@Router		/cakes/{id} [put]
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
