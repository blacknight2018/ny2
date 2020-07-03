package entity

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
)

type Shop struct {
	//
	db *gorm.DB

	Id       int64  `json:"id";gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	SchoolId int64  `json:"school_id";gorm:"column:school_id;NOT NULL;"`
	Name     string `json:"name";gorm:"column:name;NOT NULL"`
}

func (s *Shop) getDb() *gorm.DB {
	if s.db == nil {
		return db.GetDB()
	}
	return s.db
}
func (s *Shop) setDb(d *gorm.DB) {
	s.db = d
}

func (s *Shop) SelectById() bool {
	return nil == s.getDb().Raw("select * from shop where id = ?", s.Id).Scan(&s).Error
}
