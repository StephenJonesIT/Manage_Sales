package transport

import (
	"manage_sales/common"
	"manage_sales/modules/suppliers/biz"
	"manage_sales/modules/suppliers/storage/mysql"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItems(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		paging.Process()

	
		store := mysql.NewSQLStore(db)
		business := biz.NewListSupplierBiz(store)

		result, err := business.ListSupplierById(ctx.Request.Context(), &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponseSupplier(result,paging))
	}
}