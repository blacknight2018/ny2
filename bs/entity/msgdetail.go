package entity

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
)

type Detail struct {
	//
	db     *gorm.DB
	Id     int64 `gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	MsgId  int64 `gorm:"column:msg_id;NUL NULL;"`
	StuId  int64 `gorm:"column:stu_id;NOT NULL;"`
	IsRead bool  `json:"is_read,omitempty";gorm:"column:is_read;NOT NULL;"`
}

func (md *Detail) getDb() *gorm.DB {
	if md.db == nil {
		return db.GetDB()
	}
	return md.db
}
func (md *Detail) setDb(d *gorm.DB) {
	md.db = d
}
func (md *Detail) Update() bool {
	err := nil == md.getDb().Exec("update msgdetail set msg_id = ?,stu_id = ?,is_read = ? where id = ?", md.MsgId, md.StuId, md.IsRead, md.Id).Error
	return err
}
func (md *Detail) SelectById() bool {
	return nil == md.getDb().Raw("select * from msgdetail where id = ?", md.Id).Scan(&md).Error
}
