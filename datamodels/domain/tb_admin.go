package domain

import "time"

// 管理员模型
type TbAdmin struct {
	Id            int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid          string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	UserName      string    `json:"userName" gorm:"type:VARCHAR(255);"`
	Pass          string    `json:"pass" gorm:"type:VARCHAR(255);"`
	CreatedAt     time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
	UpdatedAt     time.Time `json:"UpdatedAt" gorm:"type:DATETIME(0);"`
	OperationAt   time.Time `json:"operationAt" gorm:"type:DATETIME(0);"`
	Operation     int32     `json:"operation" gorm:"type:INT(2);"`
	OperationText string    `json:"operationText" gorm:"type:VARCHAR(500);"`
}
