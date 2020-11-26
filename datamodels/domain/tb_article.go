package domain

import "time"

// 文章模型
type TbArticle struct {
	Id          int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `gorm:"type:CHAR(32);NOT NULL"`
	ClubUuid    string    `gorm:"type:CHAR(32);"`
	UserUuid    string    `gorm:"type:CHAR(32);"`
	Title       string    `gorm:"type:VARCHAR(255);"`
	Content     string    `gorm:"type:TEXT;"`
	CreateAt    time.Time `gorm:"type:DATETIME(0);"`
	UpdateAt    time.Time `gorm:"type:DATETIME(0);"`
	Creater     string    `gorm:"type:VARCHAR(255);"`
	Updater     string    `gorm:"type:VARCHAR(255);"`
	OpenNum     int32     `gorm:"type:INT(11);"`
	Description string    `gorm:"type:VARCHAR(300);"`
	IsGood      int32     `gorm:"type:INT(2);"`
	Flag        int32     `gorm:"type:INT(2);"`
	IsMoney     int32     `gorm:"type:INT(2);"`
	Money       int32     `gorm:"type:INT(11);"`
}
