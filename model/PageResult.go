package model

type PageResult struct {
	Total     int64         `json:"total"`
	TotalPage int64         `json:"totalPage"`
	Data      []interface{} `json:"data"`
}
