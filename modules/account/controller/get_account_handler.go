package controller

import (
	"manage_sales/common"
	"manage_sales/modules/account/biz"
	"manage_sales/modules/account/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAccount(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {

		id := strings.ToUpper(ctx.Param("matk"))
	

		store := mysql.NewAccountRepository(db)

		business := biz.NewGetAccountBiz(store)

		data, err := business.FindAccountById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}