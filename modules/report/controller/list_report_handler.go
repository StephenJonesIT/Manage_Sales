package controller

import (
	"manage_sales/common"
	"manage_sales/modules/report/biz"
	stora "manage_sales/modules/report/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListReport(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err!=nil{
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging.Process()

		storage := stora.NewSQLReport(db)
		business := biz.NewReportStorage(storage)
		data, err := business.ListReport(ctx.Request.Context(),&paging);

		if err != nil {
			ctx.JSON(http.StatusInternalServerError,common.ErrIntenal(err))
			return
		}

		ctx.JSON(http.StatusOK,common.NewSuccessResponseSupplier(data,paging))
	}
}