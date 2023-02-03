package controllers

import (
	"DuolinGo/mappers"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetErrors(w http.ResponseWriter, r *http.Request) {
	// 解析url中的参数/error/{uid}
	uid := mux.Vars(r)["uid"]
	errors := mappers.GetUserErrors(uid)

	json.NewEncoder(w).Encode(errors)
}
