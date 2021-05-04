package master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//程序的配置
type Config struct {
	ApiPort         int `json:"apiPort"`
	ApiReadTimeout  int `json:"apiReadTimeout"`
	ApiWriteTimeout int `json:"apiWriteTimeout"`
}

var G_config *Config

//加载配置
func InitConfig(filename string) (err error) {
	var conf Config
	//1、读取文件
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//2、JSON反序列化

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return
	}

	//3、赋值单例
	G_config = &conf
	fmt.Println(conf)
	return
}
