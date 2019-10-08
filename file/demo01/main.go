package main

import (
	"os"
	"log"
	"fmt"
)

func main() {

	//ioutil.ReadFile()
	//os.NewFile()

	//文件指针，文件对象
	file, err := os.Open("file/demo01/test.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("open file err", err)
	}

	fmt.Println("file=", file)

}
