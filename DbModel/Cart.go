package DbModel

import (
	"admin/Config"
	"admin/Utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Cart struct {
	Id         int        `json:"id" gorm:"column:id;primary_key"`
	UserId     int        `json:"user_id" gorm:"column:user_id"`
	SubGoodsId int        `json:"sub_goods_id" gorm:"column:sub_goods_id"`
	Amount     int        `json:"amount" gorm:"column:amount"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time" sql:"-"`
}

func (c *Cart) TableName() string {
	return "cart"
}

func (c *Cart) Update() bool {
	return UpdateDBObj(c)
}

func (c *Cart) Insert() bool {
	return InsertDBObj(c)
}

func (c *Cart) Delete() bool {
	return DeleteDBObj(c)
}

func (c *Cart) DeleteWithDB(db *gorm.DB) bool {
	return db.Delete(c).Error == nil
}

/**
 * @Description: 删除userId下的cart记录
 * @param cartId
 * @param userId
 * @return bool
 */
func DeleteCartByUserId(cartId int, userId int) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	defer db.Close()
	return nil == db.Where("user_id = ?", userId).Where("id = ?", cartId).Delete(Cart{}).Error
}

/**
 * @Description:根据cartId获取一条cart记录
 * @param cartId
 * @return bool
 * @return *Cart
 */
func SelectCartByCartId(cartId int) (bool, *Cart) {
	var cart Cart
	return SelectTableRecordById((&Cart{}).TableName(), cartId, nil, &cart), &cart
}

/**
 * @Description: 多条件查询cart集合
 * @param condition
 * @param limit
 * @param offset
 * @return bool
 * @return []Cart
 */
func SelectCartSet(condition map[string]interface{}, limit int, offset int) (bool, []Cart) {
	var cartSet []Cart
	return SelectTableRecordSet((&Cart{}).TableName(), &cartSet, condition, &limit, &offset, Utils.EmptyString), cartSet
}

/**
 * @Description: 获取userId下的所有cart记录
 * @param userId
 * @param limit
 * @param offset
 * @return bool
 * @return []Cart
 */
func SelectCartSetByUserId(userId int, limit int, offset int) (bool, []Cart) {
	var cartSet []Cart
	return SelectTableRecordSet((&Cart{}).TableName(), &cartSet, map[string]interface{}{"user_id": userId}, &limit, &offset, "create_time desc"), cartSet
}

/**
 * @Description:
 * @param userId
 * @param subGoodsId
 * @return bool
 * @return *Cart
 */
func SelectUserCartBySubGoodsId(userId int, subGoodsId int) (bool, *Cart) {
	var cartSet []Cart
	condition := map[string]interface{}{"user_id": userId, "sub_goods_id": subGoodsId}
	limit := 1
	offset := 0
	SelectTableRecordSet((&Cart{}).TableName(), &cartSet, condition, &limit, &offset, Utils.EmptyString)
	if len(cartSet) >= 1 {
		return true, &cartSet[0]
	}
	return false, nil
}
