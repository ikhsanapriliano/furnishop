package controller

import (
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

func (p *ProductCategoryController) createHandler(ctx *gin.Context) {
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

func (p *ProductCategoryController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := p.uc.GetById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "success", response)
}

func (p *ProductCategoryController) getAllHandler(ctx *gin.Context) {
	response, err := p.uc.GetAll()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "success", response)
}

func (p *ProductCategoryController) updateHandler(ctx *gin.Context) {
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

func (p *ProductCategoryController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := p.uc.Delete(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "success", response)
}

func (p *ProductCategoryController) Router() {
	pg := p.rg.Group(appconfig.ProductCategoryGroup)
	{
		pg.POST(appconfig.ProductCategoryCreate, p.createHandler)
		pg.GET(appconfig.ProductCategoryGetById, p.getByIdHandler)
		pg.GET(appconfig.ProductCategoryGetAll, p.getAllHandler)
		pg.PUT(appconfig.ProductCategoryUpdate, p.updateHandler)
		pg.DELETE(appconfig.ProductCategoryDelete, p.deleteHandler)
	}
}

func NewProductCategoryController(uc usecase.ProductCategoryUseCase, rg *gin.RouterGroup) *ProductCategoryController {
	return &ProductCategoryController{
		uc: uc,
		rg: rg,
	}
}
