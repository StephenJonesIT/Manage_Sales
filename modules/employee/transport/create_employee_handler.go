package transport

import (
	"log"
	"manage_sales/common"
	"manage_sales/modules/employee/biz"
	"manage_sales/modules/employee/model"
	"manage_sales/modules/employee/storage/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEmployee(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var data model.EmployeeCreate
		if err := ctx.Bind(&data); err != nil{
			log.Println("Error binding data:", err)

			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		log.Println("Data bound successfully:", data)

		if db == nil {
			log.Println("DB is nil")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database connection is not initialized",
			})
			return
		}

		storage := mysql.NewSQLStore(db)
		business := biz.NewCreateEmployeeBiz(storage)

		if err := business.CreateNewEmployee(ctx.Request.Context(), &data); err!=nil{
			ctx.JSON(http.StatusBadRequest, err)
		}

		log.Println("Employee created successfully!")
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}