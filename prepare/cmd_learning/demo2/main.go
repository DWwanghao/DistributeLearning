package main

/*
	强制结束任务
 */
import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type  result struct{
	err error
	output []byte
}

func main() {
	resultChan := make(chan *result, 1000)

	ctx, cancelFunc := context.WithCancel(context.TODO())

	//协程
	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;ls -l")
		 output, err := cmd.CombinedOutput();
		 resultChan <- &result{
		 	err: err,
		 	output: output,
		 }
	}()
	time.Sleep(1*time.Second)
	cancelFunc()
	res := <-resultChan
	fmt.Println(res.err,string(res.output))

}
