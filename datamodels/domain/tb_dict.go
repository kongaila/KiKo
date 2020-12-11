package domain

import "time"

// 字典模型
type TbDict struct {
	Id        int32     `json:"id" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Type      int32     `json:"type" gorm:"type:INT(11);"`
	TypeName  string    `json:"typeName" gorm:"type:VARCHAR(10);"`
	Name      string    `json:"name" gorm:"type:VARCHAR(30);"`
	Pid       int32     `json:"pid" gorm:"type:INT(11);"`
	Level     string    `json:"level" gorm:"type:VARCHAR(10);"`
	Order     int32     `json:"order" gorm:"type:INT(11);"`
	Other     string    `json:"other" gorm:"type:VARCHAR(255);"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
}
