## Duolingo
### server：9009
### go run/build main.go
### Database
	tb_user , tb_sentence , tb_word , tb_error

## gorm
  go get -u gorm.io/driver/mysql
  go get -u gorm.io/gorm/logger
  go get -u gorm.io/gorm
## router
  go get -u github.com/gorilla/mux 
## 热加载
  go get github.com/pilu/fresh
  go install github.com/pilu/fresh@latest

## router
#### 初始化路由
  router = mux.NewRouter()
#### 设置路由
  setRouter()
#### 开启监听服务
  http.ListenAndServe(":9009", router)