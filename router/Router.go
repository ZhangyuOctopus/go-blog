package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	// 在路由里面又要做一些区分: 1. 返回页面(views); 2. 返回api/json数据; 3. 返回静态数据

	// 因为在路由里面做了处理所以还是比较清晰的, 下面第一个是返回对应的页面, 注意方法名需要大写才可以访问到
	http.HandleFunc("/", views.HTML.Index)

	// 处理点击左边文章分类链接的请求: http://localhost:8081/c/1?page=1
	// 因为返回也是页面所以使用view.HTML来处理
	http.HandleFunc("/c/", views.HTML.Category)

	// 登录的请求处理
	http.HandleFunc("/login", views.HTML.Login)

	// 点击文章列表之后跳转到对应的文章详情页面, 类似于之前的点击分类链接的处理:
	http.HandleFunc("/p/", views.HTML.Detail)

	// 点击写文章的请求处理
	http.HandleFunc("/writing/", views.HTML.Writing)

	// 文章归档
	http.HandleFunc("/pigeonhole/", views.HTML.Pigeonhole)

	// 2. 返回对应的api处理
	http.HandleFunc("/api/v1/post", api.API.SaveAndPost)

	// 点击编辑文章发布之后更新发布信息
	http.HandleFunc("/api/v1/post/", api.API.GetPost)

	// 根据搜索的文字查询所有的文章
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	// 处理登录的api请求
	http.HandleFunc("/api/v1/login", api.API.Login)

	// 所有的静态资源走下面这个处理器, url的请求路径映射到文件中的
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/public/resource"))))
}
