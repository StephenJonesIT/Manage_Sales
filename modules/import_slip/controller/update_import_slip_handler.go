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

func UpdateImportSlip(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var id = ctx.Param("mapn")
		var updateImportSlip model.PhieuNhapUpdate

		if err := ctx.ShouldBind(&updateImportSlip); err != nil{
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		storage := mysql.NewSQLStore(db)
		business := business.NewUpdateImportSlip(storage)
		if err:= business.UpdateImportSlip(ctx.Request.Context(),id, &updateImportSlip); err!=nil {
			ctx.JSON(http.StatusInternalServerError, common.ErrIntenal(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}