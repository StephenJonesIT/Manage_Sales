package controller

import (
	"manage_sales/common"
	"manage_sales/modules/bonsai/biz"
	"manage_sales/modules/bonsai/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		id := strings.ToUpper(ctx.Param("masp"))
	

		store := mysql.NewSQLStore(db)

		business := biz.NewGetBonsaiBiz(store)

		data, err := business.GetBonsaiById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}