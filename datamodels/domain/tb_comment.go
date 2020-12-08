package domain

import "time"

// 评论模型
type TbComment struct {
	Id          int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	PId         int32     `json:"pId" gorm:"type:INT(11);"`
	PUuid       string    `json:"pUuid" gorm:"type:CHAR(32);"`
	Type        int32     `json:"type" gorm:"type:INT(2);"`
	UserUuid    string    `json:"userUuid" gorm:"type:CHAR(32);"`
	ArticleUuid string    `json:"articleUuid" gorm:"type:CHAR(32);"`
	Content     string    `json:"content" gorm:"type:TEXT;"`
	CreateAt    time.Time `json:"createAt" gorm:"type:DATETIME(0);"`
	State       int32     `json:"state" gorm:"type:INT(2);"`
}
