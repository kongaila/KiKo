package domain

import "time"

// 贴吧模型
type TbClub struct {
	Id           int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid         string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	MasterName   string    `json:"masterName" gorm:"type:VARCHAR(255);"`
	MasterUuid   string    `json:"masterUuid" gorm:"type:CHAR(32);"`
	Name         string    `json:"name" gorm:"type:VARCHAR(255);"`
	HeadImg      string    `json:"headImg" gorm:"type:VARCHAR(255);"`
	Description  string    `json:"description" gorm:"type:VARCHAR(500);"`
	Type         string    `json:"type" gorm:"type:VARCHAR(255);"`
	TypeName     string    `json:"typeName" gorm:"type:VARCHAR(255);"`
	AttentionNum int32     `json:"attentionNum" gorm:"type:INT(11);"`
	ArticleNum   int32     `json:"articleNum" gorm:"type:INT(11);"`
	MemberNum    int32     `json:"memberNum" gorm:"type:INT(11);"`
	CreatedAt    time.Time `json:"createdAt" gorm:"type:DATETIME(0);"`
	UpdatedAt    time.Time `json:"-" gorm:"type:DATETIME(0);"`
	Creater      string    `json:"creater" gorm:"type:VARCHAR(255);"`
	Updater      string    `json:"-" gorm:"type:VARCHAR(255);"`
	Status       int32     `json:"status" gorm:"type:type:INT(2);"`
}
