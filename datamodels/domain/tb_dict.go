package domain

// 字典模型
type TbDict struct {
	Id       int32  `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Type     int32  `gorm:"type:INT(11);"`
	TypeName string `gorm:"type:VARCHAR(10);"`
	Name     string `gorm:"type:VARCHAR(30);"`
	Pid      int32  `gorm:"type:INT(11);"`
	Level    string `gorm:"type:VARCHAR(10);"`
	Order    int32  `gorm:"type:INT(11);"`
	Other    string `gorm:"type:VARCHAR(255);"`
}
