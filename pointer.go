package main

import "fmt"

func fun1(x int){
	x = 0
}

func fun2(xPtr *int){
	*xPtr = 0 // dereferencing
}

func num(yPtr *int){
	*yPtr = 5
}

func main () {
	x := 10
	fun1(x)
	fmt.Println("x: ", x) // this will print out 10

	fun2(&x) // variable address
	fmt.Println("x: ", x) // this will print out 0

	// New Function Pointer Example
	yPtr := new(int)
	num(yPtr)
	fmt.Println("yPtr: ", *yPtr)
}