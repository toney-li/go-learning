package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Printf("Lock the lock, Thread id \n")
	mutex.Lock()
	// 确保程序运行完成就释放锁,这个程序里需要主动释放锁，所以注释
	// defer mutex.Unlock()
	fmt.Printf("The lock is locked \n")
	chls := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		chls[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, chls[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Main unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range chls {
		fmt.Printf("Run goroutine func \n")
		data := <-c
		fmt.Printf("Recieve data %v \n", data)
	}
	time.Sleep(time.Second)
}
