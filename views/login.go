package views

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

// Login 处理登录的请求
func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
