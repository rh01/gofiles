package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, request := range []string{"", "hello"} {
		resp, err := echo(request)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}

	response = fmt.Sprintf("echo: %s", request)
	return
}
