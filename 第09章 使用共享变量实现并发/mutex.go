package main

import (
	"fmt"
	"sync"
)

var (
	//mu   sync.Mutex
	mu   sync.RWMutex
	data int
)

func GotData() int {
	mu.RLock()
	defer mu.RUnlock()
	return gotData()
}

func gotData() int {
	return data
}

func SetData(param int) {
	mu.Lock()
	defer mu.Unlock()
	setData(param)
}

func setData(param int) {
	data = param
}

func ChangeDataForTimes(n int, ch chan bool) {
	for i := 0; i < 10000; i++ {
		SetData(i)
		fmt.Println("set data to ", i, "by", n, "goroutine")
	}
	ch <- true
}

func ShowData() {
	for {
		showData()
	}
}

func showData() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("data = ", gotData())
}
func main() {
	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go ChangeDataForTimes(i, ch)
	}

	go ShowData()

	for i := 0; i < 10; i++ {
		<-ch
	}
}
