package configuration

import (
	ginitem "manage_sales/modules/bonsai/transport"
	gincustomer "manage_sales/modules/customer/transport"
	ginemployee "manage_sales/modules/employee/transport"
	ginsupplier "manage_sales/modules/suppliers/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHttp() {
  r := gin.Default()

 createBonsaiHttp(r)
 createSupplierHttp(r)
 createCustomerHttp(r)
 createEmployeeHttp(r)

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
 

  r.Run(":1000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createBonsaiHttp(r *gin.Engine){
      v1 := r.Group("/shop")
      {
        bonsais := v1.Group("/bonsais")
        {
            bonsais.POST("", ginitem.CreateItem(DB))
            bonsais.GET("", ginitem.ListItems(DB))
            bonsais.GET("/:masp",ginitem.GetItem(DB))
            bonsais.DELETE("/:masp", ginitem.DeleteItem(DB))
            bonsais.PATCH("/:masp",ginitem.UpdateItem(DB))
        }
      }
}

func createSupplierHttp(r *gin.Engine){
    v1 := r.Group("/shop")
    {
      suppliers := v1.Group("/suppliers")
      {
        suppliers.POST("", ginsupplier.CreateItem(DB))
        suppliers.GET("", ginsupplier.ListItems(DB))
        suppliers.GET("/:mancc", ginsupplier.GetItem(DB))
        suppliers.DELETE("/:mancc", ginsupplier.DeleteItem(DB))
        suppliers.PATCH("/:mancc", ginsupplier.UpdateItem(DB))
      }
    }
}

func createCustomerHttp(r *gin.Engine){
   v1 := r.Group("/shop")
   {
      customers := v1.Group("/customer")
      {
        customers.POST("",gincustomer.CreateItem(DB))
        customers.GET("", gincustomer.ListItems(DB))
        customers.GET("/:makh", gincustomer.GetCustomer(DB))
        customers.PATCH("/:makh", gincustomer.UpdateItem(DB))
      }
   }
}

func createEmployeeHttp(r *gin.Engine){
  v1 := r.Group("/shop")
  {
    employees := v1.Group("/employee")
    {
      employees.GET("/:manv", ginemployee.GetItem(DB))
      employees.GET("", ginemployee.ListEmployees(DB))
      employees.POST("", ginemployee.CreateEmployee(DB))
      employees.DELETE("/:manv", ginemployee.DeleteEmployee(DB))
      employees.PATCH("/:manv", ginemployee.UpdateEmployee(DB))
    }
  }
}