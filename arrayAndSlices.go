package main

import "fmt"

func main(){
	//arrays
	arr := [5]int{0,1,2,3,4}
	var array1 = [...]int{0,0,0}
	fmt.Println(arr)
	fmt.Println(array1, len(array1))
	//2D array
	var twoDArr [4][4]int
	fmt.Println(twoDArr)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			twoDArr[i][j] = i+j
		}
	}
	fmt.Println(twoDArr)
	//slices
	slc := make([]int, 4)
	fmt.Println(slc)
	slc1 := []int{1,2,3,4}
	fmt.Println(slc1)
	slc2 := slc1[1:2]
	fmt.Println(slc2)

}
