package main

import "fmt"

func main() {
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	case uint8, uint16:
		fmt.Println("uint8")
	// case byte:
	// 	fmt.Println("byte")
	default:
		fmt.Printf("unsupport type: %T\n", t)
	}
}
