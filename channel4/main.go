package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	time.Sleep(5 * time.Second)
	close(ch)
	//for v := range ch { // range遍历，会在缓存没有的时候还去取，导致panic
	//	fmt.Println(v)
	//}

	for { // 正确的遍历方式
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
