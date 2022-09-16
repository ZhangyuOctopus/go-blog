package main

import (
	"go-blog/common"
	"go-blog/server"
)

// 加载template模板
func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {
	server.App.Start("127.0.0.1", "8081")
}
