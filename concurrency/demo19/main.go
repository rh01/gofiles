package main

import "sync"

func main() {
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	for condtionTrue() == false {
		c.Wait()
	}
	c.L.Unlock()
}
