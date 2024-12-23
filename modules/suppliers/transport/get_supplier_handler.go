package transport

import (
	"manage_sales/common"
	"manage_sales/modules/suppliers/biz"
	"manage_sales/modules/suppliers/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		id := strings.ToUpper(ctx.Param("mancc"))

		store := mysql.NewSQLStore(db)

		business := biz.NewGetSupplierBiz(store)

		data, err := business.GetSupplierById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}