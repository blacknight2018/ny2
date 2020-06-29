package bs

import (
	"github.com/jinzhu/gorm"
	"ny2/bs/entity"
	"ny2/db"
)

type Dorm struct {
	//
	db *gorm.DB

	Id            int64   `json:"dorm_id";gorm:"column:id"`
	SchoolId      int64   `json:"school_id";gorm:"column:id;"`
	DormName      string  `json:"dorm_name";gorm:"column:dorm_name;"`
	DormLongitude float32 `json:"dorm_longitude";gorm:"column:dorm_longitude;"`
	DormLatitude  float32 `json:"dorm_latitude";gorm:"column:dorm_latitude;"`
}

func (u *Dorm) getDb() *gorm.DB {
	if u.db == nil {
		return db.GetDB()
	}
	return u.db
}
func (u *Dorm) setDb(d *gorm.DB) {
	u.db = d
}

func (u *Dorm) SelectOrder(limit int64, offset int64) (bool, []entity.Order) {
	var o []entity.Order
	sql := "select * from `order` where dorm_id = ? order by id desc limit  ? offset  ?"
	r := u.getDb().Raw(sql, u.Id, limit, offset).Scan(&o).Error
	return r == nil, o
}

func (u *Dorm) SelectOrderSize() (bool, int64) {
	sql := "select count(*) as size from `order` where dorm_id = ?"
	type t struct {
		Size int64 `gorm:"column:size"`
	}
	var tmp t
	u.getDb().Raw(sql, u.Id).Scan(&tmp)
	return true, tmp.Size
}

func (u *Dorm) SelectById() bool {
	sql := `select  * from dorm where id = ?`
	r := u.getDb().Raw(sql, u.Id).Scan(&u).Error
	return r == nil
}
