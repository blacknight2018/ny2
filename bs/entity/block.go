package entity

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
)

type Block struct {
	//
	db *gorm.DB

	Id     int64 `json:"id";gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	StuId  int64 `gorm:"column:stu_id;"`
	DstStu int64 `gorm:"column:dst_stu;"`
}

func (m *Block) getDb() *gorm.DB {
	if m.db == nil {
		return db.GetDB()
	}
	return m.db
}
func (m *Block) setDb(d *gorm.DB) {
	m.db = d
}

func (m *Block) SelectById() bool {
	return nil == m.getDb().Raw("select * from block where id = ?", m.Id).Error
}

func (m *Block) Insert() bool {
	return nil == m.getDb().Exec("insert into block(stu_id,dst_stu) values(?,?)", m.StuId, m.DstStu).Error
}
