package main

import (
	"fmt"
	"os"
	"github.com/lianglee123/easyconfig"
)

type MyConfig struct {
	Debug    bool   `config:"-"`  // do not load this field
	LogLevel string `config:"default:debug"`
	DB       DBConfig
}

type DBConfig struct {
	Host     string `config:"default:127.0.0.1"`
	Port     int    `config:"default:5432"`
	UserName string `config:"default:lianglee"`
	Pwd      string `config:"default:qwer1234"`
	DBName   string `config:"default:config_demo"`
}

func main() {
	opt := &easyconfig.LoadOption{
		EnvPrefix:      "DEMO",
		ConfigFilePath: "test.yaml",
	}
	config := &MyConfig{}
	os.Setenv("DEMO_LOG_LEVEL", "info")
	err := easyconfig.LoadConfig(config, opt)
	if err != nil {
		fmt.Printf("err happen when load config: %v \n", err)
		return
	}
	fmt.Printf("config: %+v", config)
}
