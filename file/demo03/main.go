package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	file, err := ioutil.ReadFile("file/demo02/test.txt")

	if err != nil {
		log.Fatal("open file err", err)
	}

	// 创建一个带有缓冲reader
	// const (
	//		MaxScanTokenSize = 64 * 1024 // 用于缓冲一个token，实际需要的最大token尺寸可能小一些，例如缓冲中需要保存一整行内容
	//)
	fmt.Println(string(file))
}
