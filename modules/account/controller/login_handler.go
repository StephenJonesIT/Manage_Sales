package controller

import (
	"manage_sales/common"
	"manage_sales/modules/account/biz"
	"manage_sales/modules/account/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginController(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var creds common.LoginRequest

		if err := ctx.ShouldBindJSON(&creds); err != nil{
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := mysql.NewAccountRepository(db)

		business := biz.NewGetAccountBiz(storage)

		data, err := business.FindAccountByUsername(ctx, creds.TenDangNhap)	
		
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.ErrUnauthorized(err))
			return 
		}

		if !common.CheckPasswordHash(creds.MatKhau, data.MatKhau){
			ctx.JSON(http.StatusUnauthorized, common.ErrUnauthorized(err))
			return 
		}
		role := data.Loai.String()
		
		token, err := common.GenerateJWT(data.TenDangNhap, role)

		if err != nil{
			ctx.JSON(http.StatusInternalServerError, common.ErrIntenal(err))
			return 
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponseToken(token))
	}
}