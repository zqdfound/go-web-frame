package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	fmt.Println("test")
}

var set = make(map[int]bool, 0)
var m sync.Mutex

func printOnce(num int) {
	m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}

func TestPrintOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
