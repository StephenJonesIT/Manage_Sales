package controller

import (
	"manage_sales/common"
	"manage_sales/modules/bill/business"
	"manage_sales/modules/bill/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteBill(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var id = ctx.Param("mahd")
		storage := mysql.NewBillRepository(db)
		business := business.NewDeleteBill(storage)
		if err := business.DeleteBillBiz(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}