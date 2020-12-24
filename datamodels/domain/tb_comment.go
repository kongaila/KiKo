package domain

import "time"

// 评论模型
type TbComment struct {
	Id           int32       `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid         string      `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	PId          int32       `json:"pId" gorm:"type:INT(11);"`
	PUuid        string      `json:"pUuid" gorm:"type:CHAR(32);"`
	Type         int32       `json:"type" gorm:"type:INT(2);"`
	UserUuid     string      `json:"userUuid" gorm:"type:CHAR(32);"`
	ArticleUuid  string      `json:"articleUuid" gorm:"type:CHAR(32);"`
	Content      string      `json:"content" gorm:"type:TEXT;"`
	CreatedAt    time.Time   `json:"CreatedAt" gorm:"type:DATETIME(0);"`
	Status       int32       `json:"status" gorm:"type:INT(2);"`
	Like         int32       `json:"like" gorm:"type:INT(11);"`
	ArticleTitle string      `json:"articleTitle" gorm:"type:VARCHAR(100);"`
	User         TbUser      `json:"user" gorm:"embedded;"`
	SonComment   []TbComment `json:"sonComment"`
	ReportMsg    string      `json:"state" gorm:"type:TEXT;"` // 举报信息
}
