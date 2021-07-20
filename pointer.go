package main

import "fmt"

func passByValue(num int) {
	num++
}
func passByReference(num *int) {
	*num++
}

func main(){
	i := 0
	fmt.Println(i)
	passByValue(i)
	fmt.Println(i)
	passByReference(&i)
	fmt.Println(i, &i)
}
