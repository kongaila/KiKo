package domain

import "time"

// 文章模型
type TbArticle struct {
	Id          int32     `json:"id" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	ClubUuid    string    `json:"clubUuid" gorm:"type:CHAR(32);"`
	UserUuid    string    `json:"userUuid" gorm:"type:CHAR(32);"`
	Title       string    `json:"title" gorm:"type:VARCHAR(255);"`
	Content     string    `json:"content" gorm:"type:TEXT;"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME(0);"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME(0);"`
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
	Num         int32     `json:"num" gorm:"-"`               // 回复数量
	CraterNick  string    `json:"craterNick" gorm:"-"`        // 创建人昵称
	Status      int32     `json:"status" gorm:"type:INT(2);"` // 审核状态  是否被举报 0 没有 1被举报（审核中） 2 举报成功 3举报驳回
	ReportMsg   string    `json:"state" gorm:"type:TEXT;"`    // 举报信息
}
