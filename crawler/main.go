package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var start, end int
	fmt.Println("请输入起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入结束页：")
	fmt.Scan(&end)

	Run(start, end)
}

func GetData(url string) (result string, err error) {
	resp, err := http.Get(url) //作为客户端或取响应数据
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// buffer := make([]byte, 1024*4)
	for {
		// n, err1 := resp.Body.Read(buffer)
		bytes, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				log.Fatal("读取内容失败：", err1)
				break
			}

		}
		result += string(bytes)
		return result, err
	}
	return
}

func GetDateByPage(i int, page chan<- int) {
	// url := "https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html"
	url := "https://bbs.hupu.com/lakers-" + strconv.Itoa(i)
	fmt.Printf("正在爬取第%d个网页。。。\n", i)
	fmt.Println(url)
	res, err := GetData(url)
	if err != nil {
		log.Fatal("爬取失败：", err)
		return
	}
	//爬取的内容写入文件
	fileName := "page" + strconv.Itoa(i) + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal("文件创建错误：", err)
		return
	}
	//写内容
	f.WriteString(res)
	f.Close()
	page <- i //写完把页码传到page中

}

//打算为每个页面分一个协程爬取
func Run(start int, end int) {
	fmt.Printf("正在爬取第%d页到第%d页..\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go GetDateByPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页的数据爬取写入完毕！！！\n", <-page)
	}
}

//TODO 这里的数据是没过滤的，学学怎么过滤数据
