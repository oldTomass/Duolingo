package model

type User struct {
	Uid         string `gorm:"column:uid" json:"uid"`
	UserName    string `gorm:"column:userName" json:"userName"`
	PassWord    string `gorm:"column:passWord" json:"passWord"`
	Number      string `gorm:"column:number" json:"number"`
	PayLevel    int    `gorm:"column:payLevel" json:"payLevel"`
	StudyLevel  int    `gorm:"column:studyLevel" json:"studyLevel"`
	Reading     int    `gorm:"column:reading;default:(-)" json:"reading"`
	Listening   int    `gorm:"column:listening;default:(-)" json:"listening"`
	Writing     int    `gorm:"column:writing;default:(-)" json:"writing"`
	CreatedTime string `gorm:"column:createdTime" json:"createdTime"`

	Errors []Error `gorm:"foreignKey:Uid;references:Uid;" json:"errors"`
}

func (user User) TableName() string {
	return "tb_user"
}
