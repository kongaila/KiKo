package domain

import "time"

type TbGold struct {
	Id       int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid     string    `gorm:"type:CHAR(32);NOT NULL"`
	CreateAt time.Time `gorm:"type:DATETIME(0);"`
	UserUuid string    `gorm:"type:CHAR(32);"`
	GlodNum  int32     `gorm:"type:INT(11);"`
	GoldType int32     `gorm:"type:INT(2);"`
}
