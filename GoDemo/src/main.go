package main
import (
   "github.com/gin-gonic/gin" 
   "database/sql"
   "log"
 controller "GODEMO/src/controllers"
   _ "github.com/go-sql-driver/mysql"
)
func main() {

    router := gin.Default()
    router.GET("/select",controller.Select)
    router.POST("/create",controller.create)
	router.DELETE("/delete",controller.Delete)
	router.PUT("/update",controller.Update)
   router.Run()
}
