package main

import (
	"admin/Config"
	"admin/Service/Order"
	"github.com/gin-gonic/gin"
)

func main() {
	Config.GetConf()
	r := gin.Default()
	r.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		context.Next()
	})
	order := r.Group("/order")
	order.GET("/query", func(context *gin.Context) {
		context.Writer.Write([]byte(Order.QueryOrder().Get()))
	})
	r.Run(":45678")
}
