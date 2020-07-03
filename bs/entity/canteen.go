package entity

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
)

type Canteen struct {
	//
	db *gorm.DB

	Id       int64  `json:"id";gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	SchoolId int64  `json:"school_id";gorm:"column:school_id;NOT NULL;"`
	Name     string `json:"name";gorm:"column:name;NOT NULL"`
}

func (m *Canteen) getDb() *gorm.DB {
	if m.db == nil {
		return db.GetDB()
	}
	return m.db
}
func (m *Canteen) setDb(d *gorm.DB) {
	m.db = d
}

func (m *Canteen) SelectById() bool {
	return nil == m.getDb().Raw("select * from canteen where id = ?", m.Id).Scan(&m).Error
}
