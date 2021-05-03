package main
/*
  捕获任务输出
 */
import (
	"fmt"
	"os/exec"
)
func main() {
	//	cmd:=exec.Command("/bin/bash","-c","ls -l")
	cmd:=exec.Command("/bin/bash","-c","ls -l")
	if output, err := cmd.CombinedOutput();err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(output))
	}
}
