package controllers

import (
	"DuolinGo/mappers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func PageSentence(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	contentType, _ := strconv.Atoi(params["contentType"])
	contentLevel, _ := strconv.Atoi(params["contentLevel"])
	pageIndex, _ := strconv.Atoi(params["page"])
	fmt.Printf("contentType=%d contentLevel=%d pageINdex=%d", contentType, contentLevel, pageIndex)
	sentences := mappers.PageSentence(contentType, contentLevel, pageIndex)
	if sentences == nil || sentences == 0 {
		json.NewEncoder(w).Encode("Fail")
	} else {
		json.NewEncoder(w).Encode(sentences)
	}
}
