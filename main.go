package main

import (
	"gin_test/controller"
	"gin_test/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.POST("/add", controller.UserAdd)
			v1.GET("/list", controller.UserList)
			v1.PUT("/update", controller.UserUpdate)
			v1.DELETE("/delete", controller.UserDelete)
		}
	}
	engine.Run(":3000")
}
