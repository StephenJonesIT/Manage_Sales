package controller

import (
	"manage_sales/common"
	"manage_sales/modules/bill/biz"
	"manage_sales/modules/bill/model"
	"manage_sales/modules/bill/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func  CreateBill(db *gorm.DB) func (*gin.Context){
	return func(ctx *gin.Context){
		var requestData model.HoaDonRequest

		if err := ctx.ShouldBindJSON(&requestData); err!=nil{
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		var bill = common.HoaDon{
			MaHD: 			requestData.MaHD, 
			TongTien:		requestData.TongTien,
			MaKH:			requestData.MaKH,
			MaBaoCao:		requestData.MaBaoCao,
			MaNV:			requestData.MaNV,
		}

		storage := mysql.NewBillRepository(db)

		business := biz.InitCreateBill(storage)

		if err := business.CreateBill(ctx.Request.Context(), &bill); err != nil{
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		businessDetail := biz.InitCreateBillDetail(storage)
		for _, detail := range requestData.ChiTietHD {
			detail.MaHD = bill.MaHD
			if err := businessDetail.CreateDetailBill(ctx.Request.Context(),&detail); err != nil{
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
		} 

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(requestData))
	}
}