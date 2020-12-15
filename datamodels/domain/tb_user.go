package domain

import (
	"time"
)

// 用户模型
type TbUser struct {
	Id          int32     `json:"-" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Uuid        string    `json:"uuid" gorm:"type:CHAR(32);NOT NULL"`
	PUuid       string    `json:"pUuid" gorm:"type:CHAR(32);"`
	UserName    string    `json:"userName" gorm:"type:VARCHAR(50);"`
	Pass        string    `json:"-" gorm:"type:VARCHAR(255);"`
	Sex         int32     `json:"sex" gorm:"type:INT(2);"`
	Nick        string    `json:"nick" gorm:"type:VARCHAR(255);"`
	Phone       string    `json:"phone" gorm:"type:VARCHAR(30);"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"type:DATETIME(0);"`
	Experience  int32     `json:"experience" gorm:"type:INT(11);"`
	Level       int32     `json:"level" gorm:"type:INT(11);"`
	Gold        int32     `json:"gold" gorm:"type:INT(11);"`
	Flag        int32     `json:"flag" gorm:"type:INT(2);"`
	Description string    `json:"description" gorm:"type:VARCHAR(300);"`
	HeadImg     string    `json:"headImg" gorm:"type:VARCHAR(255);"`
	Birthday    time.Time `json:"birthday" gorm:"type:DATE;"`
	Email       string    `json:"email" gorm:"type:VARCHAR(50);"`
	Address     string    `json:"address" gorm:"type:VARCHAR(255);"`
	Type        int32     `json:"type" gorm:"type:INT(2);"`
	Modify      string    `json:"modify" gorm:"type:VARCHAR(50);"`
	ModifyAt    time.Time `json:"modifyAt" gorm:"type:DATETIME(0);"`
	Creater     string    `json:"creater" gorm:"type:VARCHAR(50);"`
	Age         int32     `json:"age" gorm:"type:INT(11);"`
	PostNum     int32     `json:"postNum" gorm:"type:INT(11);"`
	CommentNum  int32     `json:"commentNum" gorm:"type:INT(11);"`
}

// 校验用户名和密码
func (user *TbUser) CheckUserNameAndPass() (ok bool) {
	// TODO 校验用户名和密码
	ok = true
	return
}
