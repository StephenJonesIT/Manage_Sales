package transport

import (
	"fmt"
	"manage_sales/common"
	"manage_sales/modules/suppliers/biz"
	"manage_sales/modules/suppliers/storage/mysql"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		macc := ctx.Param("mancc")

		id := strings.ToUpper(macc)

		fmt.Println(id)
		
		store := mysql.NewSQLStore(db)
		business := biz.NewDeleteSupplierBiz(store)

		if err := business.DeleteSupplierById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
