package router

import (
	"DuolinGo/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router

func init() {
	// 初始化路由
	router = mux.NewRouter()
	// 设置路由
	setRouter()
	// 开启监听服务
	http.ListenAndServe(":9009", router)
}

/* 设置路由导航 */
func setRouter() {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/error/{uid}", controllers.GetErrors).Methods("GET")

	// 查询第一页的课程/sentence/{contentType}/{contentLevel}/{page}
	router.HandleFunc("/sentence/{contentType}/{contentLevel}/{page}", controllers.PageSentence).Methods("GET")

}
