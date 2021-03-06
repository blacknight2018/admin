package Goods

import (
	"admin/DbModel"
	"admin/Result"
	"encoding/json"
)

/**
 * @Description: 分页查询商品表
 * @param goodsTitle
 * @param limit
 * @param offset
 * @return Result.Result
 */
func QueryGoods(goodsTitle string, limit int, offset int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	type name struct {
		Total int             `json:"total"`
		Goods []DbModel.Goods `json:"goods"`
	}
	var retData name
	if ok, data := DbModel.SelectGoodsSetByNickName(goodsTitle, limit, offset); ok {
		retData.Goods = data
		retData.Total = DbModel.SelectGoodsSetCountByNickName(goodsTitle, nil, nil)
		ret.Code = Result.Ok
		ret.Data = retData
	}
	return ret
}

/**
 * @Description:更新商品信息
 * @param goodsId
 * @param title
 * @param desc
 * @param template
 * @param banner
 * @param detailImg
 * @return Result.Result
 */
func UpdateGoods(goodsId int, title string, desc string, template string, banner []string, detailImg []string) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	var bannerJson string
	var detailImgJson string
	if jsonBytes, err := json.Marshal(banner); err == nil {
		bannerJson = string(jsonBytes)
	}
	if jsonBytes, err := json.Marshal(detailImg); err == nil {
		detailImgJson = string(jsonBytes)
	}
	if ok, goods := DbModel.SelectGoodsByGoodsId(goodsId); ok {
		goods.DetailImg = detailImgJson
		goods.Banner = bannerJson
		goods.Title = title
		goods.Desc = desc
		goods.Template = template
		if goods.Update() {
			ret.Code = Result.Ok
		}
	}
	return ret
}

/**
 * @Description: 添加商品
 * @param title
 * @param desc
 * @param template
 * @param banner
 * @param detailImg
 * @return Result.Result
 */
func AddGoods(title string, desc string, template string, banner []string, detailImg []string) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	var bannerJson string
	var detailImgJson string
	if jsonBytes, err := json.Marshal(banner); err == nil {
		bannerJson = string(jsonBytes)
	}
	if jsonBytes, err := json.Marshal(detailImg); err == nil {
		detailImgJson = string(jsonBytes)
	}
	var goods DbModel.Goods
	goods.Title = title
	goods.Desc = desc
	goods.Template = template
	goods.Banner = bannerJson
	goods.DetailImg = detailImgJson
	if goods.Insert() {
		ret.Code = Result.Ok
	}
	return ret
}

func QueryAllGoods() Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectGoodsSet(nil, nil, nil); ok {
		ret.Data = data
		ret.Code = Result.Ok
	}
	return ret
}
