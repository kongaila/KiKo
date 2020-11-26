package domain

import "time"

// 贴吧模型
type TbClub struct {
	Id           int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid         string    `gorm:"type:CHAR(32);NOT NULL"`
	MasterName   string    `gorm:"type:VARCHAR(255);"`
	MasterUuid   string    `gorm:"type:CHAR(32);"`
	Name         string    `gorm:"type:VARCHAR(255);"`
	HeadImg      string    `gorm:"type:VARCHAR(255);"`
	Description  string    `gorm:"type:VARCHAR(500);"`
	Type         string    `gorm:"type:VARCHAR(255);"`
	TypeName     string    `gorm:"type:VARCHAR(255);"`
	AttentionNum int32     `gorm:"type:INT(11);"`
	ArticleNum   int32     `gorm:"type:INT(11);"`
	MemberNum    int32     `gorm:"type:INT(11);"`
	CreateAt     time.Time `gorm:"type:DATETIME(0);"`
	UpdateAt     time.Time `gorm:"type:DATETIME(0);"`
	Creater      string    `gorm:"type:VARCHAR(255);"`
	Updater      string    `gorm:"type:VARCHAR(255);"`
}
