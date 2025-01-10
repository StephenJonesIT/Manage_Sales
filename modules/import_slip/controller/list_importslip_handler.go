package controller

import (
	"manage_sales/common"
	"manage_sales/modules/import_slip/business"
	"manage_sales/modules/import_slip/model"
	"manage_sales/modules/import_slip/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListEmployees(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		paging.Process()

		var filter model.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		store := mysql.NewSQLStore(db)
		business := business.NewListImportSlip(store)

		result, err := business.ListImportSlips(ctx, &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}