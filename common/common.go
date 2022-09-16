package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	// 因为加载模板比较耗时所以需要开启goroutine使得可以并发执行
	// 使用sync包中的相关操作使得开启的goroutine和main()goroutine同步
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// 当函数返回的时候注册的goroutine数目减1
		defer wg.Done()
		// 耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/template/")
		if err != nil {
			// 因为模板都加载不出来所以直接挂掉即可
			panic(err)
		}
	}()
	wg.Wait()
}

// GetRequestJsonParam 获取对应的json参数(例如登录的时候)
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

// Success 因为写入成功是可以复用的, data为传递的数据
func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	// 将结构体转为json格式的字符串
	result.Error = ""
	result.Data = data
	result.Code = 200
	resultJson, _ := json.Marshal(result)
	// 告诉浏览器我们返回的是json方面的数据
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

// Error 写入数据失败是可以复用的
func Error(w http.ResponseWriter, err error) {
	var result models.Result
	// 将结构体转为json格式的字符串
	result.Error = ""
	result.Code = 721
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println("登录失败: ", err)
	}
}
