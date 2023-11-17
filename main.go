package main

import "fmt"

func main() {
	fmt.Println("begin")
	panic("panic!!!")
	fmt.Println("end")
}
