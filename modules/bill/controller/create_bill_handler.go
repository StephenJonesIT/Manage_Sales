package controller

import (
	"log"
	"manage_sales/common"
	"manage_sales/modules/bill/business"
	"manage_sales/modules/bill/model"
	"manage_sales/modules/bill/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBill(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var newHoaDon model.CreateHoaDon
		if err := ctx.ShouldBind(&newHoaDon); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return 
		}

		storage := mysql.NewBillRepository(db)
		business := business.NewCreateBill(storage)
		if err := business.CreateBill(ctx.Request.Context(),&newHoaDon); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.ErrIntenal(err))
			return
		}

		log.Println("Bill created successfully!")
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}