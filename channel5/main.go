package main

import "fmt"

/**
使用channel也是能实现并发的，通过channel的通信来实现并发，如下
main和子函数的并发就是通过通道来实现的，但是可以看到这样很不方便,需要创建通道数组，如果子协程派生出新的协程就不方便了
用waitGroup和context都可以
*/
func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go func(ch chan int, i int) {
			ch <- i
		}(chs[i], i)
	}

	for _, ch := range chs {
		v := <-ch
		fmt.Println(v)
	}
}
