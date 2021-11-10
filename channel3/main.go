package main

import "fmt"

//缓冲通道

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2 //因为是带缓存的，所以可以发多个数据
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
