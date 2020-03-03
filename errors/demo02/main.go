package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	printError := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}

		err = underlyingError(err) 
		switch err {
		case os.ErrClosed:
			fmt.Printf("error(closed)[%d] %s",i, err)
			// return 
		case os.ErrInvalid:
			fmt.Printf("error(invalid)[%d] %s", i, err)
		case os.ErrPermission:
			fmt.Printf("error(permission)[%d] %s", i, err)
		}
	}

	printError()
}

// 潜在的错误
func underlyingError(err error) error {
	// 类型断言语句
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}
