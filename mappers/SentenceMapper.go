package mappers

import (
	"DuolinGo/model"
	"DuolinGo/utils"
)

/*	PageSentence 默认每页 10 条数据 */
func PageSentence(contentType, contentLevel, pageIndex int) interface{} {
	var sentence []model.Sentence
	result := utils.GetDB().Where("contentType = ? and contentLevel = ?", contentType, contentLevel).Limit(10).Offset((pageIndex - 1) * 10).Find(&sentence)
	err := result.Error
	num := result.RowsAffected
	if err != nil || num == 0 {
		return nil
	}
	return sentence
}
