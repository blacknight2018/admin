package Banner

import (
	"admin/DbModel"
	"admin/Result"
)

func GetBannerList() Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectBannerSet(nil, nil, nil); ok {
		ret.Code = Result.Ok
		ret.Data = data
	}
	return ret
}

func UpdateBanner(bannerId int, img string, subGoodsId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectBannerByBannerId(bannerId); ok {
		ret.Code = Result.Ok
		ret.Data = data
	}
	return ret
}

func AddBanner(img string, subGoodsId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	var banner DbModel.Banner
	banner.Img = img
	banner.SubGoodsId = subGoodsId
	if banner.Insert() {
		ret.Code = Result.Ok
	}
	return ret
}

func RemoveBanner(bannerId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectBannerByBannerId(bannerId); ok {
		if data.Delete() {
			ret.Code = Result.Ok
		}
	}
	return ret
}
