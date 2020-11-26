package domain

import "time"

// 热榜模型
type TbTop struct {
	Id          int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Description string    `gorm:"type:VARCHAR(200);"`
	CreateAt    time.Time `gorm:"type:DATETIME(0);"`
	Num         int32     `gorm:"type:INT(11);"`
	ArticleUuid string    `gorm:"type:CHAR(32);"`
}
