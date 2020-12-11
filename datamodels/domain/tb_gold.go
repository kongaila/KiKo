package domain

import "time"

type TbGold struct {
	Id        int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid      string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
	UserUuid  string    `json:"userUuid" gorm:"type:CHAR(32);"`
	GlodNum   int32     `json:"glodNum" gorm:"type:INT(11);"`
	GoldType  int32     `json:"goldType" gorm:"type:INT(2);"`
}
