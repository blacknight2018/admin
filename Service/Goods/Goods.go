package Goods

import (
	"admin/DbModel"
	"admin/Result"
)

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
