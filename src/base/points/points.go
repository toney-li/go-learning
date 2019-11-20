package main

import (
	"fmt"
)

func main() {
	a := 8
	var b *int
	//Print result:nil Type b :*int
	fmt.Printf("nil Type b: %T \n", b)
	//Print result:nil Address b :<nil>
	fmt.Println("nil Value b:", b)
	// 报错:panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println("nil Address b:", *b)
	// b 获取了a变量的地址
	b = &a
	// Print result:Type b :*int
	fmt.Printf("Type b : %T\n", b)
	// Print result: Address b : 0xc000082000
	fmt.Printf("Address b : %p \n", b)
	// Print result: Value b : 8
	fmt.Printf("Value b: %d \n", *b)
	// Go 不支持指针运算 ,invalid operation: b++ (non-numeric type *int)
	// b++

	// 改变a值b值也改变
	a++
	fmt.Printf("Address a: %p \t b:%p \n", a, b)
	fmt.Printf("Value a:%v b: %v \n", a, *b)
	// 改变b值a值也改变
	*b++
	fmt.Printf("Address a: %p \t b:%p \n", a, b)
	fmt.Printf("Value a:%v b: %v \n", a, *b)
	// 传递指针给函数
	changeB(b)
	fmt.Printf("Address a: %p \t b:%p \n", a, b)
	fmt.Printf("Value a:%v b: %v \n", a, *b)
	// 值传递不改变值
	changeA(a)
	fmt.Printf("Address a: %p \t b:%p \n", a, b)
	fmt.Printf("Value a:%v b: %v \n", a, *b)
	// 数组，使用切片传递，而不是指针
	arr := []int{10, 20, 30, 40}
	changeArray(arr)
	fmt.Println(arr)
	changeArr(&arr)
	fmt.Println(arr)
	// 虽然可以通过传递数组指针给函数的方式来修改原始数组的值，但这在 Go 中不是惯用的方式，我们可以使用切片完成同样的事情。
	//切片是对数组的抽象。Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
	changeSlice(arr[:])
	fmt.Println(arr)

}
func changeSlice(slice []int) {
	slice[3] = slice[3] + 1
}
func changeArray(arr []int) {
	arr[0] = arr[0] + 1
}
func changeArr(arr *[]int) {
	(*arr)[1] = (*arr)[1] + 1
}
func changeB(b *int) {
	*b++
}
func changeA(a int) {
	a++
}
