package views

import (
	"errors"
	"fmt"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Detail views包是用来渲染html页面的, 调用service层的方法将数据返回填充大奥html页面
func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	// 获取模板最后渲染模板
	detail := common.Template.Detail
	// 获取路径参数
	path := r.URL.Path
	// 输出对应的路径
	fmt.Println("====", path)
	// 去除掉前面的/c/那么就可以得到对应的分类id, 下面得到的是7.html所以还需要进一步截取
	pIdStr := strings.Trim(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	fmt.Println(pIdStr)
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		log.Println("转换分类id为整数的时候出现错误: ", err)
		return
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错!"))
		return
	}
	detail.WriteData(w, postRes)
}
