package controller

import (
	"manage_sales/common"
	"manage_sales/modules/import_slip/business"
	"manage_sales/modules/import_slip/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteImportSlip(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var id = ctx.Param("mapn")
		storage := mysql.NewSQLStore(db)
		business := business.NewDeleteImportSlip(storage)
		if err := business.DeleteImportSlipBiz(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}