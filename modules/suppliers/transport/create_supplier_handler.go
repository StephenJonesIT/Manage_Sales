package transport

import (
	"log"
	"manage_sales/common"
	"manage_sales/modules/suppliers/biz"
	"manage_sales/modules/suppliers/model"
	"manage_sales/modules/suppliers/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data model.SupplierItemCreate

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
		business := biz.NewCreateSupplierBiz(store)

		if err := business.CreateNewSupplier(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())

			return
		}

		log.Println("Supplier created successfully:", data)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}