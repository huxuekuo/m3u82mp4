package main

import (
	"fmt"
	"sync"
)

var Cnt int
var lock sync.Mutex

func Add(iter int) {
	lock.Lock()
	for i := 0; i < iter; i++ {
		Cnt++ // 操作共享资源
	}
	lock.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	// 两个线程操作同一共享资源
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			Add(100000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(Cnt)
}
