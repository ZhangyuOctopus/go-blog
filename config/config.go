package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

// 定义配置文件结构体
type tomlConfig struct {
	// 映射配置文件的view和system, 变量名大写外部文件才可以访问
	Viewer Viewer
	System SystemConfig
}

var Cfg *tomlConfig

func init() {
	//init()方法在程序启动的时候就会执行init()方法
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "go-blog"
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	Cfg.System.Version = 1.0
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
