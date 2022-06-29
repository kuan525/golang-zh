// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 若干是这个开头则true
		if strings.HasPrefix(url, "https://") == false {
			url = "https://" + url
		}
		// 创建HTTP请求 从url获取信息 resp是一个结构体，得到访问的请求结果
		resp, err := http.Get(url)

		// 获取错误
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// resp的Body字段包括一个可读的服务器响应流，获取网站内容
		siz, err := io.Copy(os.Stdout, resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)
		// 关闭流，防止资源泄露
		// resp.Body.Close()

		// 从Body获取失败
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		statu := resp.Status
		fmt.Printf("当前url的状态码是：%s\n", statu)
		fmt.Printf("当前url获取到的信息大小为：%d字节\n", siz)

		// 打印主体内容
		// fmt.Printf("%s", )
	}
}
