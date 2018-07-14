package main

import (
	"fmt"

	"github.com/boiuW0Tm8/goTime/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
