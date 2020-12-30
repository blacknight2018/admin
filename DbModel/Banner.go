package DbModel

import (
	"admin/Utils"
)

type Banner struct {
	Id         int         `json:"id" gorm:"column:id"`
	Img        string      `json:"img" gorm:"column:img"`
	SubGoodsId int         `json:"sub_goods_id" gorm:"column:sub_goods_id"`
	CreateTime *Utils.Time `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime *Utils.Time `json:"update_time" gorm:"column:update_time" sql:"-"`
}

func (c *Banner) TableName() string {
	return "banner"
}

func (c *Banner) Update() bool {
	return UpdateDBObj(c)
}

func (c *Banner) Insert() bool {
	return InsertDBObj(c)
}

func (c *Banner) Delete() bool {
	return DeleteDBObj(c)
}
func SelectBannerSet(condition map[string]interface{}, limit *int, offset *int) (bool, []Banner) {
	var bannerSet []Banner
	return SelectTableRecordSet((&Banner{}).TableName(), &bannerSet, condition, limit, offset, Utils.EmptyString), bannerSet
}

func SelectBannerByBannerId(bannerId int) (bool, *Banner) {
	var banner Banner
	return SelectTableRecordById((&Banner{}).TableName(), bannerId, nil, &banner), &banner
}
