package model

type Error struct {
	Eid         int    `gorm:"column:eid" json:"eid"`
	Uid         string `gorm:"column:uid" json:"-"`
	Sid         int    `gorm:"column:sid" json:"-"`
	Wid         int    `gorm:"column:wid" json:"-"`
	ErrorType   int    `gorm:"column:errorType" json:"errorType"`
	CreatedTime string `gorm:"column:createdTime" json:"createdTime"`

	Sentence Sentence `gorm:"foreignKey:Sid;references:Sid;" json:"sentence,omitempty"`
	Word     Word     `gorm:"foreignKey:Wid;references:Wid;" json:"word,omitempty"`
}

func (error Error) TableName() string {
	return "tb_error"
}
