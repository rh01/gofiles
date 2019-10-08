package main

import (
	"go-learn/type/struct/demo01/model"
	"fmt"
)

func main() {
	// 创一个导出类型只需要引包即可
	student := model.NewStudent("shhh", 10)
	//&student
	fmt.Println(student.GetName())
}
