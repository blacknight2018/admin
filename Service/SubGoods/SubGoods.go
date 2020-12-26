package SubGoods

import (
	"admin/DbModel"
	"admin/Result"
	"admin/Utils"
)

func QueryAllSubGoods(goodsId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectSubGoodsByGoodsId(goodsId); ok {
		ret.Code = Result.Ok
		ret.Data = data
	}
	return ret
}

func UpdateSubGoods(subGoodsId int, price float32, stoke int, sell int, img string, template []int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectSubGoodsBySubGoodsId(subGoodsId); ok {
		data.Price = price
		data.Stoke = &stoke
		data.Sell = &sell
		data.Template = Utils.IntArrayToJSON(template)
		data.Img = img
		if data.Update() {
			ret.Code = Result.Ok
		}
	}
	return ret
}
