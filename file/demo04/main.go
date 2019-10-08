package main

import (
	"log"
	"os"
	"bufio"
)

func main() {

	// 写文件，一般使用os.OpenFile()
	file, err := os.OpenFile("file/demo01/test01.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal("open file err", err)
	}

	//file.Write([]byte("hello, Garden"))
	// 写入5句
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello, Garden\n")
	}

	//因为writer是带有缓存的，因此在调用writeString方法时，其实
	//内容时先写到缓存中，所所以需要调用红Flush方法，将缓存的数据真正
	// 写到文件中，否则文件会没有数据
	defer writer.Flush()
}
