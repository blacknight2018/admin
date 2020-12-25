package DbModel

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

type Order struct {
	Id           int     `json:"id" gorm:"column:id;primary_key"`
	UserId       int     `json:"user_id" gorm:"column:user_id"`
	NickName     string  `json:"nick_name" gorm:"column:nick_name"`
	Sex          string  `json:"sex" gorm:"column:sex;type:char(1)"`
	Phone        string  `json:"phone" gorm:"column:phone"`
	Detail       string  `json:"detail" gorm:"column:detail"`
	SubGoods     string  `json:"sub_goods" gorm:"column:sub_goods"`
	Status       int     `json:"status" gorm:"column:status"`
	TotalPrice   float32 `json:"total_price" gorm:"column:total_price"`
	CreateTime   Time    `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime   Time    `json:"update_time" gorm:"column:update_time" sql:"-"`
	DeliveryCode string  `json:"delivery_code" gorm:"column:delivery_code"`
}

func (o *Order) TableName() string {
	return "order"
}

func (o *Order) Update() bool {
	return UpdateDBObj(o)
}

func (o *Order) Insert() bool {
	return InsertDBObj(o)
}

func SelectOrderByOrderId(orderId int) (bool, *Order) {
	var order Order
	return SelectTableRecordById((&Order{}).TableName(), orderId, nil, &order), &order
}

func SelectOrderSet(condition map[string]interface{}, limit int, offset int, order string) (bool, []Order) {
	var orderSet []Order
	return SelectTableRecordSet((&Order{}).TableName(), &orderSet, condition, &limit, &offset, order), orderSet
}

func SelectOrderSetByStatus(status int, limit *int, offset *int, order string) (bool, []Order) {
	var orderSet []Order
	var condition = make(map[string]interface{})
	condition["status"] = status
	return SelectTableRecordSet((&Order{}).TableName(), &orderSet, condition, limit, offset, order), orderSet

}

func (o *Order) InsertOrderWithDB(db *gorm.DB) bool {
	return db.Create(o).Error == nil
}
