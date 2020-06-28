package bs

import (
	"github.com/jinzhu/gorm"
	"ny2/bs/entity"
	"ny2/db"
)

//type User struct {
//	Id         int64      `json:"user_id";gorm:"column:id;"`
//	OpenId     string     `gorm:"column:open_id;"`
//	NickName   string     `gorm:"column:nick_name;"`
//	CreateTime *time.Time `gorm:"column:create_time;"`
//	Mobile     string     `gorm:"column:mobile;"`
//	AvatarUrl  string     `gorm:"column:avatar_url;"`
//}

type Stu struct {
	//
	db *gorm.DB

	Us entity.User

	StuId     *int64 `gorm:"column:id;"`
	DormId    *int64 `gorm:"column:dorm_id;"`
	StuNumber string `gorm:"column:stu_number;"`
	DormRoom  string `gorm:"column:dorm_room;"`
	UserId    *int64 `gorm:"column:user_id;"`

	Dm Dorm
}

func (u *Stu) getDb() *gorm.DB {
	if u.db == nil {
		return db.GetDB()
	}
	return u.db
}
func (u *Stu) setDb(d *gorm.DB) {
	u.db = d
}
func (u *Stu) SelectByOpenId() bool {
	sql := `select id ,open_id,nick_name,mobile,avatar_url from user where open_id = ?;`
	r := u.getDb().Raw(sql, u.Us.OpenId).Scan(&u.Us).Error
	if r != nil {
		return false
	}
	return u.SelectById()
}

func (u *Stu) SelectById() bool {
	sql := `SELECT
user.id,
user.open_id,
user.nick_name,
user.create_time,
user.mobile,
user.avatar_url

FROM
user 

where  user.id = ?;
`
	err := u.getDb().Raw(sql, u.Us.Id).Scan(&u.Us).Error
	if err != nil {
		return false
	}
	sql = `select id ,dorm_id,stu_number,dorm_room,user_id from stu where stu.user_id = ?`
	err = u.getDb().Raw(sql, u.Us.Id).Scan(&u).Error
	if err != nil {
		return true
	}

	if u.DormId != nil {
		sql = `SELECT * from dorm where id = ?`
		err := u.getDb().Raw(sql, u.DormId).Scan(&u.Dm).Error
		return nil == err
	}
	return true
}

func (u *Stu) SelectByStuId() bool {
	sql := `select * from stu where id = ?`
	err := u.getDb().Raw(sql, u.StuId).Scan(&u).Error
	if err != nil {
		return false
	}
	u.Us.Id = *u.UserId
	ok := u.SelectById()
	if false == ok {
		return false
	}
	return true

}

func (u *Stu) UpdateById() bool {
	sql := `UPDATE USER
SET open_id = ?,
 nick_name = ?,
 mobile = ?,
 avatar_url = ?
WHERE
	id = ?;`
	err := u.getDb().Exec(sql, u.Us.OpenId, u.Us.NickName, u.Us.Mobile, u.Us.AvatarUrl, u.Us.Id).Error
	if err != nil {
		return false
	}

	if u.StuId != nil {
		sql = `UPDATE stu
SET dorm_id = ?, stu_number = ?,  dorm_room = ?
WHERE
	id = ?`
		return nil == u.getDb().Exec(sql, u.DormId, u.StuNumber, u.DormRoom, u.StuId).Error
	}
	sql = `insert into stu(dorm_id,stu_number,user_id,dorm_room) values(?,?,?,?);`
	err = u.getDb().Exec(sql, u.DormId, u.StuNumber, u.Us.Id, u.Dm.DormName).Error
	return err == nil
}

func (u *Stu) Insert() bool {

	//使用新的线程连接，因为下面要insert后获取插入后的id,如果用同一个链接多线程下结果可能错误
	newCon := db.NewDbCon()
	tx := newCon.Begin()
	u.setDb(tx)
	defer func() {
		newCon.Close()
		u.setDb(nil)
	}()
	sql := `insert into user(open_id,nick_name,mobile,avatar_url) values(?,?,?,?);`
	err := tx.Exec(sql, u.Us.OpenId, u.Us.NickName, u.Us.Mobile, u.Us.AvatarUrl).Error
	if err != nil {
		r := tx.Rollback().Error
		return r == nil
	}
	sql = `SELECT LAST_INSERT_ID() as id;`
	err = tx.Raw(sql).Scan(&u.Us).Error
	if err != nil {
		r := tx.Rollback().Error
		return r == nil
	}
	if false == u.SelectById() {
		r := tx.Rollback().Error
		return r == nil
	}
	if false == u.UpdateById() {
		tx.Rollback()
		return false
	}
	return tx.Commit().Error == nil
}

func (u *Stu) IsOpenIdExist() bool {
	sql := `select * from user where open_id = ?;`
	var t Stu
	r := u.getDb().Raw(sql, u.Us.OpenId).Scan(&t)
	return !r.RecordNotFound()
}

func (u *Stu) InsertMsg(m *entity.Msg) bool {
	newCon := db.NewDbCon()
	defer func() {
		newCon.Close()
		u.setDb(nil)
	}()

	sql := `insert into msg(sender_stu,recipient_stu,content,type) values(?,?,?,?);`
	err := u.getDb().Exec(sql, m.SenderStu, m.RecipientStu, m.Content, m.Type).Error
	if err != nil {
		return false
	}
	sql = `SELECT LAST_INSERT_ID() as id;`
	u.getDb().Raw(sql).Scan(&m)
	return true
}

func (u *Stu) InsertOrder(or *entity.Order) bool {
	sql := "insert into `order`(stu_id,price,finish_time,type,comment,recv_stu,school_id,dorm_id,avatar_url,template_id) values(?,?,?,?,?,?,?,?,?,?)"
	err := u.getDb().Exec(sql, u.StuId, or.Price, or.FinishTime, or.Type, or.Comment, or.RecvStu, or.SchoolId, u.DormId, or.AvatarUrl, or.TemplateId).Error
	return err == nil
}

func (u *Stu) InsertMsgDetail(md *entity.Detail) bool {
	newCon := db.NewDbCon()
	defer func() {
		newCon.Close()
		u.setDb(nil)
	}()
	u.setDb(newCon)
	sql := `insert into msgdetail(msg_id,stu_id,is_read) values(?,?,?)`
	err := u.getDb().Exec(sql, md.MsgId, md.StuId, md.IsRead).Error

	return err == nil
}

func (u *Stu) QueryStuMsg(stuId int64, limit int) (bool, []entity.Msg) {
	var m []entity.Msg
	sql := `SELECT
	*
FROM
	msg
WHERE
	(
		(sender_stu = ? && recipient_stu = ?) || (sender_stu = ? && recipient_stu = ?)
	) && (
		id NOT IN (
			SELECT
				msg_id
			FROM
				msgdetail
			WHERE
				stu_id = ? && is_read = 1
		)
	) limit ?`
	return nil == u.getDb().Raw(sql, u.StuId, stuId, stuId, u.StuId, u.StuId, limit).Scan(&m).Error, m
}
