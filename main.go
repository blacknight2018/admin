package main

import (
	"admin/Config"
	"admin/Service/Goods"
	"admin/Service/Order"
	"admin/Service/User"
	"admin/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	Config.GetConf()
	r := gin.Default()
	r.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		context.Writer.Header().Set("Access-Control-Max-Age", "3600")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
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
	goods.OPTIONS("", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	goods.POST("", func(context *gin.Context) {
		type name struct {
			Title     string   `json:"title"`
			Desc      string   `json:"desc"`
			Template  string   `json:"template"`
			Banner    []string `json:"banner"`
			DetailImg []string `json:"detail_img"`
		}
		var tmp name
		context.ShouldBindJSON(&tmp)
		context.Writer.Write([]byte(Goods.AddGoods(tmp.Title, tmp.Desc, tmp.Template, tmp.Banner, tmp.DetailImg).Get()))
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
