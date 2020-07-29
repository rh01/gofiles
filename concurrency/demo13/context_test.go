package context_test

import (
	"sync"
	"testing"
)

func BenchmarkContext(b *testing.B) {
	var wg sync.WaitGroup

	// init two channel
	begin := make(chan struct{}, 0) // non buffer channel
	c := make(chan struct{}, 0)

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // block
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}

	receiver := func() {
		defer wg.Done()
		<- begin // block
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
