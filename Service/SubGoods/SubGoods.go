package SubGoods

import (
	"admin/DbModel"
	"admin/Result"
	"admin/Utils"
)

/**
 * @Description: 添加一种商品规格
 * @param goodsId
 * @param price
 * @param stoke
 * @param sell
 * @param img
 * @param template
 * @return Result.Result
 */
func AddSubGoods(goodsId int, price float32, stoke int, sell int, img string, template []int) Result.Result {
	var subGoods DbModel.SubGoods
	var ret Result.Result
	ret.Code = Result.UnKnow

	subGoods.GoodsId = goodsId
	subGoods.Price = price
	subGoods.Stoke = &stoke
	subGoods.Sell = &sell
	subGoods.Img = img
	subGoods.Template = Utils.IntArrayToJSON(template)
	if subGoods.Insert() {
		ret.Code = Result.Ok
	}

	return ret
}

/**
 * @Description: 列出某个商品下的所有商品规格
 * @param goodsId
 * @return Result.Result
 */
func QueryAllSubGoods(goodsId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectSubGoodsByGoodsId(goodsId); ok {
		ret.Code = Result.Ok
		ret.Data = data
	}
	return ret
}

/**
 * @Description: 更新某个子商品
 * @param subGoodsId
 * @param price
 * @param stoke
 * @param sell
 * @param img
 * @param template
 * @return Result.Result
 */
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
