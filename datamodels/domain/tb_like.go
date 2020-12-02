package domain

import "time"

// 收藏模型
type TbLike struct {
	Id          int32     `json:"id" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	UserUuid    string    `json:"userUuid" gorm:"type:CHAR(32);"`
	ArticleUuid string    `json:"articleUuid" gorm:"type:CHAR(32);"`
	CreateAt    time.Time `json:"createAt" gorm:"type:DATETIME(0);"`
}
