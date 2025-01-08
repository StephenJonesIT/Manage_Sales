package controller

import (
	"manage_sales/common"
	"manage_sales/modules/customer/biz"
	"manage_sales/modules/customer/model"
	"manage_sales/modules/customer/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.CustomerUpdate
		id := strings.ToUpper(ctx.Param("makh"))

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := mysql.NewSQLStore(db)
		business := biz.NewUpdateCustomerBiz(store)

		if err := business.UpdateCustomerById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
