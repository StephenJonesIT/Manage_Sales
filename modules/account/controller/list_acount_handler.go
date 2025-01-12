package controller

import (
	"manage_sales/common"
	"manage_sales/modules/account/biz"
	"manage_sales/modules/account/storage/mysql"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAccount(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		paging.Process()

	
		store := mysql.NewAccountRepository(db)
		business := biz.NewListAccountBiz(store)

		result, err := business.ListAccountById(ctx.Request.Context(), &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponseSupplier(result,paging))
	}
}