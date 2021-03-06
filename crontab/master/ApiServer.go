package master

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

//任务的http接口
type ApiServer struct {
	httpServer *http.Server
}

//创建单利模式
var (
	//单利对象
	G_apiServer *ApiServer
)

//保存任务接口的回调函数
func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

//初始化服务
func InitApiServer() (err error) {
	//配置路由
	mux := http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	//启动TCP监听
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort))
	if err != nil {
		return
	}
	//创建一个http服务
	server := &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}
	//赋值单例
	G_apiServer = &ApiServer{
		httpServer: server,
	}
	//启动服务
	go server.Serve(listener)

	return

}
