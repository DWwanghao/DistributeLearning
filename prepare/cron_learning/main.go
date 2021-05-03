package main
//定时器调度
import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	parse, err := cronexpr.Parse("*/5 * * * * * *");
	if 	err != nil {
		fmt.Println(err)
	}
	now:=time.Now()
	next:=parse.Next(now)
	time.AfterFunc(next.Sub(now), func() {
		fmt.Println("调度",next)
	})
	time.Sleep(5*time.Second)


	/*for i := 0; i < 5; i++ {
		fmt.Println(next)
		next=parse.Next(next)


	}*/
	/*fmt.Println(now)

	parse=parse*/


}
