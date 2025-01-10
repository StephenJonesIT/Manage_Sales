package controller

import (
	"manage_sales/common"
	"manage_sales/modules/import_slip/storage/mysql"
	"manage_sales/modules/import_slip/business"
	"manage_sales/modules/import_slip/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetImportSlip(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var id = ctx.Param("mapn")
		storage := mysql.NewSQLStore(db)
		business := business.NewGetImportSlip(storage)
		data, err := business.GetImportSlip(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrCannotGetEntity(model.EntitySlip,err))
			return
		}
		ctx.JSON(http.StatusOK,common.SimpleSuccessResponse(data))
	}
}