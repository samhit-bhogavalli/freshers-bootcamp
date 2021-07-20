package main

import "fmt"

var java , python bool

func main(){
	var temp int
	var varcheck1 = true
	varcheck2 := true
	varcheck3, varcheck4 := 10, "string"
	fmt.Println(temp, java, python, varcheck1, varcheck2)
	fmt.Println(varcheck3, varcheck4)
}
