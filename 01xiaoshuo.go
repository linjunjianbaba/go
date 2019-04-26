package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HttpGet(url string) (rest string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4*1024)

	n, _ := resp.Body.Read(buf)
	for {
		if n == 0 {
			break
		}
		rest += string(buf[n:])
	}
	return
}
func PaQu(i int) {

	url := "http://www.quanben.co/sort/1_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页\n", i)

	rest, err := HttpGet(url)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("rest = ", rest)

}

func DoWord(start, end int) {
	fmt.Printf("准备爬取%d到%d的页面\n", start, end)
	for i := start; i <= end; i++ {
		//定义一个函数进行爬取
		// fmt.Printf("正在爬取第%d页\n", i)
		PaQu(i)

	}
}

func main() {
	var start, end int
	fmt.Printf("请输入开始爬取页面:")
	fmt.Scan(&start)
	fmt.Printf("请输入爬取结束页面:")
	fmt.Scan(&end)
	DoWord(start, end)
}
