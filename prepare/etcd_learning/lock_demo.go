package main

func main() {
	//1、lease实现锁自动过期
	//2、op操作
	//3、txn事务，支持if else then
	//TODO：step1：上锁（创建租约，自动续租，拿着租约去抢占一个key）
	//TODO:step2：处理业务
	//TODO:step3：释放锁（取消自动续租，释放租约）

}
