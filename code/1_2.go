package main

import (
	"fmt"
	"strings"
	"time"
)

func func1() { // join
	var arr []string
	for i := 1; i <= 10000; i++ {
		arr = append(arr, "asdfasdf")
	}

	start := time.Now() // 获取当前时间

	s := strings.Join(arr[:], " ")
	_ = s

	len := time.Since(start)
	fmt.Println("func1函数执行完成耗时：", len)
}

func func2() { //+
	var arr []string
	for i := 1; i <= 10000; i++ {
		arr = append(arr, "asdfasdf")
	}

	start := time.Now() // 获取当前时间
	s, sep := "", ""
	for _, arg := range arr[:] {
		s += sep + arg
		sep = " "
	}

	// fmt.Println(s)
	len := time.Since(start)
	fmt.Println("func2函数执行完成耗时：", len)
}

func main() {
	start := time.Now() // 获取当前时间
	go func1()
	go func2()

	time.Sleep(10000000000000000)

	len := time.Since(start)
	fmt.Println("main函数执行完成耗时：", len)
}
