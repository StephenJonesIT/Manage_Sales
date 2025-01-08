package controller

import (
	"fmt"
	"manage_sales/common"
	"manage_sales/modules/bonsai/biz"
	"manage_sales/modules/bonsai/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		macc := ctx.Param("masp")

		id := strings.ToUpper(macc)

		fmt.Println(id)
		
		store := mysql.NewSQLStore(db)
		business := biz.NewDeleteItemBiz(store)

		if err := business.DeleteItemById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
