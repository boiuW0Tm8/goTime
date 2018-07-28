package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println()
		for j := 1; j <= i; j++ {
			fmt.Printf("%v X %v = %v\t", i, j, i*j)
		}
	}
}
