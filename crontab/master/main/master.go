package main

import (
	"Project/crontab/master"
	"flag"
	"fmt"
	"runtime"
)

var (
	conFile string //配置文件路径
)

//初始化命令行参数
func initArgs() {
	//master -config ./master.json
	flag.StringVar(&conFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

//线程数与主机cpu核数一致
func initEnv() {

	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	var err error
	//初始化命令行参数
	initArgs()

	//初始化线程
	initEnv()

	//加载配置(文件名经常从命令行参数传进来)
	err = master.InitConfig(conFile)
	if err != nil {
		goto ERR
	}
	//任务管理器
	err = master.InitJobManager()
	if err != nil {
		goto ERR
	}

	//启动Api Http服务
	err = master.InitApiServer()
	if err != nil {
		goto ERR
	}
	//正常退出
	return
	//打印错误
ERR:
	fmt.Println(err)

}
