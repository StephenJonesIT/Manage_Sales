package controller

import (
	"manage_sales/common"
	"manage_sales/modules/bill/business"
	"manage_sales/modules/bill/model"
	"manage_sales/modules/bill/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBill(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var id = ctx.Param("mahd")
		storage := mysql.NewBillRepository(db)
		business := business.NewGetBill(storage)
		data, err := business.GetBill(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrCannotGetEntity(model.Entity,err))
			return
		}
		ctx.JSON(http.StatusOK,common.SimpleSuccessResponse(data))
	}
}