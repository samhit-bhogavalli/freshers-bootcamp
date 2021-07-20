package main

import "fmt"
//functions testing
func isEven(num int) bool {
	if num%2 == 0{
		return true
	} else {
		return false
	}
}
//function with 2 return types
func swap(x,y int) (int , int){
	return y,x
}
//variadic functions
func concat(str ...string) (result string){
	for _, i := range str {
		result+=i
	}
	return
}
//closure testing
func createSeq() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}
//sample of a recursion program
func total(num int) int{
	if num == 1 {
		return 1
	}
	return num + total(num-1)
}

func main(){
	fmt.Println(isEven(2))
	x,y :=  1,2
	x,y = swap(x,y)
	fmt.Println(x,y)
	fmt.Println(concat("samhit ","chowdary ", "bhogavalli"))
	slc := []string{"samhit ", "chowdary ", "bhogavalli"}
	fmt.Println(concat(slc...))

	newSeq1 := createSeq()
	newSeq1()
	newSeq1()
	newSeq2 := createSeq()
	newSeq2()
	newSeq2()
	fmt.Println(total(10))
}
