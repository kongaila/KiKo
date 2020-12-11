package domain

import "time"

// 用户贴吧关系模型
type TbUserClub struct {
	Id        int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid      string    `json:"uuid" gorm:"type:CHAR(32);"`
	ClubUuid  string    `json:"clubUuid" gorm:"type:CHAR(32);"`
	UserUuid  string    `json:"userUuid" gorm:"type:CHAR(32);"`
	Identity  int32     `json:"identity" gorm:"type:INT(11);"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
}
