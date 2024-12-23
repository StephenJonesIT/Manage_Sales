package transport

import (
	"manage_sales/common"
	"manage_sales/modules/suppliers/biz"
	"manage_sales/modules/suppliers/model"
	"manage_sales/modules/suppliers/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.SupplierItemUpdate
		id := strings.ToUpper(ctx.Param("mancc"))

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := mysql.NewSQLStore(db)
		business := biz.NewUpdateSupplierBiz(store)

		if err := business.UpdateSupplierById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
