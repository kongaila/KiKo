package domain

import "time"

// 热榜模型
type TbTop struct {
	Id          int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Description string    `json:"description" gorm:"type:VARCHAR(200);"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
	Num         int32     `json:"num" gorm:"type:INT(11);"`
	ArticleUuid string    `json:"articleUuid" gorm:"type:CHAR(32);"`
}
