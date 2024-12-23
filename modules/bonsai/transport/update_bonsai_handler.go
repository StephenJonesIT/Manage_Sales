package transport

import (
	"manage_sales/common"
	"manage_sales/modules/bonsai/biz"
	"manage_sales/modules/bonsai/model"
	"manage_sales/modules/bonsai/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.BonsaiItemUpdate
		id := strings.ToUpper(ctx.Param("masp"))

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := mysql.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
