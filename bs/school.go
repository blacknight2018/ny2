package bs

import (
	"github.com/jinzhu/gorm"
	"ny2/db"
)

type School struct {
	//
	db         *gorm.DB
	Id         int64  `json:"school_id";gorm:"column:id;PRIMARY_KEY;"`
	SchoolName string `json:"school_name";gorm:"column:school_name;"`
}

func (u *School) getDb() *gorm.DB {
	if u.db == nil {
		return db.GetDB()
	}
	return u.db
}
func (u *School) setDb(d *gorm.DB) {
	u.db = d
}
func SelectAllSchool() (bool, []School) {
	var s []School
	sql := `select * from school;`
	err := db.GetDB().Raw(sql).Scan(&s).Error
	return nil == err, s
}

func (s *School) SelectSchoolAllDorm() (bool, []Dorm) {
	var r []Dorm
	sql := `select id ,school_id,dorm_name,dorm_longitude,dorm_latitude from dorm where school_id = ?`
	err := s.getDb().Raw(sql, s.Id).Scan(&r).Error
	return err == nil, r
}
