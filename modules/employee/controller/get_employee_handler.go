package controller

import (
	"manage_sales/common"
	"manage_sales/modules/employee/biz"
	"manage_sales/modules/employee/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		id := strings.ToUpper(ctx.Param("manv"))

		store := mysql.NewSQLStore(db)

		business := biz.NewGetEmployeeBiz(store)

		data, err := business.GetEmployeeById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}