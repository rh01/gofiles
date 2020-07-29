package main

func main() {
	// model1: for-range,向通道发送迭代变量 
	// for _, s := range []string{"a", "b", "c"} {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case stringStream <-s:
	// 	}
	// }

	// 循环等待停止
	for {
		range {
			case <-done:
				return
			default:
		}
	}


}