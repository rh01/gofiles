package main

import "fmt"

func writedata(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func readData(data chan int, exitChan chan struct{}) {
	for {
		v, ok := <-data
		if !ok {

			exitChan <- struct{}{}

			break
		}

		fmt.Println(v)

	}
	// exitChan <- struct{}{}
	close(exitChan)
}

func main() {
	data := make(chan int, 10)
	go writedata(data)

	exitChan := make(chan struct{}, 2)

	// for i := 0; i < 2; i++ {
	go readData(data, exitChan)
	// }

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
