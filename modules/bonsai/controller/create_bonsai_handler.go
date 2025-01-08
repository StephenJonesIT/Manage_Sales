package controller

import (
	"log"
	"manage_sales/common"
	"manage_sales/modules/bonsai/biz"
	"manage_sales/modules/bonsai/model"
	"manage_sales/modules/bonsai/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data model.BonsaiItemCreate

		if err := c.ShouldBind(&data); err != nil {
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

		store := mysql.NewSQLStore(db)
		business := biz.NewCreateBonsaiBiz(store)

		if err := business.CreateNewBonsai(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotCreateEntity("Product",err))
			return
		}

		log.Println("Bonsai created successfully:", data)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}