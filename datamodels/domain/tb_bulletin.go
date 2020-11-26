package domain

import "time"

// 公告栏模型
type TbBulletin struct {
	Id         int32     `gorm:"type:INT(11);NOT NULL"`
	Title      string    `gorm:"type:VARCHAR(50);"`
	Content    string    `gorm:"type:TEXT;"`
	CreateAt   time.Time `gorm:"type:DATETIME(0);"`
	CreateUuid string    `gorm:"type:CHAR(32);"`
}
