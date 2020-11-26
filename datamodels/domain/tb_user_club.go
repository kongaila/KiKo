package domain

import "time"

// 用户贴吧关系模型
type TbUserClub struct {
	Id       int32  `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid     string `gorm:"type:CHAR(32);"`
	ClubUuid string `gorm:"type:CHAR(32);"`

	UserUuid string    `gorm:"type:CHAR(32);"`
	Identity int32     `gorm:"type:INT(11);"`
	CreateAt time.Time `gorm:"type:DATETIME(0);"`
}
