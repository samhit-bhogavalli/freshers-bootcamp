package main

import "fmt"

func call(){
	defer fmt.Println("defer in function")
	fmt.Println("in function")
}

func main() {
	total := 0
	//for loop testing
	for i := 0; i < 10; i++{
		j := i
		for ; j < 10;{
			total++
			j++
		}
	}
	//while testing
	i := 0
	for i < 100 {
		total++
		i++
	}
	fmt.Println(total)
	//if condition testing
	num := 9
	if temp:="string";num%2 == 0 {
		fmt.Println("num is a even number")
		fmt.Println("temp is declared inside if-", temp)
	} else {
		fmt.Println("num is a odd number")
		fmt.Println("temp is declared inside if-", temp)
	}
	if temp := num; temp < 10 {
		fmt.Println("num is less than 10")
	}
	//switch statement testing
	switch num := 9; num {
	case 1:
		fmt.Println("first is case running")
	case 2:
		fmt.Println("second case is running")
	default:
		fmt.Println("default case is running")
	}
	//switch with no condition
	switch  {
	case false:
		fmt.Println("first case is running")
	case true:
		fmt.Println("second case is running")
	}
	//defer testing
	defer fmt.Println("defer in main")
	fmt.Println("hello")

	call()
	fmt.Println("after call function")
}
