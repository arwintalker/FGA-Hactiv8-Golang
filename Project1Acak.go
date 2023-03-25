package main

import (
	"fmt"
	"time"
)

func routine1(data interface{}, done chan bool) {
	for i := 0; i < 4; i++ {
		fmt.Println("GOROUTINE 1:", data)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func routine2(data interface{}, done chan bool) {
	for i := 0; i < 4; i++ {
		fmt.Println("GOROUTINE 2:", data)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func Project1Acak() {
	done1 := make(chan bool)
	done2 := make(chan bool)
	data1 := "Data1"
	data2 := "Data2"

	go routine1(data1, done1)
	go routine2(data2, done2)

	for i := 0; i < 8; i++ {
		select {
		case <-done1:
			fmt.Println("GOROUTINE 1 selesai")
		case <-done2:
			fmt.Println("GOROUTINE 2 selesai")
		}
	}
}
