package main

import (
	"fmt"
	"time"
)

/*
>>>>>>>>>>>>>range接收channel数据

*/
func main() {
	ch := make(chan int)

	go func() {
		for i := 6; i <= 8; i++ {
			ch <- i
			time.Sleep(time.Second) //这里等待是为了输出数据时慢一点好看一些，没其他作用
		}
	}()

	for rece := range ch {
		fmt.Println(rece)
		if rece == 8 {
			break
		}
	}
}
