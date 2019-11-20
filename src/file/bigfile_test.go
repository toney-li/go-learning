package main

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestStringSearch(t *testing.T) {
	fmt.Printf("CPU %d", runtime.NumCPU())
	fmt.Println("Count d")
	keyword := "a."
	src := "a.bc.ef.a.kk"
	rs := strings.Count(src, keyword)
	fmt.Printf("Count %d", rs)
}

var sum = 0

// 两个channel汇总到一个channel

func TestChannel(t *testing.T) {
	summary := make(chan int)

	for i := 0; i < 10000; i++ {
		go cal(i, summary)
	}
	fmt.Printf("%d", totalCal(summary))
	time.Sleep(10000)
}

func cal(i int, summary chan int) {
	tmp := i
	summary <- tmp
}
func totalCal(summary chan int) int {
	calResult := <-summary
	sum += calResult
	return sum
}
