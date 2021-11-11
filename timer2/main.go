package main

import (
	"fmt"
	"time"
)

//如果不用timer，自己手动实现定时器就是这样实现
func main() {
	timetout := Timer(5 * time.Second)
	for {
		select {
		case <-timetout: //取到就行
			fmt.Println("5秒到了")
			return
		}
	}
}

func Timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)

		//时间到了
		ch <- true
	}()
	return ch
}
