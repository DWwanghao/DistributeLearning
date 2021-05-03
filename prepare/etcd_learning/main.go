package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	//客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return

	}
	//实现put操作的demo
	/*kv := clientv3.NewKV(client)
	putResponse, err := kv.Put(context.TODO(), "/cron/jobs/job1", "wanghao",clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Revision",putResponse.Header.Revision)
		if putResponse.PrevKv != nil {
			fmt.Println("prevalue",string(putResponse.PrevKv.Value))
		}
	}*/
	//实现租约demo
	lease := clientv3.NewLease(client)
	grant, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := grant.ID

	//var keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
	var keepResp *clientv3.LeaseKeepAliveResponse

	//申请自动续租
	timeout, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	keepRespChan, err := lease.KeepAlive(timeout, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	//启动协程消费
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约到期了")
					goto END
				} else {
					fmt.Println("收到自动续租应答：", keepResp.ID)

				}

			}
		}
	END:
	}()

	kv := clientv3.NewKV(client)
	putResponse, err := kv.Put(context.TODO(), "/cron/lock/job1", "test_lease", clientv3.WithLease(id))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功", putResponse.Header.Revision)
	for {
		get, err := kv.Get(context.TODO(), "/cron/lock/job1")
		if err != nil {
			fmt.Println(err)
			return
		}
		if get.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期", get.Kvs)
		time.Sleep(2 * time.Second)

	}

}
