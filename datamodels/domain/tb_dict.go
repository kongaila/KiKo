package domain

import "time"

// 字典模型
type TbDict struct {
	Id        int32     `json:"id" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Code      string    `json:"type" gorm:"type:VARCHAR(20);"`
	CodeName  string    `json:"typeName" gorm:"type:VARCHAR(20);"`
	Name      string    `json:"name" gorm:"type:VARCHAR(30);"`
	Pid       int32     `json:"pid" gorm:"type:INT(11);"`
	Level     string    `json:"level" gorm:"type:VARCHAR(10);"`
	Order     int32     `json:"order" gorm:"type:INT(11);"`
	Other     string    `json:"other" gorm:"type:VARCHAR(255);"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:DATETIME(0);"`
}
