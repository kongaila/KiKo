package domain

import "time"

// 热榜模型
type TbTop struct {
	Id           int32     `json:"id" gorm:"type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Description  string    `json:"description" gorm:"type:VARCHAR(200);"`
	CreatedAt    time.Time `json:"createdAt" gorm:"type:DATETIME(0);"`
	Num          int32     `json:"num" gorm:"type:INT(11);"`
	ArticleUuid  string    `json:"articleUuid" gorm:"type:CHAR(32);"`
	ArticleTitle string    `json:"articleTitle" gorm:"type:VARCHAR(100);"`
	Weight       int64     `json:"weight" gorm:"type:INT(11);"` // 权重
}

// 设定排序规则。 用于热榜
type TopSlice []TbTop

func (t TopSlice) Len() int {
	return len(t)
}
func (t TopSlice) Less(i, j int) bool {
	return t[i].Weight > t[j].Weight
}

func (t TopSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
