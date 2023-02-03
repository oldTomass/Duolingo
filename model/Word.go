package model

type Word struct {
	Wid         int    `gorm:"column:wid" json:"wid"`
	WordEn      string `gorm:"column:wordEn" json:"wordEn"`
	WordZn      string `gorm:"column:wordZn" json:"wordZn"`
	IsSynonyms  int    `gorm:"column:isSynonyms" json:"isSynonyms"`
	AdminId     string `gorm:"column:adminId" json:"adminId"`
	CreatedTime string `gorm:"column:createdTime" json:"createdTime"`
	UpdateTime  string `gorm:"column:updateTime" json:"updateTime"`
}

func (word Word) TableName() string {
	return "tb_word"
}
