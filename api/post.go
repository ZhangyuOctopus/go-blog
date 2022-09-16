package api

import (
	"errors"
	"fmt"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SaveAndPost Api结构体类型指针, 写文章之后发布文章
func (*Api) SaveAndPost(w http.ResponseWriter, r *http.Request) {
	// 获取用户id, 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录过期"))
		return
	}
	uid := claim.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		// Save操作
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		catrgoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: catrgoryId,
			UserId:     uid,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//update操作, 展示文章列表的详情, 与Save操作是类似的
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		catrgoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		pidStr := params["pid"].(float64)
		// 与Save操作不一样的是这里有pid了
		pid := int(pidStr)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: catrgoryId,
			UserId:     uid,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	pIdStr := strings.Trim(path, "/api/v1/post/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	fmt.Println(pIdStr)
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	pos, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, pos)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	// 解析表单用来获取url中第一个键值对的相关信息
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
