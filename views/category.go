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

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	// 获取url请求中的参数: http://localhost:8081/c/1?page=1
	path := r.URL.Path
	// 输出对应的路径
	fmt.Println("====", path)
	// 去除掉前面的/c/那么就可以得到对应的分类id
	cIdStr := strings.Trim(path, "/c/")
	fmt.Println(cIdStr)
	cId, err := strconv.Atoi(cIdStr)
	categoryTemplate := common.Template.Category
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		log.Println("转换分类id为整数的时候出现错误: ", err)
	}

	// 因为分类也涉及到分页, 所以还需要将index中分页的代码拷贝进来
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败: ", err)
		categoryTemplate.WriteError(w, errors.New("系统错误, 请联系管理员"))
		return
	}
	// 获取url路径中第一个参数的值, 例如路径:http://localhost:8081/c/2?page=2, 获取到的值为2
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10

	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		// 将错误写回去
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
