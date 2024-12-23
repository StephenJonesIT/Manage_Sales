package transport

import (
	"manage_sales/common"
	"manage_sales/modules/customer/biz"
	"manage_sales/modules/customer/storage/mysql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCustomer(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context){

		id := strings.ToUpper(ctx.Param("makh"))

		store := mysql.NewSQLStore(db)

		business := biz.NewGetCustomerBiz(store)

		data, err := business.GetCustomerById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}