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

func UpdateBill(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var id = ctx.Param("mapn")
		var updateBill model.UpdateHoaDon

		if err := ctx.ShouldBind(&updateBill); err != nil{
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		storage := mysql.NewBillRepository(db)
		business := business.NewUpdateBill(storage)
		if err:= business.UpdateBill(ctx.Request.Context(),id, &updateBill); err!=nil {
			ctx.JSON(http.StatusInternalServerError, common.ErrIntenal(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}