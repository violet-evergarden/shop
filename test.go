package main

import (
	"fmt"
	"os"
)

//const test  = os.Getwd()
func main()  {
	dir,_ := os.Getwd()
	fmt.Println("当前路径：",dir)
}
