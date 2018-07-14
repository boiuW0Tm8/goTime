package main

import "fmt"

var x int = 42

func max(xx int) int {
	return xx + 22
}

func main() {
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	foo()
}

func foo() {
	var max = 88
	fmt.Println(max)
}
