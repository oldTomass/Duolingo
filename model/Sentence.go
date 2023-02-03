package model

type Sentence struct {
	Sid          int    `gorm:"column:sid" json:"sid"`
	SentenceEn   string `gorm:"column:sentenceEn" json:"sentenceEn"`
	SentenceZn   string `gorm:"column:sentenceZn" json:"sentenceZn"`
	ContentType  int    `gorm:"column:contentType" json:"contentType"`
	ContentLevel int    `gorm:"column:contentLevel" json:"contentLevel"`
	AdminId      string `gorm:"column:adminId" json:"adminId"`
	CreatedTime  string `gorm:"column:createdTime" json:"createdTime"`
	UpdateTime   string `gorm:"column:updateTime" json:"updateTime"`
}

func (sentence Sentence) TableName() string {
	return "tb_sentence"
}
