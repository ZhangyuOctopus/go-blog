package api

import (
	"fmt"
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名和密码, 返回对应的json数据
	// 这里需要注意一下json的数据不可以通过r.ParseForm和r.Form.Get()拿取到的, 需要对应的处理
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	fmt.Println("username: ", userName, "password: ", passwd)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	// 将返回值写入到json字符串中
	common.Success(w, loginRes)
}
