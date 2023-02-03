package mappers

import (
	"DuolinGo/model"
	"DuolinGo/utils"
)

func GetUserErrors(uid string) []model.Error {
	var errors []model.Error

	db := utils.GetDB()

	db.Preload("Sentence").Preload("Word")
	result := db.Joins("Sentence").Joins("Word").Where("uid = ?", uid).Find(&errors)

	err := result.Error
	num := result.RowsAffected
	if err != nil || num == 0 {
		return nil
	}
	return errors
}
