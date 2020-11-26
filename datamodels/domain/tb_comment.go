package domain

import "time"

// 评论模型
type TbComment struct {
	Id          int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `gorm:"type:CHAR(32);NOT NULL"`
	PId         int32     `gorm:"type:INT(11);"`
	PUuid       string    `gorm:"type:CHAR(32);"`
	Type        int32     `gorm:"type:INT(2);"`
	UserUuid    string    `gorm:"type:CHAR(32);"`
	ArticleUuid string    `gorm:"type:CHAR(32);"`
	Content     string    `gorm:"type:TEXT;"`
	CreateAt    time.Time `gorm:"type:DATETIME(0);"`
	State       int32     `gorm:"type:INT(2);"`
}
