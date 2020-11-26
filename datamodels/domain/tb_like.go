package domain

import "time"

// 收藏模型
type TbLike struct {
	Id          int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `gorm:"type:CHAR(32);NOT NULL"`
	UserUuid    string    `gorm:"type:CHAR(32);"`
	ArticleUuid string    `gorm:"type:CHAR(32);"`
	CreateAt    time.Time `gorm:"type:DATETIME(0);"`
}
