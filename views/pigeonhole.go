package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

// Pigeonhole 文章归档
func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.PigOnHole
	// 查询出数据展示在页面上即可
	// 需要写哪些数据呢? 这个时候就需要在model中定义相关的结构体
	pigeonholeRes := service.FindPostPigonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
