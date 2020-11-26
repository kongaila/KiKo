package domain

import "time"

// 管理员模型
type TbAdmin struct {
	Id            int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid          string    `gorm:"type:CHAR(32);NOT NULL"`
	UserName      string    `gorm:"type:VARCHAR(255);"`
	Pass          string    `gorm:"type:VARCHAR(255);"`
	CreateAt      time.Time `gorm:"type:DATETIME(0);"`
	UpdateAt      time.Time `gorm:"type:DATETIME(0);"`
	OperationAt   time.Time `gorm:"type:DATETIME(0);"`
	Operation     int32     `gorm:"type:INT(2);"`
	OperationText string    `gorm:"type:VARCHAR(500);"`
}
