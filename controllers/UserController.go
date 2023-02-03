package controllers

import (
	"DuolinGo/mappers"
	"DuolinGo/model"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// 解码Json到User
//json.NewDecoder(r.Body).Decode(&user)
// 编码User到Json
//json.NewEncoder(w).Encode(peter)

const CookieName = "GO_COOKIE_HJM_0627"

func Home(w http.ResponseWriter, r *http.Request) {
	users := mappers.Home()
	if users != nil && users != 0 {
		json.NewEncoder(w).Encode(users)
	} else {
		json.NewEncoder(w).Encode("fail")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 解析数据
	u := model.User{
		UserName: r.FormValue("userName"),
		PassWord: r.FormValue("passWord"),
	}
	//json.NewDecoder(r.Body).Decode(&u)

	user := mappers.Login(u)

	if user != nil {

		// 判读那是否存在cookie
		cookie, err := r.Cookie(CookieName)
		if err == nil {
			cookie.MaxAge = 10 // 存在cookie，重新设置10s过期时间
		} else {
			// 不存在就新建cookie，然后设置到request中
			c := &http.Cookie{Name: CookieName, Value: u.UserName, MaxAge: 10}
			c.Path = "/"
			http.SetCookie(w, c)
		}

		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode("fail")
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	// 未对数据进行校验
	u := model.User{
		Uid:         uuid.New().String()[0:32],
		UserName:    r.FormValue("userName"),
		PassWord:    r.FormValue("passWord"),
		Number:      r.FormValue("number"),
		CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	is := mappers.Register(u)
	if is {
		json.NewEncoder(w).Encode("Success")
	} else {
		json.NewEncoder(w).Encode("Fail")
	}
}
