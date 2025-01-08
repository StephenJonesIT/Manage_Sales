package controller

import (
	"manage_sales/common"
	"manage_sales/modules/employee/biz"
	"manage_sales/modules/employee/model"
	"manage_sales/modules/employee/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateEmployee(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.EmployeeUpdate
		id := strings.ToUpper(ctx.Param("manv"))

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := mysql.NewSQLStore(db)
		business := biz.NewUpdateEmployeeBiz(store)

		if err := business.UpdateEmployeeById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
