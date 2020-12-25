package DbModel

import (
	"admin/Utils"
	"github.com/jinzhu/gorm"
	"time"
)

type SubGoods struct {
	Id         int        `json:"id" gorm:"column:id;primary_key"`
	Price      float32    `json:"price" gorm:"column:price"`
	Stoke      *int       `json:"stoke" gorm:"column:stoke"`
	Sell       int        `json:"sell" gorm:"column:sell"`
	Img        string     `json:"img" gorm:"column:img"`
	GoodsId    int        `json:"goods_id" gorm:"column:goods_id"`
	Template   string     `json:"template" gorm:"column:template"`
	CreateTime *time.Time `json:"-" gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `json:"-" gorm:"column:update_time" sql:"-"`
}

func (s *SubGoods) TableName() string {
	return "sub_goods"
}

func (s *SubGoods) Update() bool {
	return UpdateDBObj(s)
}

func (s *SubGoods) Insert() bool {
	return InsertDBObj(s)
}
func (s *SubGoods) InsertWithDB(db *gorm.DB) bool {
	return db.Create(s).Error == nil
}
func (s *SubGoods) UpdateWithDB(db *gorm.DB) bool {
	return db.Model(&SubGoods{}).Where("id = ?", s.Id).Update(s).Error == nil
}
func SelectSubGoodsBySubGoodsId(subGoodsId int) (bool, *SubGoods) {
	var subGoods SubGoods
	return SelectTableRecordById((&SubGoods{}).TableName(), subGoodsId, nil, &subGoods), &subGoods
}

func SelectSubGoodsByGoodsId(goodsId int) (bool, []SubGoods) {
	var subGoodsSet []SubGoods
	var condition = make(map[string]interface{})
	condition["goods_id"] = goodsId
	return SelectTableRecordSet((&SubGoods{}).TableName(), &subGoodsSet, condition, nil, nil, Utils.EmptyString), subGoodsSet
}
func SelectSubGoodsSet(condition map[string]interface{}, limit int, offset int) (bool, []SubGoods) {
	var subGoodsSet []SubGoods
	return SelectTableRecordSet((&SubGoods{}).TableName(), &subGoodsSet, condition, &limit, &offset, Utils.EmptyString), subGoodsSet
}

func SelectSubGoodsSetDescCreateTime(condition map[string]interface{}, limit int, offset int) (bool, []SubGoods) {
	var subGoodsSet []SubGoods
	return SelectTableRecordSet((&SubGoods{}).TableName(), &subGoodsSet, condition, &limit, &offset, "create_time desc"), subGoodsSet
}

func SelectSubGoodsSetDescSell(condition map[string]interface{}, limit int, offset int) (bool, []SubGoods) {
	var subGoodsSet []SubGoods
	return SelectTableRecordSet((&SubGoods{}).TableName(), &subGoodsSet, condition, &limit, &offset, "sell desc"), subGoodsSet
}

func SelectSubGoodsByTemplateIndex(goodsId int, templateIndex string) (bool, *SubGoods) {
	var subGoodsSet []SubGoods
	condition := map[string]interface{}{"goods_id": goodsId, "template": templateIndex}
	limit := 1
	offset := 0
	SelectTableRecordSet((&SubGoods{}).TableName(), &subGoodsSet, condition, &limit, &offset, Utils.EmptyString)
	if len(subGoodsSet) == 0 {
		return false, nil
	}
	return true, &subGoodsSet[0]
}
