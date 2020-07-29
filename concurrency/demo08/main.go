package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 活锁的例子
// example: a demo for livelock
func main() {
	cadence := sync.NewCond(&sync.Mutex{})  // 使用锁创建一个条件变量，供线程等待或者宣布某事件的发生
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast() //唤醒所有等待cadence的线程
		}
	}()
	
	takeStep := func() {
		cadence.L.Lock() // 必须在wait之前上锁
		cadence.Wait() // Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L。
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprintf(out, ". Success!")
			return true
		}

		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool {return tryDir("right", &right, out) }

	// demo
	walk := func(walking *sync.WaitGroup, name string)  {
		var out bytes.Buffer
		defer func() {fmt.Println(out.String())}()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tossses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "alice")
	go walk(&peopleInHallway, "bob")
	peopleInHallway.Wait()
}