package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
)

// 读取文件的内容，并按照带有缓冲的方式读取
func main() {
	file, err := os.Open("file/demo02/test.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("open file err", err)
	}

	// 创建一个带有缓冲reader
	// const (
	//		MaxScanTokenSize = 64 * 1024 // 用于缓冲一个token，实际需要的最大token尺寸可能小一些，例如缓冲中需要保存一整行内容
	//)
	reader := bufio.NewReader(file)
	// 循环读取文件的内容
	// ioutil.ReadFile是一次性读取完
	for {
		// 读到一个换行就结束
		s, err := reader.ReadString('\n')
		// 表示文件的末尾
		if err == io.EOF {
			break
		}
		// 输出内容
		fmt.Print(s)
	}

}
