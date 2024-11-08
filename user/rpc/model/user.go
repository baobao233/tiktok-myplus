package model

import "database/sql"

type User struct {
	Id            uint64       `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name          string       `gorm:"column:name;type:varchar(255);comment:用户姓名;NOT NULL" json:"name"`
	Gender        uint         `gorm:"column:gender;type:tinyint(3) unsigned;default:0;comment:用户性别;NOT NULL" json:"gender"`
	Mobile        string       `gorm:"column:mobile;type:varchar(255);comment:用户电话;NOT NULL" json:"mobile"`
	Password      string       `gorm:"column:password;type:varchar(255);comment:用户密码;NOT NULL" json:"password"`
	Dec           string       `gorm:"column:dec;type:varchar(255);comment:个性签名;NOT NULL" json:"dec"`
	Avatar        string       `gorm:"column:avatar;type:varchar(255);comment:头像;NOT NULL" json:"avatar"`
	BackgroundUrl string       `gorm:"column:background_url;type:varchar(255);comment:背景图;NOT NULL" json:"background_url"`
	CreateTime    sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime    sql.NullTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"update_time"`
}
