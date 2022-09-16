package server

import (
	"go-blog/router"
	"log"
	"net/http"
)

var App = &MyServer{}

type MyServer struct {
}

func (*MyServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	// 调用自定义router包中的Router()函数
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
