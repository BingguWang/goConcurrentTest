package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	createPool(64, jobChan, resultChan) //创建工作池,先把运算协程都开起来待命

	//打印结果
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.Job.Id, result.Job.RandNum, result.Res)
		}
	}(resultChan)

	//插入任务
	var id int
	for { //无限循环是防止main退出过早
		id++
		ran := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: ran,
		}

		jobChan <- job
	}
}

/*设计一个线程池，计算每个随机数各个位的数的和*/
type Job struct {
	Id      int
	RandNum int //需要计算的随机数
}
type Result struct {
	Job *Job
	Res int //各位的和
}

//创建一个协程池,参数2：存任务的通道 参数3：存结果的通道
func createPool(num int, job chan *Job, resultChan chan *Result) { //第一个参数是要开几个协程

	for i := 0; i < num; i++ { //开启num个协程进行运算
		go func(jobChan chan *Job, resultChan chan *Result) { //开启协程来运算结果
			for job := range jobChan {
				ran := job.RandNum
				var re int
				for ran != 0 { //计算结果
					re += ran % 10
					ran = ran / 10
				}
				result := &Result{
					Job: job,
					Res: re,
				}
				resultChan <- result //结果存到结果的通道中
			}
		}(job, resultChan)
	}
}
