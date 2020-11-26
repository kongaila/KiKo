package domain

import "time"

// 用户模型
type TbUser struct {
	Id          int32     `gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `gorm:"type:CHAR(32);NOT NULL"`
	PUuid       string    `gorm:"type:CHAR(32);"`
	UserName    string    `gorm:"type:VARCHAR(50);"`
	Pass        string    `gorm:"type:VARCHAR(255);"`
	Sex         int32     `gorm:"type:INT(2);"`
	Nick        string    `gorm:"type:VARCHAR(255);"`
	Phone       string    `gorm:"type:VARCHAR(30);"`
	CreateAt    time.Time `gorm:"type:DATETIME(0);"`
	Experience  int32     `gorm:"type:INT(11);"`
	Level       int32     `gorm:"type:INT(11);"`
	Gold        int32     `gorm:"type:INT(11);"`
	Flag        int32     `gorm:"type:INT(2);"`
	Description string    `gorm:"type:VARCHAR(300);"`
	HeadImg     string    `gorm:"type:VARCHAR(255);"`
	Birthday    time.Time `gorm:"type:DATE;"`
	Email       string    `gorm:"type:VARCHAR(50);"`
	Address     string    `gorm:"type:VARCHAR(255);"`
	Type        int32     `gorm:"type:INT(2);"`
	Modify      string    `gorm:"type:VARCHAR(50);"`
	ModifyAt    time.Time `gorm:"type:DATETIME(0);"`
	Creater     string    `gorm:"type:VARCHAR(50);"`
	Age         int32     `gorm:"type:INT(11);"`
	PostNum     int32     `gorm:"type:INT(11);"`
}
