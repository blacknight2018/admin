package main

import (
	"admin/Config"
	"admin/Service/Goods"
	"admin/Service/Order"
	"admin/Service/User"
	"admin/Utils"
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
	user := r.Group("/user")
	goods := r.Group("/goods")
	goods.GET("", func(context *gin.Context) {
		goodsTitle := context.Query("title")
		limit := Utils.StrToInt(context.Query("limit"))
		offset := Utils.StrToInt(context.Query("offset"))
		context.Writer.Write([]byte(Goods.QueryGoods(goodsTitle, limit, offset).Get()))
	})
	user.GET("", func(context *gin.Context) {
		nickName := context.Query("nick_name")
		limit := Utils.StrToInt(context.Query("limit"))
		offset := Utils.StrToInt(context.Query("offset"))
		context.Writer.Write([]byte(User.QueryUser(nickName, limit, offset).Get()))
	})
	order.GET("/query", func(context *gin.Context) {
		context.Writer.Write([]byte(Order.QueryOrder().Get()))
	})
	r.Run(":45678")
}
