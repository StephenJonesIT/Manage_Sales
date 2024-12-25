package controller

import (
	"log"
	"manage_sales/common"
	"manage_sales/modules/account/biz"
	"manage_sales/modules/account/model"
	"manage_sales/modules/account/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data model.AccountCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			log.Println("Error binding data:", err)

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		log.Println("Data bound successfully:", data)

		if db == nil {
			log.Println("DB is nil")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database connection is not initialized",
			})
			return
		}

		hashedPassword ,err := common.HashPassword(data.MatKhau)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrIntenal(err))
		}

		data.MatKhau = hashedPassword
		
		store := mysql.NewAccountRepository(db)
		business := biz.NewAccount(store)

		if err := business.CreateAccountBiz(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		log.Println("Account created successfully:", data)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}