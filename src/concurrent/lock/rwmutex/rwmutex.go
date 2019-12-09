package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// time.Sleep(time.Second)
	// wmutex()
	// fmt.Printf("Do Read lock")
	// time.Sleep(10 * time.Second)
	rmutex()
	// rmutex1()
}
func rmutex1() {
	var mutex *sync.RWMutex
	mutex = new(sync.RWMutex)
	mutex.Lock()

	mutex.Unlock()
	mutex.RLock()

	mutex.RUnlock()
}
func rmutex() {
	var mutex sync.RWMutex
	fmt.Printf("Lock the lock, Thread id \n")
	mutex.Lock()
	// 确保程序运行完成就释放锁,这个程序里需要主动释放锁，所以注释
	// defer mutex.Unlock()
	fmt.Printf("The lock is locked \n")
	chls := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		go func() {
			fmt.Println("Try to lock writing operation.")
			mutex.Lock()
			fmt.Println("Writing operation is locked.")

			fmt.Println("Try to unlock writing operation.")
			mutex.Unlock()
			fmt.Println("Writing operation is unlocked.")
		}()

		chls[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.RLock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.RUnlock()
			c <- i
		}(i, chls[i])

	}
	time.Sleep(time.Second)
	fmt.Println("Main unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range chls {
		data := <-c
		fmt.Printf("Recieve data %v \n", data)
	}
	time.Sleep(time.Second)

}
func wmutex() {
	var mutex sync.RWMutex
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
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range chls {
		data := <-c
		fmt.Printf("Recieve data %v \n", data)
	}
	time.Sleep(time.Second)
}
