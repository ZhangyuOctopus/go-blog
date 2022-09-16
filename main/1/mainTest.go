package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// 定义一个IndexData(json格式的字符串)
type IndexData struct {
	// 下面是`json:title`表示展示的时候按照小写来展示
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	// 指明当前拿到的是json字符串, Contype-Type设置浏览器按照什么方式来解析当前的html代码
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go博客项目"
	indexData.Desc = "入门级别项目"
	// 将结构体转为json格式的字符串
	jsonStr, _ := json.Marshal(indexData)
	fmt.Println("index")
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	// 指明当前拿到的是json字符串
	var indexData IndexData
	indexData.Title = "go博客项目"
	indexData.Desc = "入门级别项目"
	t := template.New("indexpage.html")
	// 1.拿到当前的路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/indexpage.html")
	// 解析html代码, indexData表示在html页面中需要展示的数据,
	t.Execute(w, indexData)
}

func main() {
	// 程序入口,一个项目只能够有一个入口
	// web程序, http协议 ip:port
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	// HandleFunc函数处理请求"/"请求
	http.HandleFunc("/", index)
	// 处理"/indexpage.html"请求
	http.HandleFunc("/indexpage.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
