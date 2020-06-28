package entity

type Detail struct {
	Id     int64 `gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	MsgId  int64 `gorm:"column:msg_id;NUL NULL;"`
	StuId  int64 `gorm:"column:stu_id;NOT NULL;"`
	IsRead bool  `json:"is_read,omitempty";gorm:"column:is_read;NOT NULL;"`
}
