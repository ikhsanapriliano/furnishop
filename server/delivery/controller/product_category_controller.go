package controller

import (
	"fmt"
	appconfig "furnishop/server/config/app_config"
	"furnishop/server/dto"
	"furnishop/server/usecase"
	"furnishop/server/util/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductCategoryController struct {
	uc usecase.ProductCategoryUseCase
	rg *gin.RouterGroup
}

func (p *ProductCategoryController) CreateHandler(ctx *gin.Context) {
	var payload dto.ProductCategoryDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := p.uc.Create(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "success", response)
}

func (p *ProductCategoryController) GetByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := p.uc.GetById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "success", response)
}

func (p *ProductCategoryController) GetAllHandler(ctx *gin.Context) {
	response, err := p.uc.GetAll()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "success", response)
}

func (p *ProductCategoryController) UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload dto.ProductCategoryDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := p.uc.Update(id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "success", response)
}

func (p *ProductCategoryController) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := p.uc.Delete(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := fmt.Sprintf("product category with id %s deleted", id)
	common.SendSingleResponse(ctx, "success", response)
}

func (p *ProductCategoryController) Router() {
	pg := p.rg.Group(appconfig.ProductCategoryGroup)
	{
		pg.POST(appconfig.ProductCategoryCreate, p.CreateHandler)
		pg.GET(appconfig.ProductCategoryGetById, p.GetByIdHandler)
		pg.GET(appconfig.ProductCategoryGetAll, p.GetAllHandler)
		pg.PUT(appconfig.ProductCategoryUpdate, p.UpdateHandler)
		pg.DELETE(appconfig.ProductCategoryDelete, p.DeleteHandler)
	}
}

func NewProductCategoryController(uc usecase.ProductCategoryUseCase, rg *gin.RouterGroup) *ProductCategoryController {
	return &ProductCategoryController{
		uc: uc,
		rg: rg,
	}
}
