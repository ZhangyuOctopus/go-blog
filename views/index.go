package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Index 处理页面的请求
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// service层调用dao层查询数据局
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败: ", err)
		index.WriteError(w, errors.New("系统错误, 请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	// 需要传递每页显示的数量pageSize, 每页显示的数量
	page, pageSize := 1, 10
	if pageStr != "" {
		// 将字符串转为int类型整数
		page, _ = strconv.Atoi(pageStr)
	}
	path := r.URL.Path
	// 去除掉/前面的字符
	slug := strings.TrimPrefix(path, "/")
	// 由于涉及到分页所以在原来的基础上传递页数和每页查询的数量
	// 多传递一个slug参数这样在service层判断的时候判断路径中是否存在这个参数增加slug查询条件的限制
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据错误: ", err)
		index.WriteError(w, errors.New("系统错误, 请联系管理员"))
	}
	index.WriteData(w, hr)
}
