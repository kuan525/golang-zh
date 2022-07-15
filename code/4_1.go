package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashMethod = flag.Int("s", 256, "选择哈希版本：256、384、512。")

func main() {
	flag.Parse()
	printHash()
}

func printHash() {
	var s string
	fmt.Println("输入要解析的字符串：")
	fmt.Scanf("%s\n", s)

	switch *hashMethod {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(s)))

	}
}
