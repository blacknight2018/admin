package main

import (
	"admin/Config"
	"admin/Service/Banner"
	"admin/Service/Goods"
	"admin/Service/Order"
	"admin/Service/SubGoods"
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
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE,PUT")
		context.Writer.Header().Set("Access-Control-Max-Age", "3600")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		context.Next()
	})

	order := r.Group("/order")
	user := r.Group("/user")
	goods := r.Group("/goods")
	subGoods := r.Group("/sub_goods")
	home := r.Group("/home")
	home.GET("/banner", func(context *gin.Context) {
		context.Writer.Write([]byte(Banner.GetBannerList().Get()))
	})
	home.PUT("/banner", func(context *gin.Context) {
		type name struct {
			Id         int    `json:"id"`
			Img        string `json:"img"`
			SubGoodsId int    `json:"sub_goods_id"`
		}
		var tmp name
		context.ShouldBindJSON(&tmp)
		context.Writer.Write([]byte(Banner.UpdateBanner(tmp.Id, tmp.Img, tmp.SubGoodsId).Get()))
	})
	home.POST("/banner", func(context *gin.Context) {
		type name struct {
			Img        string `json:"img"`
			SubGoodsId int    `json:"sub_goods_id"`
		}
		var tmp name
		context.ShouldBindJSON(&tmp)
		context.Writer.Write([]byte(Banner.AddBanner(tmp.Img, tmp.SubGoodsId).Get()))
	})
	home.DELETE("/banner", func(context *gin.Context) {
		type name struct {
			Id int `json:"id"`
		}
		var tmp name
		context.ShouldBindJSON(&tmp)
		context.Writer.Write([]byte(Banner.RemoveBanner(tmp.Id).Get()))
	})
	home.OPTIONS("/banner", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	subGoods.GET("", func(context *gin.Context) {
		goodsId := Utils.StrToInt(context.Query("goods_id"))
		context.Writer.Write([]byte(SubGoods.QueryAllSubGoods(goodsId).Get()))
	})
	subGoods.OPTIONS("", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	subGoods.POST("", func(context *gin.Context) {
		type name struct {
			GoodsId  int     `json:"goods_id"`
			Price    float32 `json:"price"`
			Stoke    int     `json:"stoke"`
			Sell     int     `json:"sell"`
			Img      string  `json:"img"`
			Template []int   `json:"template"`
		}
		var tmp name
		context.ShouldBindJSON(&tmp)
		context.Writer.Write([]byte(SubGoods.AddSubGoods(tmp.GoodsId, tmp.Price, tmp.Stoke, tmp.Sell, tmp.Img, tmp.Template).Get()))
	})
	subGoods.PUT("", func(context *gin.Context) {
		type name struct {
			Id       int     `json:"id"`
			Price    float32 `json:"price"`
			Stoke    int     `json:"stoke"`
			Sell     int     `json:"sell"`
			Img      string  `json:"img"`
			Template []int   `json:"template"`
		}
		var tmp name
		context.ShouldBindJSON(&tmp)
		context.Writer.Write([]byte(SubGoods.UpdateSubGoods(tmp.Id, tmp.Price, tmp.Stoke, tmp.Sell, tmp.Img, tmp.Template).Get()))
	})
	goods.GET("", func(context *gin.Context) {
		goodsTitle := context.Query("title")
		limit := Utils.StrToInt(context.Query("limit"))
		offset := Utils.StrToInt(context.Query("offset"))
		context.Writer.Write([]byte(Goods.QueryGoods(goodsTitle, limit, offset).Get()))
	})
	goods.GET("/list", func(context *gin.Context) {
		context.Writer.Write([]byte(Goods.QueryAllGoods().Get()))
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
	goods.PUT("", func(context *gin.Context) {
		type name struct {
			Id        int      `json:"id"`
			Title     string   `json:"title"`
			Desc      string   `json:"desc"`
			Template  string   `json:"template"`
			Banner    []string `json:"banner"`
			DetailImg []string `json:"detail_img"`
		}
		var tmp name
		if context.ShouldBindJSON(&tmp) == nil {
			context.Writer.Write([]byte(Goods.UpdateGoods(tmp.Id, tmp.Title, tmp.Desc, tmp.Template, tmp.Banner, tmp.DetailImg).Get()))
		}

	})
	user.GET("", func(context *gin.Context) {
		nickName := context.Query("nick_name")
		limit := Utils.StrToInt(context.Query("limit"))
		offset := Utils.StrToInt(context.Query("offset"))
		context.Writer.Write([]byte(User.QueryUser(nickName, limit, offset).Get()))
	})
	order.GET("/query", func(context *gin.Context) {
		context.Writer.Write([]byte(Order.SummaryOrder().Get()))
	})
	order.GET("", func(context *gin.Context) {
		limit := Utils.StrToInt(context.Query("limit"))
		offset := Utils.StrToInt(context.Query("offset"))
		context.Writer.Write([]byte(Order.GetOrder(limit, offset).Get()))
	})
	r.Run(":45678")
}
