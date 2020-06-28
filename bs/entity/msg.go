package entity

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
	"time"
)

type Msg struct {
	//
	db *gorm.DB

	Id           int64      `json:"id";gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	SenderStu    int64      `json:"sender_stu";gorm:"column:sender_stu;"`
	RecipientStu int64      `json:"recipient_stu";gorm:"column:recipient_stu;"`
	Content      string     `json:"content";gorm:"column:content;NOT NULL;"`
	Type         string     `json:"type";gorm:"column:type;type:enum('TXT');NOT NULL;"`
	CreateTime   *time.Time `json:"create_time";gorm:"column:create_time;"`
}

func (m *Msg) getDb() *gorm.DB {
	if m.db == nil {
		return db.GetDB()
	}
	return m.db
}
func (m *Msg) setDb(d *gorm.DB) {
	m.db = d
}
