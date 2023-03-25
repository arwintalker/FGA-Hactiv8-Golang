package main

import (
	"fmt"
	"sync"
	"time"
)

func routine1(data interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		mutex.Lock()
		fmt.Println("GOROUTINE 1:", data)
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func routine2(data interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		mutex.Lock()
		fmt.Println("GOROUTINE 2:", data)
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func Project1Rapi() {
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	data1 := "Data1"
	data2 := "Data2"

	wg.Add(2)
	go routine1(data1, &wg, mutex)
	go routine2(data2, &wg, mutex)

	wg.Wait()
}
