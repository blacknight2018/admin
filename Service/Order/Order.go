package Order

import (
	"admin/DbModel"
	"admin/Result"
	"admin/Utils"
)

const (
	UnPay      = iota
	Pay        = iota
	UnDelivery = iota
	Delivery   = iota
	All        = iota
)

/**
 * @Description: 分页查询订单
 * @param limit
 * @param offset
 * @return Result.Result
 */
func GetOrder(limit int, offset int) Result.Result {
	var ret Result.Result
	ret.Code = Result.Ok

	type name struct {
		Total int             `json:"total"`
		Order []DbModel.Order `json:"order"`
	}
	var retData name
	if ok, data := DbModel.SelectOrderSet(nil, limit, offset, Utils.EmptyString); ok {
		_, cnt := DbModel.SelectOrderCount()
		retData.Total = cnt
		retData.Order = data
		ret.Data = retData
		ret.Code = Result.Ok
	}
	return ret
}

/**
 * @Description: 统计订单信息
 * @return Result.Result
 */
func SummaryOrder() Result.Result {
	var ret Result.Result
	ret.Code = Result.Ok
	type name struct {
		Total           int `json:"total"`
		PayCount        int `json:"pay_count"`
		UnPayCount      int `json:"unPay_count"`
		DeliveryCount   int `json:"delivery_count"`
		UnDeliveryCount int `json:"unDelivery_count"`
	}
	var retData name

	if ok, data := DbModel.SelectOrderSetByStatus(UnPay, nil, nil, Utils.EmptyString); ok {
		retData.UnPayCount = len(data)
		retData.Total += retData.UnPayCount
	}
	if ok, data := DbModel.SelectOrderSetByStatus(Pay, nil, nil, Utils.EmptyString); ok {
		retData.PayCount = len(data)
		retData.Total += retData.PayCount
	}
	if ok, data := DbModel.SelectOrderSetByStatus(Delivery, nil, nil, Utils.EmptyString); ok {
		retData.DeliveryCount = len(data)
		retData.Total += retData.DeliveryCount
	}
	if ok, data := DbModel.SelectOrderSetByStatus(UnDelivery, nil, nil, Utils.EmptyString); ok {
		retData.UnDeliveryCount = len(data)
		retData.Total += retData.UnDeliveryCount
	}

	ret.Data = retData
	return ret
}
