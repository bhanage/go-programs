package main

import (
  
  "github.com/gin-gonic/gin"
   controller "example/web-service-gin/controller"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  // r := gin.Default()
  // r.GET("/select", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{
  //     "message": "pong",
  //   })
  // })

  // r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  router := gin.Default()
  router.GET("/select",controller.Select)
  router.POST("/create",controller.Create)
router.DELETE("/delete",controller.Delete)
router.PUT("/update",controller.Update)
 router.Run()
}