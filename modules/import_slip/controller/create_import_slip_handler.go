package controller

import (
	"log"
	"manage_sales/common"
	"manage_sales/modules/import_slip/business"
	"manage_sales/modules/import_slip/model"
	"manage_sales/modules/import_slip/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateImportSlip(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var newPhieuNhap model.PhieuNhapCreate
		if err := ctx.ShouldBind(&newPhieuNhap); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return 
		}

		storage := mysql.NewSQLStore(db)
		business := business.NewCreateImportSlip(storage)
		if err := business.CreateImportSlip(ctx.Request.Context(),&newPhieuNhap); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.ErrIntenal(err))
			return
		}

		log.Println("Import Slip created successfully!")
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}