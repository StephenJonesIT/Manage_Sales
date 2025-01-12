package controller

import (
	"manage_sales/common"
	"manage_sales/modules/account/biz"
	"manage_sales/modules/account/storage/mysql"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAccount(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("matk")

		
		store := mysql.NewAccountRepository(db)
		business := biz.NewDeleteAccountBiz(store)

		if err := business.DeleteAccount(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
