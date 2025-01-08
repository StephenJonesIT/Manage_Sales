package controller

import (
	"fmt"
	"manage_sales/common"
	"manage_sales/modules/employee/biz"
	"manage_sales/modules/employee/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteEmployee(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		manv := ctx.Param("manv")

		id := strings.ToUpper(manv)

		fmt.Println(id)

		store := mysql.NewSQLStore(db)
		business := biz.NewDeleteEmployeeBiz(store)

		if err := business.DeleteEmployeeById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
