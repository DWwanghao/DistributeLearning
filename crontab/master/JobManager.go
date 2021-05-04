package master

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

//任务管理器
type JobManager struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

var (
	//单例
	G_jobManager *JobManager
)

func InitJobManager() (err error) {
	//初始化配置
	config := clientv3.Config{
		Endpoints:   G_config.EtcdEndpoints,
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond,
	}
	//建立连接
	client, err := clientv3.New(config)
	if err != nil {
		return
	}

	//得到kv和lease的API子集
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	//给单例赋值
	G_jobManager = &JobManager{
		client: client,
		kv:     kv,
		lease:  lease,
	}

	return

}
