package main

import (
	"fmt"
	"time"
)

/*
	>>>>>>>>>>>>Ticker
*/
func main() {
	ticker := time.NewTicker(time.Second * 1) //一次定义可以多次执行，其余和timer一样

	i := 0
	go func() {
		for {
			i++
			fmt.Printf("i:%d	%v\n", i, <-ticker.C)
			if i == 6 {
				break
			}
		}
	}()

	for {
	}
}
