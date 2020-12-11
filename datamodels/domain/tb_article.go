package domain

import "time"

// 文章模型
type TbArticle struct {
	Id          int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	ClubUuid    string    `json:"clubUuid" gorm:"type:CHAR(32);"`
	UserUuid    string    `json:"userUuid" gorm:"type:CHAR(32);"`
	Title       string    `json:"title" gorm:"type:VARCHAR(255);"`
	Content     string    `json:"content" gorm:"type:TEXT;"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
	UpdatedAt   time.Time `json:"UpdatedAt" gorm:"type:DATETIME(0);"`
	Creater     string    `json:"-" gorm:"type:VARCHAR(255);"`
	Updater     string    `json:"-" gorm:"type:VARCHAR(255);"`
	OpenNum     int32     `json:"openNum" gorm:"type:INT(11);"`
	Description string    `json:"description" gorm:"type:VARCHAR(300);"`
	IsGood      int32     `json:"isGood" gorm:"type:INT(2);"`
	Flag        int32     `json:"flag" gorm:"type:INT(2);"`
	IsMoney     int32     `json:"isMoney" gorm:"type:INT(2);"`
	Money       int32     `json:"money" gorm:"type:INT(11);"`
	Type        string    `json:"type" gorm:"type:VARCHAR(255);"`
	TypeName    string    `json:"typeName" gorm:"type:VARCHAR(255);"`
	Num         int32     `json:"num" gorm:"type:INT(11);"`
}
