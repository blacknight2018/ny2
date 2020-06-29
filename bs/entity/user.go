package entity

import "time"

type User struct {
	Id         int64      `json:"user_id";gorm:"column:id;"`
	OpenId     string     `gorm:"column:open_id;"`
	NickName   string     `gorm:"column:nick_name;"`
	CreateTime *time.Time `gorm:"column:create_time;"`
	Mobile     string     `gorm:"column:mobile;"`
	AvatarUrl  string     `gorm:"column:avatar_url;"`
}
