package entity

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
	"time"
)

const (
	Delivery = iota
	Food     = iota
	Buy      = iota
)

type Order struct {
	//
	db *gorm.DB

	Id         int64      `json:"order_id";gorm:"column:id;unique_index;PRIMARY_KEY;"`
	StuId      int64      `json:"stu_id";gorm:"column:stu_id;NOT NULL;"`
	Price      string     `json:"price";gorm:"column:price;"`
	FinishTime *time.Time `json:"finish_time";gorm:"column:finish_time;"`
	Type       string     `json:"type";gorm:"column:type;type:enum('2','1','0');NOT NULL"`
	Comment    string     `json:"comment";gorm:"column:comment;"`
	RecvStu    *int64     `json:"recv_stu,omitempty";gorm:"column:recv_stu;"`
	SchoolId   int64      `json:"school_id";gorm:"column:school_id;NOT NULL;"`
	DormId     int64      `json:"dorm_id";gorm:"column:dorm_id;NOT NULL;"`
	AvatarUrl  string     `json:"avatar_url";gorm:"column:avatar_url;NOT NULL;"`
	TemplateId string     `json:"template_id";gorm:"column:template_id;NOT NULL;"`
}

func (u *Order) getDb() *gorm.DB {
	if u.db == nil {
		return db.GetDB()
	}
	return u.db
}
func (u *Order) setDb(d *gorm.DB) {
	u.db = d
}
